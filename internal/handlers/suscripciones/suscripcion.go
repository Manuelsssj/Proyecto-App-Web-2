package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"RideUleam/internal/models/suscripciones"
	"RideUleam/internal/service/suscripciones"
)

// SuscripcionHandler maneja las rutas HTTP de las suscripciones.
type SuscripcionHandler struct {
	srv *service.SuscripcionService
}

// NewSuscripcionHandler crea un nuevo handler de suscripciones.
func NewSuscripcionHandler(srv *service.SuscripcionService) *SuscripcionHandler {
	return &SuscripcionHandler{srv: srv}
}

// Registrar asocia las rutas al router.
func (h *SuscripcionHandler) Registrar(r chi.Router) {
	r.Get("/suscripciones", h.Listar)
	r.Post("/suscripciones", h.Crear)
	r.Get("/suscripciones/{id}", h.Obtener)
	r.Put("/suscripciones/{id}", h.Actualizar)
	r.Delete("/suscripciones/{id}", h.Eliminar)
}

// Listar maneja GET /suscripciones.
func (h *SuscripcionHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.srv.Listar()
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, lista)
}

// Crear maneja POST /suscripciones.
func (h *SuscripcionHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var s models.SuscripcionRuta

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	creada, err := h.srv.Crear(s)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creada)
}

// Obtener maneja GET /suscripciones/{id}.
func (h *SuscripcionHandler) Obtener(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerID(r)
	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	s, err := h.srv.Obtener(id)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, s)
}

// Actualizar maneja PUT /suscripciones/{id}.
func (h *SuscripcionHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerID(r)
	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	var s models.SuscripcionRuta

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	s.ID = uint(id)

	actualizada, err := h.srv.Actualizar(s)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, actualizada)
}

// Eliminar maneja DELETE /suscripciones/{id}.
func (h *SuscripcionHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
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
