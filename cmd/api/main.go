package main

import (
	"fmt"
	"net/http"

	"RideUleam/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api/v1/suscripciones", func(r chi.Router) {
		r.Post("/", handlers.CreateSuscripcion)
		r.Get("/", handlers.GetSuscripciones)
		r.Get("/{id}", handlers.GetSuscripcion)
		r.Put("/{id}", handlers.UpdateSuscripcion)
		r.Delete("/{id}", handlers.DeleteSuscripcion)
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
