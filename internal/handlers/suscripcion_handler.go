package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"RideUleam/internal/models"
	"RideUleam/internal/storage"

	"github.com/go-chi/chi/v5"
)

func (s *Server) CreateSuscripcion(w http.ResponseWriter, r *http.Request) {

	var suscripcion models.Suscripcion

	err := json.NewDecoder(r.Body).Decode(&suscripcion)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	storage.Suscripciones = append(storage.Suscripciones, suscripcion)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(suscripcion)
}

func (s *Server) GetSuscripciones(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.Suscripciones)
}
func (s *Server) GetSuscripcion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, suscripcion := range storage.Suscripciones {
		if suscripcion.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(suscripcion)
			return
		}
	}

	http.Error(w, "Suscripción no encontrada", http.StatusNotFound)
}
func (s *Server) UpdateSuscripcion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var suscripcionActualizada models.Suscripcion

	err = json.NewDecoder(r.Body).Decode(&suscripcionActualizada)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	for i, suscripcion := range storage.Suscripciones {
		if suscripcion.ID == id {
			suscripcionActualizada.ID = id
			storage.Suscripciones[i] = suscripcionActualizada

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(suscripcionActualizada)
			return
		}
	}

	http.Error(w, "Suscripción no encontrada", http.StatusNotFound)
}
func (s *Server) DeleteSuscripcion(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, suscripcion := range storage.Suscripciones {
		if suscripcion.ID == id {

			storage.Suscripciones = append(
				storage.Suscripciones[:i],
				storage.Suscripciones[i+1:]...,
			)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Suscripción no encontrada", http.StatusNotFound)
}
