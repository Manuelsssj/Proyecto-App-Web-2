// Command rideUleam arranca el servidor HTTP de la API RideULEAM.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"RideUleam/internal/config"
	"RideUleam/internal/httpserver"
	"RideUleam/internal/middleware"

	handlersVI "RideUleam/internal/handlers/viajeInmediato"
	serviceVI "RideUleam/internal/service/viajeInmediato"

	handlersUV "RideUleam/internal/handlers/usuarioVehiculo"
	serviceUV "RideUleam/internal/service/usuarioVehiculo"

	storage "RideUleam/internal/storage"
)

func main() {
	cfg := config.Cargar()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config.Config) error {
	// 1. Recursos de almacenamiento (Factory): abre DB (segun el motor elegido
	//    en la config: sqlite local o postgres en Docker), migra, siembra y elige backend.
	recursos, err := storage.Inicializar(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB, cfg.Backend)
	if err != nil {
		return err
	}
	defer func() { _ = recursos.Cerrar() }()
	log.Printf("Motor de base de datos: %s | Backend: %s", cfg.DBDriver, recursos.BackendUsado)
	// 2. Capa de servicio. AuthService recibe secreto y duracion por Options,
	//    tomados de la configuracion (antes eran globales hardcodeadas).
	viajesInmediatoSvc := serviceVI.NewViajeInmediatoService(recursos.AlmacenVI)
	solicitudViajeSvc := serviceVI.NewSolicitudViajeService(recursos.AlmacenVI)
	participanteViajeSvc := serviceVI.NewParticipanteViajeService(recursos.AlmacenVI)

	vehiculoSvc := serviceUV.NuevoVehiculoService(recursos.AlmacenUV)
	authSvc := serviceUV.NuevoAuthService(
		recursos.Usuarios,
		serviceUV.WithSecreto(cfg.JWTSecreto),
		serviceUV.WithDuracionToken(cfg.JWTDuracion),
	)

	// 3. Server con sus dependencias agrupadas en un struct (escala sin crecer
	//    la firma del constructor).
	servidorVI := handlersVI.NewServer(handlersVI.Deps{
		ViajeInmediatos:    viajesInmediatoSvc,
		SolicitudViajes:    solicitudViajeSvc,
		ParticipanteViajes: participanteViajeSvc,
	})

	servidorUV := handlersUV.NewServer(handlersUV.Deps{
		Vehiculos: vehiculoSvc,
		Auth:      authSvc,
	})

	// 4. Router + middleware global.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		// Publicas: registro y login.
		r.Post("/auth/register", servidorUV.Registrar)
		r.Post("/auth/login", servidorUV.Login)

		// Protegidas: exigen JWT valido en Authorization: Bearer <token>.
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authSvc))

			r.Get("/viajes-inmediatos", servidorVI.ListarViajeInmediatos)
			r.With(middleware.RequiereRol("admin", "conductor")).
				Post("/viajes-inmediatos", servidorVI.CrearViajeInmediato)
			r.Get("/viajes-inmediatos/{id}", servidorVI.ObtenerViajeInmediato)
			r.Put("/viajes-inmediatos/{id}", servidorVI.ActualizarViajeInmediato)
			r.Delete("/viajes-inmediatos/{id}", servidorVI.BorrarViajeInmediato)

			// Solicitudes de Viaje

			r.Get("/solicitudes-viajes", servidorVI.ListarSolicitudViajes)
			r.Post("/solicitudes-viajes", servidorVI.CrearSolicitudViaje)
			r.Get("/solicitudes-viajes/{id}", servidorVI.ObtenerSolicitudViaje)
			r.Put("/solicitudes-viajes/{id}", servidorVI.ActualizarSolicitudViaje)
			r.Delete("/solicitudes-viajes/{id}", servidorVI.BorrarSolicitudViaje)

			// Participantes de Viaje

			r.Get("/participantes-viajes", servidorVI.ListarParticipanteViajes)
			r.Post("/participantes-viajes", servidorVI.CrearParticipanteViaje)
			r.Get("/participantes-viajes/{id}", servidorVI.ObtenerParticipanteViaje)
			r.Put("/participantes-viajes/{id}", servidorVI.ActualizarParticipanteViaje)
			r.Delete("/participantes-viajes/{id}", servidorVI.BorrarParticipanteViaje)

			// =========================
			// Usuarios y Vehículos
			// =========================

			r.Get("/vehiculos", servidorUV.ListarVehiculos)
			r.Post("/vehiculos", servidorUV.CrearVehiculo)
			r.Get("/vehiculos/{id}", servidorUV.ObtenerVehiculo)
			r.Put("/vehiculos/{id}", servidorUV.ActualizarVehiculo)
			r.With(middleware.RequiereRol("admin")).
				Delete("/vehiculos/{id}", servidorUV.BorrarVehiculo)
		})
	})

	// 6. Servidor HTTP configurado por Options (puerto + timeouts desde config).
	srv := httpserver.Nuevo(
		r,
		httpserver.ConPuerto(cfg.Puerto),
		httpserver.ConReadTimeout(cfg.ReadTimeout),
		httpserver.ConWriteTimeout(cfg.WriteTimeout),
	)

	// 7. Contexto que se cancela al recibir Ctrl+C o SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 8. Arrancar el servidor en una goroutine para no bloquear la espera de la senal.
	errServidor := make(chan error, 1)
	go func() {
		log.Printf("Servidor escuchando en http://localhost%s", cfg.Puerto)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errServidor <- err
		}
	}()

	// 9. Esperar: o el servidor falla, o llega la senal de apagado.
	select {
	case err := <-errServidor:
		return err
	case <-ctx.Done():
		log.Println("Senal de apagado recibida, cerrando ordenadamente...")
	}

	// 10. Graceful shutdown: hasta 10s para terminar las requests en curso.
	ctxApagado, cancelar := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelar()
	if err := srv.Shutdown(ctxApagado); err != nil {
		return err
	}
	log.Println("Servidor detenido limpiamente.")
	return nil
}
