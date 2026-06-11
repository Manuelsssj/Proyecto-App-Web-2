package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"RideUleam/internal/models"
	"RideUleam/internal/storage"

	"github.com/go-chi/chi/v5"
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

func GetEstados(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.Estados)
}

func GetEstado(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, estado := range storage.Estados {
		if estado.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(estado)
			return
		}
	}

	http.Error(w, "Estado no encontrado", http.StatusNotFound)
}

func UpdateEstado(w http.ResponseWriter, r *http.Request) {

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

	for i, estado := range storage.Estados {

		if estado.ID == id {

			storage.Estados[i] = estadoActualizado

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(estadoActualizado)
			return
		}
	}

	http.Error(w, "Estado no encontrado", http.StatusNotFound)
}
