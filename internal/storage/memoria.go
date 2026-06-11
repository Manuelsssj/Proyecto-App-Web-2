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
