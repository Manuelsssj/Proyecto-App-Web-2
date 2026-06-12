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

	server := handlers.NewServer()

	r.Route("/api/v1/estados", func(r chi.Router) {
		r.Post("/estado", server.CreateEstado)
		r.Get("/estados", server.GetEstados)
		r.Get("/estado/{id}", server.GetEstado)
		r.Put("/estado/{id}", server.UpdateEstado)
		r.Delete("/estado/{id}", server.DeleteEstado)
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
