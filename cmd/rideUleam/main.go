// Command cafeteria-api arranca el servidor HTTP de la Cafetería Universitaria.
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite" // driver database/sql "sqlite" (pure-Go) para el backend sqlc
	"github.com/glebarez/sqlite"      // driver SQLite pure-Go (sin CGO)
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"RideUleam/internal/middleware"

	handlersVI "RideUleam/internal/handlers/viajeInmediato"
	modelsVI "RideUleam/internal/models/viajeInmediato"
	serviceVI "RideUleam/internal/service/viajeInmediato"
	storageVI "RideUleam/internal/storage/viajeInmediato"

	handlersUV "RideUleam/internal/handlers/usuarioVehiculo"
	modelsUV "RideUleam/internal/models/usuarioVehiculo"
	serviceUV "RideUleam/internal/service/usuarioVehiculo"
	storageUV "RideUleam/internal/storage/usuarioVehiculo"
)

func main() {
	// 1. Abrir SQLite y migrar el esquema (crea las tablas si no existen).
	gdb, err := gorm.Open(sqlite.Open("rideUleam.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}
	if err := gdb.AutoMigrate(
		&modelsVI.ViajeInmediato{},
		&modelsVI.SolicitudViaje{},
		&modelsVI.ParticipanteViaje{},
		&modelsUV.Usuario{},
		&modelsUV.Vehiculo{}); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Crear el almacenamiento GORM y sembrar si está vacío.
	almacenGormVI := storageVI.NuevoAlmacenSQLite(gdb)
	almacenGormUV := storageUV.NuevoAlmacenSQLite(gdb)

	almacenGormVI.SembrarSiVacio()
	almacenGormUV.SembrarSiVacio()

	var almacenVI storageVI.Almacen
	var almacenUV storageUV.Almacen

	switch os.Getenv("STORAGE") {
	case "sqlc":
		// Ya migramos y sembramos con GORM; cerramos esa conexión para que
		// sqlc sea el único dueño del archivo cafeteria.db en tiempo de servicio.
		if sqlDB, err := gdb.DB(); err == nil {
			_ = sqlDB.Close()
		}
		sdb, err := sql.Open("sqlite", "rideUleam.db")
		if err != nil {
			log.Fatal("no se pudo abrir sql.DB para sqlc: ", err)
		}
		almacenVI = storageVI.NuevoAlmacenSQLC(sdb)
		almacenUV = storageUV.NuevoAlmacenSQLC(sdb)
		log.Println("Backend de almacenamiento: sqlc (database/sql)")
	default:
		almacenVI = almacenGormVI
		almacenUV = almacenGormUV
		log.Println("Backend de almacenamiento: GORM")
	}

	// 3. Server con inyección de dependencias. No sabe qué backend recibió.

	usuarioRepo := storageUV.NuevoUsuarioGORM(gdb)
	authService := serviceUV.NuevoAuthService(usuarioRepo)
	vehiculoService := serviceUV.NuevoVehiculoService(almacenUV)

	viajeInmediatoService := serviceVI.NewViajeInmediatoService(almacenVI)
	solicitudViajeService := serviceVI.NewSolicitudViajeService(almacenVI)
	participanteService := serviceVI.NewParticipanteViajeService(almacenVI)

	servidorVI := handlersVI.NewServer(viajeInmediatoService, solicitudViajeService, participanteService)
	servidorUV := handlersUV.NewServer(vehiculoService, authService)

	// 4. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 5. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {

		r.Post("/auth/register", servidorUV.Registrar)
		r.Post("/auth/login", servidorUV.Login)
		// =========================
		// Viajes Inmediatos
		// =========================
		r.Group(func(r chi.Router) {

			r.Use(middleware.Auth(authService))

			r.Get("/viajes-inmediatos", servidorVI.ListarViajeInmediatos)
			r.Post("/viajes-inmediatos", servidorVI.CrearViajeInmediato)
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
			r.Delete("/vehiculos/{id}", servidorUV.BorrarVehiculo)
		})

	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
