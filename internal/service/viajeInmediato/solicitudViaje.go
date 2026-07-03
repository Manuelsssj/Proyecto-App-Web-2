// Validaciones de productos
package service

import (
	models "cmd/rideUleam/internal/models/viajeInmediato"
	storage "cmd/rideUleam/internal/storage/viajeInmediato"
)

type SolicitudViajeService struct {
	repo storage.SolicitudViajeRepository
}

func NewSolicitudViajeService(repo storage.SolicitudViajeRepository) *SolicitudViajeService {
	return &SolicitudViajeService{repo: repo}
}

func (s *SolicitudViajeService) Listar() []models.SolicitudViaje {
	return s.repo.ListarSolicitudViajes()
}

func (s *SolicitudViajeService) Obtener(id int) (models.SolicitudViaje, error) {
	sv, ok := s.repo.BuscarSolicitudViajePorID(id)
	if !ok {
		return models.SolicitudViaje{}, ErrSolicitudViajeNoEncontrado
	}
	return sv, nil
}

func (s *SolicitudViajeService) Crear(sv models.SolicitudViaje) (models.SolicitudViaje, error) {
	if err := validarSolicitudViaje(sv); err != nil {
		return models.SolicitudViaje{}, err
	}
	return s.repo.CrearSolicitudViaje(sv), nil
}

func (s *SolicitudViajeService) Actualizar(id int, sv models.SolicitudViaje) (models.SolicitudViaje, error) {
	if err := validarSolicitudViaje(sv); err != nil {
		return models.SolicitudViaje{}, err
	}
	sv, ok := s.repo.ActualizarSolicitudViaje(id, sv)
	if !ok {
		return models.SolicitudViaje{}, ErrSolicitudViajeNoEncontrado
	}
	return sv, nil
}

func (s *SolicitudViajeService) Borrar(id int) error {
	if !s.repo.BorrarSolicitudViaje(id) {
		return ErrSolicitudViajeNoEncontrado
	}
	return nil
}

func validarSolicitudViaje(sv models.SolicitudViaje) error {
	if sv.ViajeID == 0 {
		return ErrViajeIDInvalido
	}

	if sv.PasajeroID == 0 {
		return ErrPasajeroIDInvalido
	}
	return nil
}
