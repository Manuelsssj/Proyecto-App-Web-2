package main

import (
	"log"
	"net/http"

	"RideUleam/internal/handlers"
	"RideUleam/internal/middleware"
	"RideUleam/internal/storage"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func main() {
	almacen := storage.NuevaMemoria()
	almacen.SeedRutas()

	servidor := handlers.NewServer(almacen)

	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		// Módulo Rutas
		r.Get("/rutas", servidor.ListarRutas)
		r.Post("/rutas", servidor.CrearRuta)
		r.Get("/rutas/{id}", servidor.ObtenerRutas)
		r.Put("/rutas/{id}", servidor.ActualizarRuta)
		r.Delete("/rutas/{id}", servidor.BorrarRuta)

		// Módulo Estado
		r.Post("/estados/estado", servidor.CreateEstado)
		r.Get("/estados/estados", servidor.GetEstados)
		r.Get("/estados/estado/{id}", servidor.GetEstado)
		r.Put("/estados/estado/{id}", servidor.UpdateEstado)
		r.Delete("/estados/estado/{id}", servidor.DeleteEstado)

		// Módulo Suscripciones
		r.Post("/suscripciones", servidor.CreateSuscripcion)
		r.Get("/suscripciones", servidor.GetSuscripciones)
		r.Get("/suscripciones/{id}", servidor.GetSuscripcion)
		r.Put("/suscripciones/{id}", servidor.UpdateSuscripcion)
		r.Delete("/suscripciones/{id}", servidor.DeleteSuscripcion)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
