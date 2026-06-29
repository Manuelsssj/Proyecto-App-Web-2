// Package handlers contiene los handlers HTTP de la API de cafetería.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	models "RideUleam/internal/models/viajeInmediato"
)

// ListarProductos atiende GET /api/v1/productos.
func (s *Server) ListarViajeInmediatos(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.ViajeInmediatos.Listar())
}

// ObtenerProducto atiende GET /api/v1/productos/{id}.
func (s *Server) ObtenerViajeInmediato(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	viajeInmediato, err := s.ViajeInmediatos.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, viajeInmediato)
}

// CrearProducto atiende POST /api/v1/productos.
func (s *Server) CrearViajeInmediato(w http.ResponseWriter, r *http.Request) {
	var nuevo models.ViajeInmediato
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	creado, err := s.ViajeInmediatos.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarProducto atiende PUT /api/v1/productos/{id}.
func (s *Server) ActualizarViajeInmediato(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	var datos models.ViajeInmediato
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	actualizado, err := s.ViajeInmediatos.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarProducto atiende DELETE /api/v1/productos/{id}.
func (s *Server) BorrarViajeInmediato(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	if err := s.ViajeInmediatos.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
