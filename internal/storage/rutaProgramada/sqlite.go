package storage

import (
	models "cmd/rideUleam/internal/models/rutaProgramada"

	"gorm.io/gorm"
)

// AlmacenSQLite implementa la interfaz Almacen usando GORM sobre SQLite.
//
// Los métodos tienen las mismas firmas que Memoria.
// Por eso los services y handlers no necesitan saber si reciben memoria o SQLite.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =====================
// RUTAS PROGRAMADAS
// =====================

func (a *AlmacenSQLite) ListarRutasProgramadas() []models.RutaProgramada {
	var rutas []models.RutaProgramada
	a.db.Find(&rutas)
	return rutas
}

func (a *AlmacenSQLite) BuscarRutaProgramadaPorID(id int) (models.RutaProgramada, bool) {
	var ruta models.RutaProgramada

	if err := a.db.First(&ruta, id).Error; err != nil {
		return models.RutaProgramada{}, false
	}

	return ruta, true
}

func (a *AlmacenSQLite) CrearRutaProgramada(ruta models.RutaProgramada) models.RutaProgramada {
	a.db.Create(&ruta)
	return ruta
}

func (a *AlmacenSQLite) ActualizarRutaProgramada(id int, datos models.RutaProgramada) (models.RutaProgramada, bool) {
	var existente models.RutaProgramada

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.RutaProgramada{}, false
	}

	datos.ID = id
	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarRutaProgramada(id int) bool {
	res := a.db.Delete(&models.RutaProgramada{}, id)
	return res.RowsAffected > 0
}

// =====================
// HORARIOS DE RUTA
// =====================

func (a *AlmacenSQLite) ListarHorariosRuta() []models.HorarioRuta {
	var horarios []models.HorarioRuta
	a.db.Find(&horarios)
	return horarios
}

func (a *AlmacenSQLite) ListarHorariosPorRutaID(rutaID int) []models.HorarioRuta {
	var horarios []models.HorarioRuta
	a.db.Where("ruta_id = ?", rutaID).Find(&horarios)
	return horarios
}

func (a *AlmacenSQLite) BuscarHorarioRutaPorID(id int) (models.HorarioRuta, bool) {
	var horario models.HorarioRuta

	if err := a.db.First(&horario, id).Error; err != nil {
		return models.HorarioRuta{}, false
	}

	return horario, true
}

func (a *AlmacenSQLite) CrearHorarioRuta(horario models.HorarioRuta) models.HorarioRuta {
	a.db.Create(&horario)
	return horario
}

func (a *AlmacenSQLite) ActualizarHorarioRuta(id int, datos models.HorarioRuta) (models.HorarioRuta, bool) {
	var existente models.HorarioRuta

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.HorarioRuta{}, false
	}

	datos.ID = id
	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarHorarioRuta(id int) bool {
	res := a.db.Delete(&models.HorarioRuta{}, id)
	return res.RowsAffected > 0
}

// =====================
// MANTENIMIENTOS VEHÍCULO
// =====================

func (a *AlmacenSQLite) ListarMantenimientosVehiculo() []models.MantenimientoVehiculo {
	var mantenimientos []models.MantenimientoVehiculo
	a.db.Find(&mantenimientos)
	return mantenimientos
}

func (a *AlmacenSQLite) ListarMantenimientosPorVehiculoID(vehiculoID int) []models.MantenimientoVehiculo {
	var mantenimientos []models.MantenimientoVehiculo
	a.db.Where("vehiculo_id = ?", vehiculoID).Find(&mantenimientos)
	return mantenimientos
}

func (a *AlmacenSQLite) BuscarMantenimientoVehiculoPorID(id int) (models.MantenimientoVehiculo, bool) {
	var mantenimiento models.MantenimientoVehiculo

	if err := a.db.First(&mantenimiento, id).Error; err != nil {
		return models.MantenimientoVehiculo{}, false
	}

	return mantenimiento, true
}

func (a *AlmacenSQLite) CrearMantenimientoVehiculo(mantenimiento models.MantenimientoVehiculo) models.MantenimientoVehiculo {
	a.db.Create(&mantenimiento)
	return mantenimiento
}

func (a *AlmacenSQLite) ActualizarMantenimientoVehiculo(id int, datos models.MantenimientoVehiculo) (models.MantenimientoVehiculo, bool) {
	var existente models.MantenimientoVehiculo

	if err := a.db.First(&existente, id).Error; err != nil {
		return models.MantenimientoVehiculo{}, false
	}

	datos.ID = id
	a.db.Save(&datos)

	return datos, true
}

func (a *AlmacenSQLite) BorrarMantenimientoVehiculo(id int) bool {
	res := a.db.Delete(&models.MantenimientoVehiculo{}, id)
	return res.RowsAffected > 0
}

// =====================
// SEEDS
// =====================

// SembrarSiVacio inserta datos iniciales solo si aún no hay rutas.
// Así no se duplican datos en cada arranque del servidor.
func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64

	a.db.Model(&models.RutaProgramada{}).Count(&n)
	if n > 0 {
		return
	}

	rutas := []models.RutaProgramada{
		{
			ID:          1,
			ConductorID: 1,
			Origen:      "Los Esteros",
			Destino:     "ULEAM",
			Costo:       0.75,
		},
		{
			ID:          2,
			ConductorID: 2,
			Origen:      "Tarqui",
			Destino:     "ULEAM",
			Costo:       1.00,
		},
	}
	a.db.Create(&rutas)

	horarios := []models.HorarioRuta{
		{
			ID:     1,
			RutaID: 1,
			Dia:    "Lunes",
			Hora:   "07:00",
		},
		{
			ID:     2,
			RutaID: 2,
			Dia:    "Martes",
			Hora:   "08:00",
		},
	}
	a.db.Create(&horarios)

	mantenimientos := []models.MantenimientoVehiculo{
		{
			ID:          1,
			VehiculoID:  1,
			FechaInicio: "2026-06-22",
			FechaFin:    "2026-06-25",
			Motivo:      "Cambio de aceite",
		},
	}
	a.db.Create(&mantenimientos)
}

// Chequeo en tiempo de compilación: AlmacenSQLite debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLite)(nil)
