// Package storage gestiona el almacenamiento en memoria de la cafetería.
//
// El tipo Memoria mantiene en un solo lugar todos los datos del dominio:
// Productos y Categorías.
package storage

import (
	models "RideUleam/internal/models/usuarioVehiculo"
	"sync"
)

// Memoria es un almacén unificado de la cafetería.
type Memoria struct {
	vehiculos      []models.Vehiculo
	nextVehiculoID int

	mu sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{

		vehiculos:      []models.Vehiculo{},
		nextVehiculoID: 1,
	}
}

// Módulo 4: Usuarios y Vehículos

// =========================================================
// usuarios
// =========================================================

// =========================================================
// vehiculos
// =========================================================

// SeedViajesInmediatos carga ViajesInmediatos iniciales en memoria.
func (m *Memoria) SeedVehiculos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.vehiculos = []models.Vehiculo{
		{ID: 1, ConductorID: 1, Placa: "MBA-1234", Marca: "Toyota", Modelo: "Hiace", Capacidad: 30},
		{ID: 2, ConductorID: 2, Placa: "MBB-5678", Marca: "Hyundai", Modelo: "H1", Capacidad: 25},
		{ID: 3, ConductorID: 3, Placa: "MBC-9012", Marca: "Kia", Modelo: "Pregio", Capacidad: 20},
		{ID: 4, ConductorID: 1, Placa: "MBD-3456", Marca: "Mercedes", Modelo: "Sprinter", Capacidad: 35},
		{ID: 5, ConductorID: 2, Placa: "MBE-7890", Marca: "Chevrolet", Modelo: "N300", Capacidad: 15},
		{ID: 6, ConductorID: 3, Placa: "MBF-1122", Marca: "Renault", Modelo: "Master", Capacidad: 28},
	}

	m.nextVehiculoID = 7

}

// ListarViajesInmediatos devuelve todos los productos en memoria.
func (m *Memoria) ListarVehiculos() []models.Vehiculo {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Vehiculo, len(m.vehiculos))
	copy(copia, m.vehiculos)
	return copia
}

// BuscarViajeInmediatoPorID devuelve el producto con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarVehiculoPorID(id int) (models.Vehiculo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, v := range m.vehiculos {
		if v.ID == id {
			return v, true
		}
	}
	return models.Vehiculo{}, false
}

// CrearViajeInmediato agrega un producto nuevo y devuelve el producto con ID asignado.
func (m *Memoria) CrearVehiculo(v models.Vehiculo) models.Vehiculo {
	m.mu.Lock()
	defer m.mu.Unlock()

	v.ID = m.nextVehiculoID
	m.nextVehiculoID++
	m.vehiculos = append(m.vehiculos, v)
	return v
}

// ActualizarViajeInmediato reemplaza el producto con el ID dado.
func (m *Memoria) ActualizarVehiculo(id int, datos models.Vehiculo) (models.Vehiculo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, v := range m.vehiculos {
		if v.ID == id {
			datos.ID = id
			m.vehiculos[i] = datos
			return datos, true
		}
	}
	return models.Vehiculo{}, false
}

// BorrarViajeInmediato elimina el producto con el ID dado
func (m *Memoria) BorrarVehiculo(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, sr := range m.vehiculos {
		if sr.ID == id {
			m.vehiculos = append(m.vehiculos[:i], m.vehiculos[i+1:]...)
			return true
		}
	}
	return false
}
