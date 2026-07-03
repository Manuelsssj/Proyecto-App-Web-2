// Validaciones de productos
package service

import (
	models "cmd/rideUleam/internal/models/viajeInmediato"
	storage "cmd/rideUleam/internal/storage/viajeInmediato"
)

type ViajeInmediatoService struct {
	repo storage.ViajeInmediatoRepository
}

func NewViajeInmediatoService(repo storage.ViajeInmediatoRepository) *ViajeInmediatoService {
	return &ViajeInmediatoService{repo: repo}
}

func (s *ViajeInmediatoService) Listar() []models.ViajeInmediato {
	return s.repo.ListarViajeInmediatos()
}

func (s *ViajeInmediatoService) Obtener(id int) (models.ViajeInmediato, error) {
	vi, ok := s.repo.BuscarViajeInmediatoPorID(id)
	if !ok {
		return models.ViajeInmediato{}, ErrViajeInmediatoNoEncontrado
	}
	return vi, nil
}

func (s *ViajeInmediatoService) Crear(vi models.ViajeInmediato) (models.ViajeInmediato, error) {
	if err := validarViajeInmediato(vi); err != nil {
		return models.ViajeInmediato{}, err
	}
	return s.repo.CrearViajeInmediato(vi), nil
}

func (s *ViajeInmediatoService) Actualizar(id int, vi models.ViajeInmediato) (models.ViajeInmediato, error) {
	if err := validarViajeInmediato(vi); err != nil {
		return models.ViajeInmediato{}, err
	}
	vi, ok := s.repo.ActualizarViajeInmediato(id, vi)
	if !ok {
		return models.ViajeInmediato{}, ErrViajeInmediatoNoEncontrado
	}
	return vi, nil
}

func (s *ViajeInmediatoService) Borrar(id int) error {
	if !s.repo.BorrarViajeInmediato(id) {
		return ErrViajeInmediatoNoEncontrado
	}
	return nil
}

func validarViajeInmediato(vi models.ViajeInmediato) error {
	if vi.ConductorID <= 0 {
		return ErrConductorIDInvalido
	}

	return nil
}
