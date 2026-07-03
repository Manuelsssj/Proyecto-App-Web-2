package service

import (
	"errors"

	models "cmd/rideUleam/internal/models/rutaProgramada"
	storage "cmd/rideUleam/internal/storage/rutaProgramada"
)

type RutaProgramadaService struct {
	almacen storage.Almacen
}

func NewRutaProgramadaService(almacen storage.Almacen) *RutaProgramadaService {
	return &RutaProgramadaService{
		almacen: almacen,
	}
}

// =====================
// VALIDACIONES
// =====================

func validarRutaProgramada(ruta models.RutaProgramada) error {
	if ruta.ConductorID <= 0 {
		return errors.New("El conductor_id es obligatorio")
	}

	if ruta.Origen == "" {
		return errors.New("El origen es obligatorio")
	}

	if ruta.Destino == "" {
		return errors.New("El destino es obligatorio")
	}

	if ruta.Costo < 0 {
		return errors.New("El costo no puede ser negativo")
	}

	return nil
}

func validarHorarioRuta(horario models.HorarioRuta) error {
	if horario.RutaID <= 0 {
		return errors.New("El ruta_id es obligatorio")
	}

	if horario.Dia == "" {
		return errors.New("El día es obligatorio")
	}

	if horario.Hora == "" {
		return errors.New("La hora es obligatoria")
	}

	return nil
}

func validarMantenimientoVehiculo(mantenimiento models.MantenimientoVehiculo) error {
	if mantenimiento.VehiculoID <= 0 {
		return errors.New("El vehiculo_id es obligatorio")
	}

	if mantenimiento.FechaInicio == "" {
		return errors.New("La fecha de inicio es obligatoria")
	}

	if mantenimiento.FechaFin == "" {
		return errors.New("La fecha de fin es obligatoria")
	}

	if mantenimiento.Motivo == "" {
		return errors.New("El motivo es obligatorio")
	}

	return nil
}

// =====================
// SERVICE RUTAS PROGRAMADAS
// =====================

func (s *RutaProgramadaService) ListarRutasProgramadas() []models.RutaProgramada {
	return s.almacen.ListarRutasProgramadas()
}

func (s *RutaProgramadaService) ObtenerRutaProgramada(id int) (models.RutaProgramada, error) {
	ruta, ok := s.almacen.BuscarRutaProgramadaPorID(id)
	if !ok {
		return models.RutaProgramada{}, ErrNoEncontrado
	}

	return ruta, nil
}

func (s *RutaProgramadaService) CrearRutaProgramada(ruta models.RutaProgramada) (models.RutaProgramada, error) {
	if err := validarRutaProgramada(ruta); err != nil {
		return models.RutaProgramada{}, err
	}

	creada := s.almacen.CrearRutaProgramada(ruta)

	return creada, nil
}

func (s *RutaProgramadaService) ActualizarRutaProgramada(id int, datos models.RutaProgramada) (models.RutaProgramada, error) {
	if err := validarRutaProgramada(datos); err != nil {
		return models.RutaProgramada{}, err
	}

	actualizada, ok := s.almacen.ActualizarRutaProgramada(id, datos)
	if !ok {
		return models.RutaProgramada{}, ErrNoEncontrado
	}

	return actualizada, nil
}

func (s *RutaProgramadaService) BorrarRutaProgramada(id int) error {
	ok := s.almacen.BorrarRutaProgramada(id)
	if !ok {
		return ErrNoEncontrado
	}

	return nil
}

func (s *RutaProgramadaService) ListarHorariosDeRutaProgramada(rutaID int) ([]models.HorarioRuta, error) {
	_, ok := s.almacen.BuscarRutaProgramadaPorID(rutaID)
	if !ok {
		return nil, ErrNoEncontrado
	}

	horarios := s.almacen.ListarHorariosPorRutaID(rutaID)

	return horarios, nil
}

func (s *RutaProgramadaService) ObtenerDetalleRutaProgramada(rutaID int) (models.RutaProgramada, []models.HorarioRuta, error) {
	ruta, ok := s.almacen.BuscarRutaProgramadaPorID(rutaID)
	if !ok {
		return models.RutaProgramada{}, nil, ErrNoEncontrado
	}

	horarios := s.almacen.ListarHorariosPorRutaID(rutaID)

	return ruta, horarios, nil
}

// =====================
// SERVICE HORARIOS DE RUTA
// =====================

func (s *RutaProgramadaService) ListarHorariosRuta() []models.HorarioRuta {
	return s.almacen.ListarHorariosRuta()
}

func (s *RutaProgramadaService) ObtenerHorarioRuta(id int) (models.HorarioRuta, error) {
	horario, ok := s.almacen.BuscarHorarioRutaPorID(id)
	if !ok {
		return models.HorarioRuta{}, ErrNoEncontrado
	}

	return horario, nil
}
func (s *RutaProgramadaService) CrearHorarioRuta(horario models.HorarioRuta) (models.HorarioRuta, error) {
	if err := validarHorarioRuta(horario); err != nil {
		return models.HorarioRuta{}, err
	}

	_, ok := s.almacen.BuscarRutaProgramadaPorID(horario.RutaID)
	if !ok {
		return models.HorarioRuta{}, errors.New("La ruta programada no existe")
	}

	creado := s.almacen.CrearHorarioRuta(horario)

	return creado, nil
}

func (s *RutaProgramadaService) ActualizarHorarioRuta(id int, datos models.HorarioRuta) (models.HorarioRuta, error) {
	if err := validarHorarioRuta(datos); err != nil {
		return models.HorarioRuta{}, err
	}

	_, ok := s.almacen.BuscarRutaProgramadaPorID(datos.RutaID)
	if !ok {
		return models.HorarioRuta{}, errors.New("La ruta programada no existe")
	}

	actualizado, ok := s.almacen.ActualizarHorarioRuta(id, datos)
	if !ok {
		return models.HorarioRuta{}, ErrNoEncontrado
	}

	return actualizado, nil
}

func (s *RutaProgramadaService) BorrarHorarioRuta(id int) error {
	ok := s.almacen.BorrarHorarioRuta(id)
	if !ok {
		return ErrNoEncontrado
	}

	return nil
}

// =====================
// SERVICE MANTENIMIENTOS VEHÍCULO
// =====================

func (s *RutaProgramadaService) ListarMantenimientosVehiculo() []models.MantenimientoVehiculo {
	return s.almacen.ListarMantenimientosVehiculo()
}

func (s *RutaProgramadaService) ObtenerMantenimientoVehiculo(id int) (models.MantenimientoVehiculo, error) {
	mantenimiento, ok := s.almacen.BuscarMantenimientoVehiculoPorID(id)
	if !ok {
		return models.MantenimientoVehiculo{}, ErrNoEncontrado
	}

	return mantenimiento, nil
}

func (s *RutaProgramadaService) CrearMantenimientoVehiculo(mantenimiento models.MantenimientoVehiculo) (models.MantenimientoVehiculo, error) {
	if err := validarMantenimientoVehiculo(mantenimiento); err != nil {
		return models.MantenimientoVehiculo{}, err
	}

	creado := s.almacen.CrearMantenimientoVehiculo(mantenimiento)

	return creado, nil
}

func (s *RutaProgramadaService) ActualizarMantenimientoVehiculo(id int, datos models.MantenimientoVehiculo) (models.MantenimientoVehiculo, error) {
	if err := validarMantenimientoVehiculo(datos); err != nil {
		return models.MantenimientoVehiculo{}, err
	}

	actualizado, ok := s.almacen.ActualizarMantenimientoVehiculo(id, datos)
	if !ok {
		return models.MantenimientoVehiculo{}, ErrNoEncontrado
	}

	return actualizado, nil
}

func (s *RutaProgramadaService) BorrarMantenimientoVehiculo(id int) error {
	ok := s.almacen.BorrarMantenimientoVehiculo(id)
	if !ok {
		return ErrNoEncontrado
	}

	return nil
}

func (s *RutaProgramadaService) ListarMantenimientosDeVehiculo(vehiculoID int) []models.MantenimientoVehiculo {
	return s.almacen.ListarMantenimientosPorVehiculoID(vehiculoID)
}
