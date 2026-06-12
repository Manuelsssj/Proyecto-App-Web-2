package main

import (
	"fmt"
	"net/http"

	"RideUleam/internal/handlers"
	"RideUleam/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.CORS)

	r.Route("/api/v1/estados", func(r chi.Router) {
		r.Post("/estado", handlers.CreateEstado)
		r.Get("/estados", handlers.GetEstados)
		r.Get("/estado/{id}", handlers.GetEstado)
		r.Put("/estado/{id}", handlers.UpdateEstado)
		r.Delete("/estado/{id}", handlers.DeleteEstado)
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
