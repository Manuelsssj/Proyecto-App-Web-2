package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"RideUleam/internal/models"

	"github.com/go-chi/chi/v5"
)

// ListarCategorias atiende GET /api/v1/categorias.
func (s *Server) ListarEstados(w http.ResponseWriter, _ *http.Request) {
	estados := s.Storage.ListarEstados()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(estados); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ObtenerCategoria atiende GET /api/v1/categorias/{id}.
func (s *Server) ObtenerEstado(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	estado, encontrado := s.Storage.BuscarEstadoPorID(id)
	if !encontrado {
		http.Error(w, "Estado no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(estado); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// CrearCategoria atiende POST /api/v1/categorias.
func (s *Server) CrearEstado(w http.ResponseWriter, r *http.Request) {
	var nueva models.Estado
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if nueva.RutaID <= 0 {
		http.Error(w, "el campo ruta_id es obligatorio", http.StatusBadRequest)
		return
	}

	creada := s.Storage.CrearEstado(nueva)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creada); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ActualizarCategoria atiende PUT /api/v1/categorias/{id}.
func (s *Server) ActualizarEstado(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Estado
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if datos.RutaID <= 0 {
		http.Error(w, "el campo ruta_id es obligatorio", http.StatusBadRequest)
		return
	}

	actualizada, encontrada := s.Storage.ActualizarEstado(id, datos)
	if !encontrada {
		http.Error(w, "categoría no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizada); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// BorrarCategoria atiende DELETE /api/v1/categorias/{id}.
func (s *Server) BorrarEstado(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	if !s.Storage.BorrarEstado(id) {
		http.Error(w, "categoría no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
