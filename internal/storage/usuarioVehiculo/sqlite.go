package storage

import (
	models "RideUleam/internal/models/usuarioVehiculo"

	"gorm.io/gorm"
)

// AlmacenSQLite implementa la interfaz Almacen usando GORM sobre SQLite.
//
// Fíjense: los métodos tienen EXACTAMENTE las mismas firmas que los de Memoria.
// Por eso el Server y los handlers no se enteran de cuál de los dos reciben.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// Módulo 4: Usuarios y Vehículos

// =========================================================
// vehiculos
// =========================================================

func (a *AlmacenSQLite) ListarVehiculos() []models.Vehiculo {
	var vehiculos []models.Vehiculo
	a.db.Find(&vehiculos)
	return vehiculos
}

func (a *AlmacenSQLite) BuscarVehiculoPorID(id int) (models.Vehiculo, bool) {
	var v models.Vehiculo
	if err := a.db.First(&v, id).Error; err != nil {
		return models.Vehiculo{}, false
	}
	return v, true
}

func (a *AlmacenSQLite) CrearVehiculo(v models.Vehiculo) models.Vehiculo {
	a.db.Create(&v)
	return v
}

func (a *AlmacenSQLite) ActualizarVehiculo(id int, datos models.Vehiculo) (models.Vehiculo, bool) {
	var existente models.Vehiculo
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.Vehiculo{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarVehiculo(id int) bool {
	res := a.db.Delete(&models.Vehiculo{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

// SembrarSiVacio inserta datos iniciales solo si aún no hay categorías.
// Así no duplicamos datos en cada arranque del servidor.
func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.Vehiculo{}).Count(&n)
	if n > 0 {
		return
	}

	vehiculos := []models.Vehiculo{
		{ID: 1, ConductorID: 1, Placa: "MBA-1234", Marca: "Toyota", Modelo: "Hiace", Capacidad: 30},
		{ID: 2, ConductorID: 2, Placa: "MBB-5678", Marca: "Hyundai", Modelo: "H1", Capacidad: 25},
		{ID: 3, ConductorID: 3, Placa: "MBC-9012", Marca: "Kia", Modelo: "Pregio", Capacidad: 20},
		{ID: 4, ConductorID: 1, Placa: "MBD-3456", Marca: "Mercedes", Modelo: "Sprinter", Capacidad: 35},
		{ID: 5, ConductorID: 2, Placa: "MBE-7890", Marca: "Chevrolet", Modelo: "N300", Capacidad: 15},
		{ID: 6, ConductorID: 3, Placa: "MBF-1122", Marca: "Renault", Modelo: "Master", Capacidad: 28},
	}
	a.db.Create(&vehiculos)

}

// Chequeo en tiempo de compilación: AlmacenSQLite debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLite)(nil)
