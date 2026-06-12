// Command cafeteria-api arranca el servidor HTTP de la Cafetería Universitaria.
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	chimw "github.com/go-chi/chi/v5/middleware"

	"RideUleam/internal/handlers"
	"RideUleam/internal/middleware"
	"RideUleam/internal/storage"
)

func main() {
	// 1. Crear el almacenamiento y cargar datos iniciales.
	almacen := storage.NuevaMemoria()
	almacen.SeedRutas()

	// 2. Crear el Server inyectándole el almacenamiento.
	servidor := handlers.NewServer(almacen)

	// 3. Configurar el router con versionado /api/v1/.
	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(middleware.CORS)

	r.Route("/api/v1", func(r chi.Router) {
		// Rutas: CRUD completo.
		r.Get("/rutas", servidor.ListarRutas)
		r.Post("/rutas", servidor.CrearRuta)
		r.Get("/rutas/{id}", servidor.ObtenerRutas)
		r.Put("/rutas/{id}", servidor.ActualizarRuta)
		r.Delete("/rutas/{id}", servidor.BorrarRuta)

		r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
			panic("panic de prueba")
		})
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
