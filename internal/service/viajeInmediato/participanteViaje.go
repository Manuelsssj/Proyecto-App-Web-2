package service

import (
	models "cmd/rideUleam/internal/models/viajeInmediato"
	storage "cmd/rideUleam/internal/storage/viajeInmediato"
)

type ParticipanteViajeService struct {
	repo storage.ParticipanteViajeRepository
}

func NewParticipanteViajeService(repo storage.ParticipanteViajeRepository) *ParticipanteViajeService {
	return &ParticipanteViajeService{repo: repo}
}

func (s *ParticipanteViajeService) Listar() []models.ParticipanteViaje {
	return s.repo.ListarParticipanteViajes()
}

func (s *ParticipanteViajeService) Obtener(id int) (models.ParticipanteViaje, error) {
	pv, ok := s.repo.BuscarParticipanteViajePorID(id)
	if !ok {
		return models.ParticipanteViaje{}, ErrParticipanteViajeNoEncontrado
	}
	return pv, nil
}

func (s *ParticipanteViajeService) Crear(pv models.ParticipanteViaje) (models.ParticipanteViaje, error) {
	if err := validarParticipanteViaje(pv); err != nil {
		return models.ParticipanteViaje{}, err
	}

	return s.repo.CrearParticipanteViaje(pv), nil
}

func (s *ParticipanteViajeService) Actualizar(id int, pv models.ParticipanteViaje) (models.ParticipanteViaje, error) {
	if err := validarParticipanteViaje(pv); err != nil {
		return models.ParticipanteViaje{}, err
	}

	pv, ok := s.repo.ActualizarParticipanteViaje(id, pv)
	if !ok {
		return models.ParticipanteViaje{}, ErrParticipanteViajeNoEncontrado
	}

	return pv, nil
}

func (s *ParticipanteViajeService) Borrar(id int) error {
	if !s.repo.BorrarParticipanteViaje(id) {
		return ErrParticipanteViajeNoEncontrado
	}

	return nil
}

func validarParticipanteViaje(pv models.ParticipanteViaje) error {
	if pv.ViajeID == 0 {
		return ErrViajeIDInvalido
	}

	if pv.UsuarioID == 0 {
		return ErrUsuarioIDInvalido
	}

	return nil
}
