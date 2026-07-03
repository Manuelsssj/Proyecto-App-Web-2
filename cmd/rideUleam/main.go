package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"cmd/rideUleam/internal/middleware"

	handlersRP "cmd/rideUleam/internal/handlers/rutaProgramada"
	modelsRP "cmd/rideUleam/internal/models/rutaProgramada"
	storageRP "cmd/rideUleam/internal/storage/rutaProgramada"
	usuarioStorage "cmd/rideUleam/internal/storage/usuario"

	handlersVI "cmd/rideUleam/internal/handlers/viajeInmediato"
	modelsVI "cmd/rideUleam/internal/models/viajeInmediato"
	serviceVI "cmd/rideUleam/internal/service/viajeInmediato"
	storageVI "cmd/rideUleam/internal/storage/viajeInmediato"

	handlersUV "cmd/rideUleam/internal/handlers/usuarioVehiculo"
	modelsUV "cmd/rideUleam/internal/models/usuarioVehiculo"
	serviceUV "cmd/rideUleam/internal/service/usuarioVehiculo"
	storageUV "cmd/rideUleam/internal/storage/usuarioVehiculo"
)

func main() {
	gdb, err := gorm.Open(sqlite.Open("rideUleam.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := gdb.AutoMigrate(
		&modelsVI.ViajeInmediato{},
		&modelsVI.SolicitudViaje{},
		&modelsVI.ParticipanteViaje{},
		&modelsUV.Usuario{},
		&modelsUV.Vehiculo{},
		&modelsRP.RutaProgramada{},
		&modelsRP.HorarioRuta{},
		&modelsRP.MantenimientoVehiculo{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	almacenVI := storageVI.NuevoAlmacenSQLite(gdb)
	almacenUV := storageUV.NuevoAlmacenSQLite(gdb)
	almacenRP := storageRP.NuevoAlmacenSQLite(gdb)

	almacenVI.SembrarSiVacio()
	almacenUV.SembrarSiVacio()
	almacenRP.SembrarSiVacio()

	log.Println("Backend de almacenamiento: GORM")

	usuarioRepoUV := storageUV.NuevoUsuarioGORM(gdb)
	authService := serviceUV.NuevoAuthService(usuarioRepoUV)
	vehiculoService := serviceUV.NuevoVehiculoService(almacenUV)

	viajeInmediatoService := serviceVI.NewViajeInmediatoService(almacenVI)
	solicitudViajeService := serviceVI.NewSolicitudViajeService(almacenVI)
	participanteService := serviceVI.NewParticipanteViajeService(almacenVI)

	usuarioRepoRP := usuarioStorage.NewUsuarioGORM(gdb)

	servidorVI := handlersVI.NewServer(viajeInmediatoService, solicitudViajeService, participanteService)
	servidorUV := handlersUV.NewServer(vehiculoService, authService)
	servidorRP := handlersRP.NewServer(almacenRP, usuarioRepoRP)

	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Servidor funcionando correctamente"}`))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidorUV.Registrar)
		r.Post("/auth/login", servidorUV.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			r.Get("/viajes-inmediatos", servidorVI.ListarViajeInmediatos)
			r.Post("/viajes-inmediatos", servidorVI.CrearViajeInmediato)
			r.Get("/viajes-inmediatos/{id}", servidorVI.ObtenerViajeInmediato)
			r.Put("/viajes-inmediatos/{id}", servidorVI.ActualizarViajeInmediato)
			r.Delete("/viajes-inmediatos/{id}", servidorVI.BorrarViajeInmediato)

			r.Get("/solicitudes-viajes", servidorVI.ListarSolicitudViajes)
			r.Post("/solicitudes-viajes", servidorVI.CrearSolicitudViaje)
			r.Get("/solicitudes-viajes/{id}", servidorVI.ObtenerSolicitudViaje)
			r.Put("/solicitudes-viajes/{id}", servidorVI.ActualizarSolicitudViaje)
			r.Delete("/solicitudes-viajes/{id}", servidorVI.BorrarSolicitudViaje)

			r.Get("/participantes-viajes", servidorVI.ListarParticipanteViajes)
			r.Post("/participantes-viajes", servidorVI.CrearParticipanteViaje)
			r.Get("/participantes-viajes/{id}", servidorVI.ObtenerParticipanteViaje)
			r.Put("/participantes-viajes/{id}", servidorVI.ActualizarParticipanteViaje)
			r.Delete("/participantes-viajes/{id}", servidorVI.BorrarParticipanteViaje)

			r.Get("/vehiculos", servidorUV.ListarVehiculos)
			r.Post("/vehiculos", servidorUV.CrearVehiculo)
			r.Get("/vehiculos/{id}", servidorUV.ObtenerVehiculo)
			r.Put("/vehiculos/{id}", servidorUV.ActualizarVehiculo)
			r.Delete("/vehiculos/{id}", servidorUV.BorrarVehiculo)

			r.Get("/rutas-programadas", servidorRP.ListarRutasProgramadas)
			r.Post("/rutas-programadas", servidorRP.CrearRutaProgramada)
			r.Get("/rutas-programadas/{id}", servidorRP.ObtenerRutaProgramada)
			r.Get("/rutas-programadas/{id}/horarios", servidorRP.ListarHorariosDeRutaProgramada)
			r.Get("/rutas-programadas/{id}/detalle", servidorRP.ObtenerDetalleRutaProgramada)
			r.Put("/rutas-programadas/{id}", servidorRP.ActualizarRutaProgramada)
			r.Delete("/rutas-programadas/{id}", servidorRP.BorrarRutaProgramada)

			r.Get("/horarios-ruta", servidorRP.ListarHorariosRuta)
			r.Post("/horarios-ruta", servidorRP.CrearHorarioRuta)
			r.Get("/horarios-ruta/{id}", servidorRP.ObtenerHorarioRuta)
			r.Put("/horarios-ruta/{id}", servidorRP.ActualizarHorarioRuta)
			r.Delete("/horarios-ruta/{id}", servidorRP.BorrarHorarioRuta)

			r.Get("/mantenimientos", servidorRP.ListarMantenimientosVehiculo)
			r.Post("/mantenimientos", servidorRP.CrearMantenimientoVehiculo)
			r.Get("/mantenimientos/{id}", servidorRP.ObtenerMantenimientoVehiculo)
			r.Put("/mantenimientos/{id}", servidorRP.ActualizarMantenimientoVehiculo)
			r.Delete("/mantenimientos/{id}", servidorRP.BorrarMantenimientoVehiculo)

			r.Get("/vehiculos/{vehiculoID}/mantenimientos", servidorRP.ListarMantenimientosDeVehiculo)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
