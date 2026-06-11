package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"RideUleam/internal/storage"
)

// Server agrupa las dependencias compartidas por los handlers.
// Recibe el storage por inyección de dependencias desde main.
type Server struct {
	Storage *storage.Memoria
}

// NewServer construye un Server listo para usar.
func NewServer(s *storage.Memoria) *Server {
	return &Server{Storage: s}
}

// ListarRutas atiende GET /api/v1/categorias.
func (s *Server) ListarRutas(w http.ResponseWriter, _ *http.Request) {
	rutas := s.Storage.ListarRutas()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(rutas); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ObtenerRutas atiende GET /api/v1/rutas/{id}.
func (s *Server) ObtenerRutas(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	ruta, encontrado := s.Storage.BuscarRutaPorID(id)
	if !encontrado {
		http.Error(w, "ruta no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ruta); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}
