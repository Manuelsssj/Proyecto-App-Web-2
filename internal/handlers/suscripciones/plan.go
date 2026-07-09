package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"suscripciones-api/internal/models/suscripciones"
	"suscripciones-api/internal/service/suscripciones"
)

// PlanHandler maneja las rutas HTTP de los planes de pago.
type PlanHandler struct {
	srv *service.PlanService
}

// NewPlanHandler crea un nuevo handler de planes de pago.
func NewPlanHandler(srv *service.PlanService) *PlanHandler {
	return &PlanHandler{srv: srv}
}

// Registrar asocia las rutas al router.
func (h *PlanHandler) Registrar(r chi.Router) {
	r.Get("/planes", h.Listar)
	r.Post("/planes", h.Crear)
	r.Get("/planes/{id}", h.Obtener)
	r.Put("/planes/{id}", h.Actualizar)
	r.Delete("/planes/{id}", h.Eliminar)
}

// Listar maneja GET /planes.
func (h *PlanHandler) Listar(w http.ResponseWriter, r *http.Request) {
	lista, err := h.srv.Listar()
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, lista)
}

// Crear maneja POST /planes.
func (h *PlanHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var p models.PlanPago

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	creado, err := h.srv.Crear(p)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

// Obtener maneja GET /planes/{id}.
func (h *PlanHandler) Obtener(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerID(r)
	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	p, err := h.srv.Obtener(id)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, p)
}

// Actualizar maneja PUT /planes/{id}.
func (h *PlanHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := obtenerID(r)
	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	var p models.PlanPago

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	p.ID = uint(id)

	actualizado, err := h.srv.Actualizar(p)
	if err != nil {
		responderError(w, codigoDeError(err), err.Error())
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

// Eliminar maneja DELETE /planes/{id}.
func (h *PlanHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
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
