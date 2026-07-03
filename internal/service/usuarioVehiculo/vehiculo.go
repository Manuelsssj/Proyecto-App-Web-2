package service

import (
	models "cmd/rideUleam/internal/models/usuarioVehiculo"
	storage "cmd/rideUleam/internal/storage/usuarioVehiculo"
)

type VehiculoService struct {
	repo storage.VehiculoRepository
}

func NuevoVehiculoService(repo storage.VehiculoRepository) *VehiculoService {
	return &VehiculoService{repo: repo}
}

func (s *VehiculoService) Listar() []models.Vehiculo {
	return s.repo.ListarVehiculos()
}

func (s *VehiculoService) Obtener(id int) (models.Vehiculo, error) {
	v, ok := s.repo.BuscarVehiculoPorID(id)
	if !ok {
		return models.Vehiculo{}, ErrVehiculoNoEncontrado
	}
	return v, nil
}

func (s *VehiculoService) Crear(v models.Vehiculo) (models.Vehiculo, error) {
	if err := validarVehiculo(v); err != nil {
		return models.Vehiculo{}, err
	}
	return s.repo.CrearVehiculo(v), nil
}

func (s *VehiculoService) Actualizar(id int, datos models.Vehiculo) (models.Vehiculo, error) {
	if err := validarVehiculo(datos); err != nil {
		return models.Vehiculo{}, err
	}
	actualizado, ok := s.repo.ActualizarVehiculo(id, datos)
	if !ok {
		return models.Vehiculo{}, ErrVehiculoNoEncontrado
	}
	return actualizado, nil
}

func (s *VehiculoService) Borrar(id int) error {
	if !s.repo.BorrarVehiculo(id) {
		return ErrVehiculoNoEncontrado
	}
	return nil
}

// validarProducto centraliza las reglas de negocio que antes vivian en el handler.
func validarVehiculo(v models.Vehiculo) error {

	if v.ConductorID <= 0 {
		return ErrConductorIDInvalido
	}
	return nil
}
