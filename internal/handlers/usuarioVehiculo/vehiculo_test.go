package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	handlers "cmd/rideUleam/internal/handlers/usuarioVehiculo"
	"cmd/rideUleam/internal/middleware"
	models "cmd/rideUleam/internal/models/usuarioVehiculo"
	service "cmd/rideUleam/internal/service/usuarioVehiculo"
	storage "cmd/rideUleam/internal/storage/usuarioVehiculo"
)

// usuarioRepoFake: repositorio de usuarios en memoria para los tests de handler.
type usuarioRepoFake struct {
	porEmail map[string]models.Usuario
	nextID   int
}

func nuevoUsuarioRepoFake() *usuarioRepoFake {
	return &usuarioRepoFake{porEmail: map[string]models.Usuario{}, nextID: 1}
}

func (f *usuarioRepoFake) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *usuarioRepoFake) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

// construirEntorno arma el MISMO router que main.go (mismas rutas, mismo
// middleware.Auth real) pero con almacen en memoria y repo de usuarios fake.
// Devuelve el handler listo para httptest y un token valido ya emitido.
//
// Clave pedagogica: probamos a traves del middleware REAL, no de uno simplificado.
// Si el wiring de la ruta protegida se rompe, este test se entera.
func construirEntorno(t *testing.T) (http.Handler, string) {
	t.Helper()

	almacen := storage.NuevaMemoria()
	almacen.SeedVehiculos()
	usuarios := nuevoUsuarioRepoFake()

	vehiculoSvc := service.NuevoVehiculoService(almacen)

	authSvc := service.NuevoAuthService(usuarios)
	srv := handlers.NewServer(vehiculoSvc, authSvc)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", srv.Registrar)
		r.Post("/auth/login", srv.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc)) // <- el middleware real de S10
			r.Get("/vehiculos", srv.ListarVehiculos)
			r.Post("/vehiculos", srv.CrearVehiculo)
			r.Get("/vehiculos/{id}", srv.ObtenerVehiculo)
			r.Put("/vehiculos/{id}", srv.ActualizarVehiculo)
			r.Delete("/vehiculos/{id}", srv.BorrarVehiculo)
		})
	})

	token := registrarYObtenerToken(t, r)
	return r, token
}

// registrarYObtenerToken hace register + login contra el propio router para
// conseguir un JWT valido, igual que lo haria un cliente real.
func registrarYObtenerToken(t *testing.T, h http.Handler) string {
	t.Helper()
	cred := `{"email":"docente@uleam.edu.ec","password":"secreta123"}`

	reqReg := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(cred))
	h.ServeHTTP(httptest.NewRecorder(), reqReg)

	reqLogin := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(cred))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, reqLogin)
	require.Equal(t, http.StatusOK, rec.Code, "el login deberia devolver 200")

	var resp struct {
		Token string `json:"token"`
	}
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	require.NotEmpty(t, resp.Token)
	return resp.Token
}

// TestCrearProducto_Exitoso: POST con token y cuerpo valido -> 201 Created.

func TestCrearVehiculo_Exitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"conductor_id":1,
		"placa":"ABC-123",
		"marca":"Toyota",
		"modelo":"Corolla",
		"capacidad":5
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/vehiculos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado models.Vehiculo
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.ID)
	assert.Equal(t, 1, creado.ConductorID)
	assert.Equal(t, "ABC-123", creado.Placa)
	assert.Equal(t, "Toyota", creado.Marca)
	assert.Equal(t, "Corolla", creado.Modelo)
	assert.Equal(t, 5, creado.Capacidad)
}

// TestObtenerProducto_NoEncontrado: id inexistente -> 404 Not Found.

func TestObtenerVehiculo_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/vehiculos/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCrearVehiculo_Invalido(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{
    "conductor_id":0,
    "placa":"ABC-123",
    "marca":"Toyota",
    "modelo":"Corolla",
    "capacidad":5
}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/vehiculos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// TestRutaProtegida_SinToken: sin header Authorization, el middleware corta
// antes de llegar al handler -> 401 Unauthorized.
func TestRutaVehiculoProtegida_SinToken(t *testing.T) {
	h, _ := construirEntorno(t)
	body := `{
    "conductor_id":1,
    "placa":"ABC-123",
    "marca":"Toyota",
    "modelo":"Corolla",
    "capacidad":5
}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/vehiculos", strings.NewReader(body))
	// A proposito: NO seteamos Authorization.
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
