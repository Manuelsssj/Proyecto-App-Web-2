package storage

import (
	"RideUleam/internal/models"
	"sync"
)

// Memoria es un almacén unificado de la cafetería.
type Memoria struct {
	rutas      []models.Ruta
	nextRutaID int

	mu sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{
		rutas:      []models.Ruta{},
		nextRutaID: 1,
	}
}

// =========================================================
// RUTAS
// =========================================================

// SeedRutas carga rutas iniciales en memoria.
func (m *Memoria) SeedRutas() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.rutas = []models.Ruta{
		{
			ID: 1, Sector: "Los Esteros", Destino: "Terminal Terrestre", HoraSalida: "07:00", Cupos: 30, Estado: "Disponible", Conductor: "Juan Pérez",
		},
		{
			ID: 2, Sector: "Tarqui", Destino: "Universidad", HoraSalida: "08:00", Cupos: 25, Estado: "Disponible", Conductor: "María López",
		},
		{
			ID: 3, Sector: "Centro", Destino: "Hospital General", HoraSalida: "09:30", Cupos: 20, Estado: "Completa", Conductor: "Carlos Mendoza",
		},
		{
			ID: 4, Sector: "La Pradera", Destino: "Aeropuerto", HoraSalida: "11:00", Cupos: 35, Estado: "Disponible", Conductor: "Luis Zambrano",
		},
		{
			ID: 5, Sector: "Manta 2000", Destino: "Centro Comercial", HoraSalida: "14:00", Cupos: 28, Estado: "Disponible", Conductor: "Ana García",
		},
		{
			ID: 6, Sector: "Jocay", Destino: "Terminal Terrestre", HoraSalida: "16:30", Cupos: 22, Estado: "En ruta", Conductor: "Pedro Cedeño",
		},
	}
	m.nextRutaID = 7
}

// ListarRutas devuelve todos los productos en memoria.
func (m *Memoria) ListarRutas() []models.Ruta {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Ruta, len(m.rutas))
	copy(copia, m.rutas)
	return copia
}

// BuscarRutaPorID devuelve el producto con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarRutaPorID(id int) (models.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, rt := range m.rutas {
		if rt.ID == id {
			return rt, true
		}
	}
	return models.Ruta{}, false
}

// CrearRuta agrega un producto nuevo y devuelve el producto con ID asignado.
func (m *Memoria) CrearRuta(rt models.Ruta) models.Ruta {
	m.mu.Lock()
	defer m.mu.Unlock()

	rt.ID = m.nextRutaID
	m.nextRutaID++
	m.rutas = append(m.rutas, rt)
	return rt
}

// ActualizarRuta reemplaza el producto con el ID dado.
func (m *Memoria) ActualizarRuta(id int, datos models.Ruta) (models.Ruta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, rt := range m.rutas {
		if rt.ID == id {
			datos.ID = id
			m.rutas[i] = datos
			return datos, true
		}
	}
	return models.Ruta{}, false
}

// BorrarRuta elimina el producto con el ID dado
func (m *Memoria) BorrarRuta(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, rt := range m.rutas {
		if rt.ID == id {
			m.rutas = append(m.rutas[:i], m.rutas[i+1:]...)
			return true
		}
	}
	return false
}
