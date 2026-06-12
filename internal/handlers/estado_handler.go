package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"RideUleam/internal/models"

	"github.com/go-chi/chi/v5"
)

func (s *Server) CreateEstado(w http.ResponseWriter, r *http.Request) {
	var estado models.Estado

	err := json.NewDecoder(r.Body).Decode(&estado)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	creado := s.Storage.CrearEstado(estado)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}

func (s *Server) GetEstados(w http.ResponseWriter, r *http.Request) {
	estados := s.Storage.ListarEstados()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(estados)
}

func (s *Server) GetEstado(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	estado, encontrado := s.Storage.BuscarEstadoPorID(id)
	if !encontrado {
		http.Error(w, "Estado no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(estado)
}

func (s *Server) UpdateEstado(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var estadoActualizado models.Estado

	err = json.NewDecoder(r.Body).Decode(&estadoActualizado)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.Storage.ActualizarEstado(id, estadoActualizado)
	if !encontrado {
		http.Error(w, "Estado no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actualizado)
}

func (s *Server) DeleteEstado(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !s.Storage.BorrarEstado(id) {
		http.Error(w, "Estado no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
