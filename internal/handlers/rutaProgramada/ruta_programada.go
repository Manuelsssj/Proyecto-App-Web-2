package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	models "RideUleam/internal/models/rutaProgramada"
	service "RideUleam/internal/service/rutaProgramada"

	"github.com/go-chi/chi/v5"
)

// =====================
// RUTAS PROGRAMADAS
// =====================

func (s *Server) ListarRutasProgramadas(w http.ResponseWriter, r *http.Request) {
	rutas := s.RutaProgramadaService.ListarRutasProgramadas()

	responderJSON(w, http.StatusOK, rutas)
}

func (s *Server) ObtenerRutaProgramada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	ruta, err := s.RutaProgramadaService.ObtenerRutaProgramada(id)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Ruta programada no encontrada")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, ruta)
}

func (s *Server) CrearRutaProgramada(w http.ResponseWriter, r *http.Request) {
	var ruta models.RutaProgramada

	if errores := json.NewDecoder(r.Body).Decode(&ruta); err != nil {
		responderError(w, http.StatusBadRequest, "Datos inválidos")
		return
	}

	creada, err := s.RutaProgramadaService.CrearRutaProgramada(ruta)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarRutaProgramada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var datos models.RutaProgramada

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "Datos inválidos")
		return
	}

	actualizada, err := s.RutaProgramadaService.ActualizarRutaProgramada(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Ruta programada no encontrada")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, actualizada)
}

func (s *Server) BorrarRutaProgramada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	err = s.RutaProgramadaService.BorrarRutaProgramada(id)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Ruta programada no encontrada")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) ListarHorariosDeRutaProgramada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	horarios, err := s.RutaProgramadaService.ListarHorariosDeRutaProgramada(id)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Ruta programada no encontrada")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, horarios)
}

func (s *Server) ObtenerDetalleRutaProgramada(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	ruta, horarios, err := s.RutaProgramadaService.ObtenerDetalleRutaProgramada(id)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Ruta programada no encontrada")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	respuesta := struct {
		Ruta     models.RutaProgramada `json:"ruta"`
		Horarios []models.HorarioRuta  `json:"horarios"`
	}{
		Ruta:     ruta,
		Horarios: horarios,
	}

	responderJSON(w, http.StatusOK, respuesta)
}

// =====================
// HORARIOS DE RUTA
// =====================

func (s *Server) ListarHorariosRuta(w http.ResponseWriter, r *http.Request) {
	horarios := s.RutaProgramadaService.ListarHorariosRuta()

	responderJSON(w, http.StatusOK, horarios)
}

func (s *Server) ObtenerHorarioRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	horario, err := s.RutaProgramadaService.ObtenerHorarioRuta(id)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Horario de ruta no encontrado")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, horario)
}

func (s *Server) CrearHorarioRuta(w http.ResponseWriter, r *http.Request) {
	var horario models.HorarioRuta

	if err := json.NewDecoder(r.Body).Decode(&horario); err != nil {
		responderError(w, http.StatusBadRequest, "Datos inválidos")
		return
	}

	creado, err := s.RutaProgramadaService.CrearHorarioRuta(horario)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarHorarioRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var datos models.HorarioRuta

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "Datos inválidos")
		return
	}

	actualizado, err := s.RutaProgramadaService.ActualizarHorarioRuta(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Horario de ruta no encontrado")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarHorarioRuta(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	err = s.RutaProgramadaService.BorrarHorarioRuta(id)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Horario de ruta no encontrado")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// =====================
// MANTENIMIENTOS VEHÍCULO
// =====================

func (s *Server) ListarMantenimientosVehiculo(w http.ResponseWriter, r *http.Request) {
	mantenimientos := s.RutaProgramadaService.ListarMantenimientosVehiculo()

	responderJSON(w, http.StatusOK, mantenimientos)
}

func (s *Server) ObtenerMantenimientoVehiculo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	mantenimiento, err := s.RutaProgramadaService.ObtenerMantenimientoVehiculo(id)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Mantenimiento no encontrado")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, mantenimiento)
}

func (s *Server) CrearMantenimientoVehiculo(w http.ResponseWriter, r *http.Request) {
	var mantenimiento models.MantenimientoVehiculo

	if err := json.NewDecoder(r.Body).Decode(&mantenimiento); err != nil {
		responderError(w, http.StatusBadRequest, "Datos inválidos")
		return
	}

	creado, err := s.RutaProgramadaService.CrearMantenimientoVehiculo(mantenimiento)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarMantenimientoVehiculo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var datos models.MantenimientoVehiculo

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "Datos inválidos")
		return
	}

	actualizado, err := s.RutaProgramadaService.ActualizarMantenimientoVehiculo(id, datos)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Mantenimiento no encontrado")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarMantenimientoVehiculo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	err = s.RutaProgramadaService.BorrarMantenimientoVehiculo(id)
	if err != nil {
		if errors.Is(err, service.ErrNoEncontrado) {
			responderError(w, http.StatusNotFound, "Mantenimiento no encontrado")
			return
		}

		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) ListarMantenimientosDeVehiculo(w http.ResponseWriter, r *http.Request) {
	vehiculoID, err := strconv.Atoi(chi.URLParam(r, "vehiculoID"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "ID de vehículo inválido")
		return
	}

	mantenimientos := s.RutaProgramadaService.ListarMantenimientosDeVehiculo(vehiculoID)

	responderJSON(w, http.StatusOK, mantenimientos)
}
