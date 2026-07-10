package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	handlers "RideUleam/internal/handlers/viajeInmediato"
	"RideUleam/internal/middleware"
	models "RideUleam/internal/models/viajeInmediato"
	service "RideUleam/internal/service/viajeInmediato"
	storage "RideUleam/internal/storage/viajeInmediato"

	handlersU "RideUleam/internal/handlers/usuarioVehiculo"
	modelsU "RideUleam/internal/models/usuarioVehiculo"
	serviceU "RideUleam/internal/service/usuarioVehiculo"
	storageU "RideUleam/internal/storage/usuarioVehiculo"
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
	viajeInmediatos    map[int]models.ViajeInmediato
	solicitudViajes    map[int]models.SolicitudViaje
	participanteViajes map[int]models.ParticipanteViaje
	nextProd           int
	nextCat            int
}

func nuevoAlmacenFake() *almacenFake {
	return &almacenFake{
		viajeInmediatos:    map[int]models.ViajeInmediato{},
		solicitudViajes:    map[int]models.SolicitudViaje{},
		participanteViajes: map[int]models.ParticipanteViaje{},
		nextProd:           1, nextCat: 1,
	}
}

func (a *almacenFake) ListarViajeInmediatos() []models.ViajeInmediato {
	out := make([]models.ViajeInmediato, 0, len(a.viajeInmediatos))
	for _, vi := range a.viajeInmediatos {
		out = append(out, vi)
	}
	return out
}
func (a *almacenFake) BuscarViajeInmediatoPorID(id int) (models.ViajeInmediato, bool) {
	vi, ok := a.viajeInmediatos[id]
	return vi, ok
}
func (a *almacenFake) CrearViajeInmediato(vi models.ViajeInmediato) models.ViajeInmediato {
	vi.ID = a.nextProd
	a.nextProd++
	a.viajeInmediatos[vi.ID] = vi
	return vi
}
func (a *almacenFake) ActualizarViajeInmediato(id int, datos models.ViajeInmediato) (models.ViajeInmediato, bool) {
	if _, ok := a.viajeInmediatos[id]; !ok {
		return models.ViajeInmediato{}, false
	}
	datos.ID = id
	a.viajeInmediatos[id] = datos
	return datos, true
}
func (a *almacenFake) BorrarViajeInmediato(id int) bool {
	if _, ok := a.viajeInmediatos[id]; !ok {
		return false
	}
	delete(a.viajeInmediatos, id)
	return true
}

func (a *almacenFake) ListarSolicitudViajes() []models.SolicitudViaje {
	out := make([]models.SolicitudViaje, 0, len(a.solicitudViajes))
	for _, sv := range a.solicitudViajes {
		out = append(out, sv)
	}
	return out
}
func (a *almacenFake) BuscarSolicitudViajePorID(id int) (models.SolicitudViaje, bool) {
	sv, ok := a.solicitudViajes[id]
	return sv, ok
}
func (a *almacenFake) CrearSolicitudViaje(sv models.SolicitudViaje) models.SolicitudViaje {
	sv.ID = a.nextCat
	a.nextCat++
	a.solicitudViajes[sv.ID] = sv
	return sv
}
func (a *almacenFake) ActualizarSolicitudViaje(id int, datos models.SolicitudViaje) (models.SolicitudViaje, bool) {
	if _, ok := a.solicitudViajes[id]; !ok {
		return models.SolicitudViaje{}, false
	}
	datos.ID = id
	a.solicitudViajes[id] = datos
	return datos, true
}
func (a *almacenFake) BorrarSolicitudViaje(id int) bool {
	if _, ok := a.solicitudViajes[id]; !ok {
		return false
	}
	delete(a.solicitudViajes, id)
	return true
}

func (a *almacenFake) ListarParticipanteViajes() []models.ParticipanteViaje {
	out := make([]models.ParticipanteViaje, 0, len(a.participanteViajes))
	for _, pv := range a.participanteViajes {
		out = append(out, pv)
	}
	return out
}
func (a *almacenFake) BuscarParticipanteViajePorID(id int) (models.ParticipanteViaje, bool) {
	pv, ok := a.participanteViajes[id]
	return pv, ok
}
func (a *almacenFake) CrearParticipanteViaje(pv models.ParticipanteViaje) models.ParticipanteViaje {
	pv.ID = a.nextCat
	a.nextCat++
	a.participanteViajes[pv.ID] = pv
	return pv
}
func (a *almacenFake) ActualizarParticipanteViaje(id int, datos models.ParticipanteViaje) (models.ParticipanteViaje, bool) {
	if _, ok := a.participanteViajes[id]; !ok {
		return models.ParticipanteViaje{}, false
	}
	datos.ID = id
	a.participanteViajes[id] = datos
	return datos, true
}
func (a *almacenFake) BorrarParticipanteViaje(id int) bool {
	if _, ok := a.participanteViajes[id]; !ok {
		return false
	}
	delete(a.participanteViajes, id)
	return true
}

var _ storage.Almacen = (*almacenFake)(nil)

// usuarioFake implementa storage.UserRepository en memoria.
type usuarioFake struct {
	porEmail map[string]modelsU.Usuario
	nextID   int
}

func nuevoUsuarioFake() *usuarioFake {
	return &usuarioFake{porEmail: map[string]modelsU.Usuario{}, nextID: 1}
}
func (f *usuarioFake) CrearUsuario(u modelsU.Usuario) (modelsU.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}
func (f *usuarioFake) BuscarUsuarioPorEmail(email string) (modelsU.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

var _ storageU.UsuarioRepository = (*usuarioFake)(nil)

// =====================================================================
// Router de prueba: el MISMO que main.go (mismas rutas, mismo
// middleware.Auth real), pero con los dobles en memoria.
// =====================================================================

// construirEntorno devuelve el handler listo y un producto sembrado (id 1).
func construirEntorno() (http.Handler, *almacenFake, *usuarioFake) {
	almacen := nuevoAlmacenFake()
	almacen.CrearViajeInmediato(models.ViajeInmediato{ConductorID: 1, Origen: "Los Esteros", Destino: "Universidad", HoraSalida: "07:30", Cupos: 4, Estado: "Disponible"})
	usuarios := nuevoUsuarioFake()
	srv := handlers.NewServer(handlers.Deps{
		ViajeInmediatos:    service.NewViajeInmediatoService(almacen),
		SolicitudViajes:    service.NewSolicitudViajeService(almacen),
		ParticipanteViajes: service.NewParticipanteViajeService(almacen),
	})

	srvU := handlersU.NewServer(handlersU.Deps{
		Auth: serviceU.NuevoAuthService(usuarios),
	})
	authSvc := serviceU.NuevoAuthService(usuarios)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", srvU.Registrar)
		r.Post("/auth/login", srvU.Login)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc)) // <- el middleware real de S10
			r.Get("/viajes-inmediatos", srv.ListarViajeInmediatos)
			r.Post("/viajes-inmediatos", srv.CrearViajeInmediato)
			r.Get("/viajes-inmediatos/{id}", srv.ObtenerViajeInmediato)
			r.Put("/viajes-inmediatos/{id}", srv.ActualizarViajeInmediato)
			r.Delete("/viajes-inmediatos/{id}", srv.BorrarViajeInmediato)
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
