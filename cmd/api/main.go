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

	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
