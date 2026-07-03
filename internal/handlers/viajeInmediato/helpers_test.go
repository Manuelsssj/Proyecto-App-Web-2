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
	service "RideUleam/internal/service/viajeInmediato"
	storage "RideUleam/internal/storage/viajeInmediato"

	handlersU "RideUleam/internal/handlers/usuarioVehiculo"
	modelsU "RideUleam/internal/models/usuarioVehiculo"
	serviceU "RideUleam/internal/service/usuarioVehiculo"

	storageU "RideUleam/internal/storage/usuarioVehiculo"
)

// usuarioRepoFake: repositorio de usuarios en memoria para los tests de handler.
type usuarioRepoFake struct {
	porEmail map[string]modelsU.Usuario
	nextID   int
}

func nuevoUsuarioRepoFake() *usuarioRepoFake {
	return &usuarioRepoFake{porEmail: map[string]modelsU.Usuario{}, nextID: 1}
}

func (f *usuarioRepoFake) CrearUsuario(u modelsU.Usuario) (modelsU.Usuario, error) {
	u.ID = f.nextID
	f.nextID++
	f.porEmail[u.Email] = u
	return u, nil
}

func (f *usuarioRepoFake) BuscarUsuarioPorEmail(email string) (modelsU.Usuario, bool) {
	u, ok := f.porEmail[email]
	return u, ok
}

func construirEntorno(t *testing.T) (http.Handler, string) {
	t.Helper()

	almacenU := storageU.NuevaMemoria()

	almacen := storage.NuevaMemoria()
	almacen.SeedParticipanteViajes()
	usuarios := nuevoUsuarioRepoFake()

	viajeInmediatoViajeSvc := service.NewViajeInmediatoService(almacen)
	solicitudViajeSvc := service.NewSolicitudViajeService(almacen)
	participanteViajeSvc := service.NewParticipanteViajeService(almacen)

	vehiculoSvc := serviceU.NuevoVehiculoService(almacenU)

	authSvc := serviceU.NuevoAuthService(usuarios)
	srv := handlers.NewServer(viajeInmediatoViajeSvc, solicitudViajeSvc, participanteViajeSvc)
	srvU := handlersU.NewServer(vehiculoSvc, authSvc)

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

			// Solicitudes de Viaje

			r.Get("/solicitudes-viajes", srv.ListarSolicitudViajes)
			r.Post("/solicitudes-viajes", srv.CrearSolicitudViaje)
			r.Get("/solicitudes-viajes/{id}", srv.ObtenerSolicitudViaje)
			r.Put("/solicitudes-viajes/{id}", srv.ActualizarSolicitudViaje)
			r.Delete("/solicitudes-viajes/{id}", srv.BorrarSolicitudViaje)

			// Participantes de Viaje

			r.Get("/participantes-viajes", srv.ListarParticipanteViajes)
			r.Post("/participantes-viajes", srv.CrearParticipanteViaje)
			r.Get("/participantes-viajes/{id}", srv.ObtenerParticipanteViaje)
			r.Put("/participantes-viajes/{id}", srv.ActualizarParticipanteViaje)
			r.Delete("/participantes-viajes/{id}", srv.BorrarParticipanteViaje)
		})
	})

	token := registrarYObtenerToken(t, r)
	return r, token
}

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
