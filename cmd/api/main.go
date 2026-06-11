package main

import (
	"fmt"
	"net/http"

	"RideUleam/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Route("/api/v1/estados", func(r chi.Router) {
		r.Post("/", handlers.CreateEstado)
		r.Get("/", handlers.GetEstados)
		r.Get("/{id}", handlers.GetEstado)
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
