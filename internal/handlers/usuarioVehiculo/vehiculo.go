// Package handlers contiene los handlers HTTP de la API de cafetería.
package handlers

import (
	models "RideUleam/internal/models/usuarioVehiculo"
	"encoding/json"
	"net/http"
)

// ListarVehiculos atiende GET /api/v1/vehiculos.
func (s *Server) ListarVehiculos(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Vehiculos.Listar())
}

// ObtenerVehiculo atiende GET /api/v1/vehiculos/{id}.
func (s *Server) ObtenerVehiculo(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	vehiculo, err := s.Vehiculos.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, vehiculo)
}

// CrearVehiculo atiende POST /api/v1/vehiculos.
func (s *Server) CrearVehiculo(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Vehiculo
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	creado, err := s.Vehiculos.Crear(nuevo)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

// ActualizarVehiculo atiende PUT /api/v1/vehiculos/{id}.
func (s *Server) ActualizarVehiculo(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	var datos models.Vehiculo
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON invalido: "+err.Error())
		return
	}
	actualizado, err := s.Vehiculos.Actualizar(id, datos)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

// BorrarVehiculo atiende DELETE /api/v1/vehiculos/{id}.
func (s *Server) BorrarVehiculo(w http.ResponseWriter, r *http.Request) {
	id, err := idDeURL(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un numero entero")
		return
	}
	if err := s.Vehiculos.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusNoContent, nil)
}
