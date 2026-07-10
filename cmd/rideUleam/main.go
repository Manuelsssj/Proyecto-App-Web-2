// Command suscripciones-api arranca el servidor HTTP del sistema de suscripción de rutas.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	handlers "suscripciones-api/internal/handlers/suscripciones"
	"suscripciones-api/internal/middleware"
	models "suscripciones-api/internal/models/suscripciones"
	service "suscripciones-api/internal/service/suscripciones"
	storage "suscripciones-api/internal/storage/suscripciones"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=rideuleam port=5432 sslmode=disable"
	}

	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo conectar a PostgreSQL: ", err)
	}

	if err := gdb.AutoMigrate(
		&models.SuscripcionRuta{},
		&models.PlanPago{},
		&models.HistorialSuscripcion{},
		&models.Usuario{},
	); err != nil {
		log.Fatal("falló AutoMigrate: ", err)
	}
	if err := storage.SembrarDatosIniciales(gdb); err != nil {
		log.Fatal("falló el seeder: ", err)
	}

	almacen := storage.NuevoAlmacenGORM(gdb)
	log.Println("Backend de almacenamiento: GORM + PostgreSQL")

	usuarioRepo := storage.NewUsuarioGORM(gdb)
	authService := service.NewAuthService(usuarioRepo)
	suscripcionSrv := service.NewSuscripcionService(almacen)
	planSrv := service.NewPlanService(almacen)
	historialSrv := service.NewHistorialService(almacen)

	authH := handlers.NewAuthHandler(authService)
	suscripcionH := handlers.NewSuscripcionHandler(suscripcionSrv)
	planH := handlers.NewPlanHandler(planSrv)
	historialH := handlers.NewHistorialHandler(historialSrv)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

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

	puerto := os.Getenv("PORT")
	if puerto == "" {
		puerto = "8080"
	}

	log.Println("Servidor escuchando en http://localhost:" + puerto)
	log.Fatal(http.ListenAndServe(":"+puerto, r))

}
