package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"suscripciones-api/internal/service"
)

// responderJSON escribe v como JSON con el código de estado indicado.
func responderJSON(w http.ResponseWriter, codigo int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codigo)

	if v != nil {
		_ = json.NewEncoder(w).Encode(v)
	}
}

// responderError escribe un error en formato JSON.
func responderError(w http.ResponseWriter, codigo int, mensaje string) {
	responderJSON(w, codigo, map[string]string{
		"error": mensaje,
	})
}

// codigoDeError traduce un error de servicio al código HTTP correspondiente.
func codigoDeError(err error) int {
	switch {
	case errors.Is(err, service.ErrNoEncontrado):
		return http.StatusNotFound

	case errors.Is(err, service.ErrDatosInvalidos):
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}

// obtenerID obtiene el parámetro {id} de la URL usando Chi.
func obtenerID(r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(r, "id"))
}
