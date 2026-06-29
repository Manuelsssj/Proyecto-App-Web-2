// Command suscripciones-api arranca el servidor HTTP del sistema de suscripción de rutas.
package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"suscripciones-api/internal/handlers"
	"suscripciones-api/internal/middleware"
	"suscripciones-api/internal/models"
	"suscripciones-api/internal/service"
	"suscripciones-api/internal/storage"
)

func main() {
	// 1. GORM es el DUEÑO DEL ESQUEMA: abre la DB y migra las tablas.
	gdb, err := gorm.Open(sqlite.Open("suscripciones.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo abrir la base de datos: ", err)
	}

	if err := gdb.AutoMigrate(
		&models.SuscripcionRuta{},
		&models.PlanPago{},
		&models.HistorialSuscripcion{},
		&models.Usuario{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}

	// 2. Crear el backend de almacenamiento con GORM.
	almacen := storage.NuevoAlmacenSQLite(gdb)
	log.Println("Backend de almacenamiento: GORM")

	// 3. Servicios con inyección de dependencias.
	usuarioRepo := storage.NewUsuarioGORM(gdb)
	authService := service.NewAuthService(usuarioRepo)
	suscripcionSrv := service.NewSuscripcionService(almacen)
	planSrv := service.NewPlanService(almacen)
	historialSrv := service.NewHistorialService(almacen)

	// 4. Handlers.
	authH := handlers.NewAuthHandler(authService)
	suscripcionH := handlers.NewSuscripcionHandler(suscripcionSrv)
	planH := handlers.NewPlanHandler(planSrv)
	historialH := handlers.NewHistorialHandler(historialSrv)

	// 5. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	// 6. Rutas versionadas /api/v1/.
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", authH.Registrar)
		r.Post("/auth/login", authH.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(authService))

			r.Get("/suscripciones", suscripcionH.Listar)
			r.Post("/suscripciones", suscripcionH.Crear)
			r.Get("/suscripciones/{id}", suscripcionH.Obtener)
			r.Put("/suscripciones/{id}", suscripcionH.Actualizar)
			r.Delete("/suscripciones/{id}", suscripcionH.Eliminar)

			r.Get("/planes", planH.Listar)
			r.Post("/planes", planH.Crear)
			r.Get("/planes/{id}", planH.Obtener)
			r.Put("/planes/{id}", planH.Actualizar)
			r.Delete("/planes/{id}", planH.Eliminar)

			r.Get("/historial", historialH.Listar)
			r.Post("/historial", historialH.Crear)
			r.Get("/historial/{id}", historialH.Obtener)
			r.Put("/historial/{id}", historialH.Actualizar)
			r.Delete("/historial/{id}", historialH.Eliminar)
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
