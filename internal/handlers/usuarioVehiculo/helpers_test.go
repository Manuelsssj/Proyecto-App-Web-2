package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	handlers "RideUleam/internal/handlers/usuarioVehiculo"
	"RideUleam/internal/middleware"
	models "RideUleam/internal/models/usuarioVehiculo"
	service "RideUleam/internal/service/usuarioVehiculo"
	storage "RideUleam/internal/storage/usuarioVehiculo"
)

// =====================================================================
// Dobles de prueba
//
// Un test de la capa HTTP NO debe tocar la base de datos real: la
// reemplazamos por estos dobles en memoria. Son SOLO para los tests
// (viven en archivos _test.go); en produccion corre GORM. La base real
// tiene su propio test en internal/storage/sqlite_test.go.
// =====================================================================

// almacenFake implementa storage.Almacen (productos + categorias) en memoria.
type almacenFake struct {
	vehiculos map[int]models.Vehiculo

	nextProd int
	nextCat  int
}

func nuevoAlmacenFake() *almacenFake {
	return &almacenFake{
		vehiculos: map[int]models.Vehiculo{},

		nextProd: 1, nextCat: 1,
	}
}

func (a *almacenFake) ListarVehiculos() []models.Vehiculo {
	out := make([]models.Vehiculo, 0, len(a.vehiculos))
	for _, p := range a.vehiculos {
		out = append(out, p)
	}
	return out
}
func (a *almacenFake) BuscarVehiculoPorID(id int) (models.Vehiculo, bool) {
	v, ok := a.vehiculos[id]
	return v, ok
}
func (a *almacenFake) CrearVehiculo(v models.Vehiculo) models.Vehiculo {
	v.ID = a.nextProd
	a.nextProd++
	a.vehiculos[v.ID] = v
	return v
}
func (a *almacenFake) ActualizarVehiculo(id int, datos models.Vehiculo) (models.Vehiculo, bool) {
	if _, ok := a.vehiculos[id]; !ok {
		return models.Vehiculo{}, false
	}
	datos.ID = id
	a.vehiculos[id] = datos
	return datos, true
}
func (a *almacenFake) BorrarVehiculo(id int) bool {
	if _, ok := a.vehiculos[id]; !ok {
		return false
	}
	delete(a.vehiculos, id)
	return true
}

var _ storage.Almacen = (*almacenFake)(nil)

// usuarioFake implementa storage.UserRepository en memoria.
type usuarioFake struct {
	porEmail map[string]models.Usuario
	nextID   int
}

func nuevoUsuarioFake() *usuarioFake {
	return &usuarioFake{porEmail: map[string]models.Usuario{}, nextID: 1}
}
func (f *usuarioFake) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}
func (f *usuarioFake) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

var _ storage.UsuarioRepository = (*usuarioFake)(nil)

// =====================================================================
// Router de prueba: el MISMO que main.go (mismas rutas, mismo
// middleware.Auth real), pero con los dobles en memoria.
// =====================================================================

// construirEntorno devuelve el handler listo y un producto sembrado (id 1).
func construirEntorno() (http.Handler, *almacenFake, *usuarioFake) {
	almacen := nuevoAlmacenFake()
	almacen.CrearVehiculo(models.Vehiculo{
		ConductorID: 1, Placa: "MBT-456", Marca: "Chevrolet", Modelo: "Spark", Capacidad: 4})
	usuarios := nuevoUsuarioFake()

	srv := handlers.NewServer(handlers.Deps{
		Vehiculos: service.NuevoVehiculoService(almacen),
		Auth:      service.NuevoAuthService(usuarios),
	})
	authSvc := service.NuevoAuthService(usuarios)

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
	return r, almacen, usuarios
}

// jsonReq arma una peticion con cuerpo JSON y, si se pasa token, el header Bearer.
func jsonReq(metodo, ruta, cuerpo, token string) *http.Request {
	var body *strings.Reader
	if cuerpo == "" {
		body = strings.NewReader("")
	} else {
		body = strings.NewReader(cuerpo)
	}
	req := httptest.NewRequest(metodo, ruta, body)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return req
}

// tokenValido registra y loguea un usuario contra el router y devuelve su JWT.
func tokenValido(t *testing.T, h http.Handler) string {
	t.Helper()
	cred := `{"email":"docente@uleam.edu.ec","password":"secreta123"}`
	h.ServeHTTP(httptest.NewRecorder(), jsonReq(http.MethodPost, "/api/v1/auth/register", cred, ""))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, jsonReq(http.MethodPost, "/api/v1/auth/login", cred, ""))
	require.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Token string `json:"token"`
	}
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	require.NotEmpty(t, resp.Token)
	return resp.Token
}
