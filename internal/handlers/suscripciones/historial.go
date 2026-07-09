package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"suscripciones-api/internal/models/suscripciones"
	"suscripciones-api/internal/service/suscripciones"
)

// HistorialHandler maneja las rutas HTTP del historial de suscripciones.
type HistorialHandler struct {
	srv *service.HistorialService
}

// NewHistorialHandler crea un nuevo handler de historial.
func NewHistorialHandler(srv *service.HistorialService) *HistorialHandler {
	return &HistorialHandler{srv: srv}
}

// Registrar asocia las rutas al router.
func (h *HistorialHandler) Registrar(r chi.Router) {
	r.Get("/historial", h.Listar)
	r.Post("/historial", h.Crear)
	r.Get("/historial/{id}", h.Obtener)
	r.Put("/historial/{id}", h.Actualizar)
	r.Delete("/historial/{id}", h.Eliminar)
}

// Listar maneja GET /historial.
func (h *HistorialHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.srv.Listar()
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, lista)
}

// Crear maneja POST /historial.
func (h *HistorialHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var hist models.HistorialSuscripcion

	if err := json.NewDecoder(r.Body).Decode(&hist); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	creado, err := h.srv.Crear(hist)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

// Obtener maneja GET /historial/{id}.
func (h *HistorialHandler) Obtener(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerID(r)
	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	hist, err := h.srv.Obtener(id)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, hist)
}

// Actualizar maneja PUT /historial/{id}.
func (h *HistorialHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerID(r)
	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	var hist models.HistorialSuscripcion

	if err := json.NewDecoder(r.Body).Decode(&hist); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	hist.ID = uint(id)

	actualizado, err := h.srv.Actualizar(hist)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

// Eliminar maneja DELETE /historial/{id}.
func (h *HistorialHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerID(r)
	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	if err := h.srv.Eliminar(id); err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
