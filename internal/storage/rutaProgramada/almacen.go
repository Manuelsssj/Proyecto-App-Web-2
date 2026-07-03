package storage

import models "cmd/rideUleam/internal/models/rutaProgramada"

// Almacen define QUÉ sabe hacer el almacenamiento del proyecto,
// sin decir CÓMO lo hace.
//
// Memoria y SQLite pueden cumplir esta interfaz.
// El service depende de esta interfaz y no de una implementación concreta.
type RutaProgramadaRepository interface {
	ListarRutasProgramadas() []models.RutaProgramada
	BuscarRutaProgramadaPorID(id int) (models.RutaProgramada, bool)
	CrearRutaProgramada(ruta models.RutaProgramada) models.RutaProgramada
	ActualizarRutaProgramada(id int, datos models.RutaProgramada) (models.RutaProgramada, bool)
	BorrarRutaProgramada(id int) bool
}

type HorarioRutaRepository interface {
	ListarHorariosRuta() []models.HorarioRuta
	ListarHorariosPorRutaID(rutaID int) []models.HorarioRuta
	BuscarHorarioRutaPorID(id int) (models.HorarioRuta, bool)
	CrearHorarioRuta(horario models.HorarioRuta) models.HorarioRuta
	ActualizarHorarioRuta(id int, datos models.HorarioRuta) (models.HorarioRuta, bool)
	BorrarHorarioRuta(id int) bool
}

type MantenimientoVehiculoRepository interface {
	ListarMantenimientosVehiculo() []models.MantenimientoVehiculo
	ListarMantenimientosPorVehiculoID(vehiculoID int) []models.MantenimientoVehiculo
	BuscarMantenimientoVehiculoPorID(id int) (models.MantenimientoVehiculo, bool)
	CrearMantenimientoVehiculo(mantenimiento models.MantenimientoVehiculo) models.MantenimientoVehiculo
	ActualizarMantenimientoVehiculo(id int, datos models.MantenimientoVehiculo) (models.MantenimientoVehiculo, bool)
	BorrarMantenimientoVehiculo(id int) bool
}

type Almacen interface {
	RutaProgramadaRepository
	HorarioRutaRepository
	MantenimientoVehiculoRepository
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto no compila.
var _ Almacen = (*Memoria)(nil)
