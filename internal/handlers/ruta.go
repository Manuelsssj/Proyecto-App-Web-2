package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"RideUleam/internal/models"
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

// CrearRuta atiende POST /api/v1/categorias.
func (s *Server) CrearRuta(w http.ResponseWriter, r *http.Request) {
	var nueva models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.Sector) == "" {
		http.Error(w, "el campo sector es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.Destino) == "" {
		http.Error(w, "el campo destino es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nueva.HoraSalida) == "" {
		http.Error(w, "el campo hora de salida es obligatorio", http.StatusBadRequest)
		return
	}
	if nueva.Cupos <= 0 {
		http.Error(w, "el campo cupos es obligatorio y debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	creada := s.Storage.CrearRuta(nueva)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creada); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ActualizarRuta atiende PUT /api/v1/categorias/{id}.
func (s *Server) ActualizarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(datos.Sector) == "" {
		http.Error(w, "el campo sector es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Destino) == "" {
		http.Error(w, "el campo destino es obligatorio", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.HoraSalida) == "" {
		http.Error(w, "el campo hora de salida es obligatorio", http.StatusBadRequest)
		return
	}
	if datos.Cupos <= 0 {
		http.Error(w, "el campo cupos es obligatorio y debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	actualizada, encontrada := s.Storage.ActualizarRuta(id, datos)
	if !encontrada {
		http.Error(w, "ruta no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizada); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// BorrarRuta atiende DELETE /api/v1/categorias/{id}.
func (s *Server) BorrarRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	if !s.Storage.BorrarRuta(id) {
		http.Error(w, "categoría no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
