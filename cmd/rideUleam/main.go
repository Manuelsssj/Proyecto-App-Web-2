package main

import (
	"log"
	"net/http"

	"RideUleam/internal/config"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	handlers "RideUleam/internal/handlers/rutaProgramada"
	mw "RideUleam/internal/middleware"
	rutaModels "RideUleam/internal/models/rutaProgramada"
	usuarioModels "RideUleam/internal/models/usuario"
	rutaStorage "RideUleam/internal/storage/rutaProgramada"
	usuarioStorage "RideUleam/internal/storage/usuario"
)

func main() {
	// 1. GORM es el dueño del esquema: abre la base, migra y siembra.
	cfg := config.Cargar()

	gdb, err := rutaStorage.AbrirGorm(cfg.DBDriver, cfg.DBDsn, cfg.RutaDB)
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := gdb.AutoMigrate(
		&rutaModels.RutaProgramada{},
		&rutaModels.HorarioRuta{},
		&rutaModels.MantenimientoVehiculo{},
		&usuarioModels.Usuario{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	almacenGorm := rutaStorage.NuevoAlmacenSQLite(gdb)
	almacenGorm.SembrarSiVacio()

	// 2. Repositorio de usuario para Auth.
	usuarioRepo := usuarioStorage.NewUsuarioGORM(gdb)

	// 3. Server con inyección de dependencias.
	servidor := handlers.NewServer(almacenGorm, usuarioRepo)

	// 4. Router + middleware.
	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(mw.CORS)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Servidor funcionando correctamente"}`))
	})

	r.Route("/api/v1", func(r chi.Router) {
		// Auth - rutas públicas
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		// Rutas protegidas con JWT
		r.Group(func(r chi.Router) {
			r.Use(mw.AuthJWT(servidor.Auth))

			// Rutas Programadas
			r.Get("/rutas-programadas", servidor.ListarRutasProgramadas)
			r.Post("/rutas-programadas", servidor.CrearRutaProgramada)
			r.Get("/rutas-programadas/{id}", servidor.ObtenerRutaProgramada)
			r.Get("/rutas-programadas/{id}/horarios", servidor.ListarHorariosDeRutaProgramada)
			r.Get("/rutas-programadas/{id}/detalle", servidor.ObtenerDetalleRutaProgramada)
			r.Put("/rutas-programadas/{id}", servidor.ActualizarRutaProgramada)
			r.Delete("/rutas-programadas/{id}", servidor.BorrarRutaProgramada)

			// Horarios de Ruta
			r.Get("/horarios-ruta", servidor.ListarHorariosRuta)
			r.Post("/horarios-ruta", servidor.CrearHorarioRuta)
			r.Get("/horarios-ruta/{id}", servidor.ObtenerHorarioRuta)
			r.Put("/horarios-ruta/{id}", servidor.ActualizarHorarioRuta)
			r.Delete("/horarios-ruta/{id}", servidor.BorrarHorarioRuta)

			// Mantenimientos de Vehículo
			r.Get("/mantenimientos", servidor.ListarMantenimientosVehiculo)
			r.Post("/mantenimientos", servidor.CrearMantenimientoVehiculo)
			r.Get("/mantenimientos/{id}", servidor.ObtenerMantenimientoVehiculo)
			r.Put("/mantenimientos/{id}", servidor.ActualizarMantenimientoVehiculo)
			r.Delete("/mantenimientos/{id}", servidor.BorrarMantenimientoVehiculo)

			// Mantenimientos por Vehículo
			r.Get("/vehiculos/{vehiculoID}/mantenimientos", servidor.ListarMantenimientosDeVehiculo)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
