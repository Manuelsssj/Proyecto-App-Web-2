package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	models "RideUleam/internal/models/viajeInmediato"
)

// ListarProductos atiende GET /api/v1/productos.
func (s *Server) ListarSolicitudViajes(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.SolicitudViajes.Listar())
}

// ObtenerProducto atiende GET /api/v1/productos/{id}.
func (s *Server) ObtenerSolicitudViaje(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	solicitudViaje, err := s.SolicitudViajes.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, solicitudViaje)
}

// CrearProducto atiende POST /api/v1/productos.
func (s *Server) CrearSolicitudViaje(w http.ResponseWriter, r *http.Request) {
	var nuevo models.SolicitudViaje
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	creado, err := s.SolicitudViajes.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarProducto atiende PUT /api/v1/productos/{id}.
func (s *Server) ActualizarSolicitudViaje(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	var datos models.SolicitudViaje
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	actualizado, err := s.SolicitudViajes.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarProducto atiende DELETE /api/v1/productos/{id}.
func (s *Server) BorrarSolicitudViaje(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	if err := s.SolicitudViajes.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
