package handlers

import (
	"encoding/json"
	"net/http"

	"RideUleam/internal/models"
	"RideUleam/internal/storage"
)

func CreateEstado(w http.ResponseWriter, r *http.Request) {

	var estado models.Estado

	err := json.NewDecoder(r.Body).Decode(&estado)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	storage.Estados = append(storage.Estados, estado)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(estado)
}
