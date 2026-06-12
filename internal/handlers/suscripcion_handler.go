package handlers

import (
	"encoding/json"
	"net/http"

	"RideUleam/internal/models"
	"RideUleam/internal/storage"
)

func CreateSuscripcion(w http.ResponseWriter, r *http.Request) {

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

func GetSuscripciones(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.Suscripciones)
}
