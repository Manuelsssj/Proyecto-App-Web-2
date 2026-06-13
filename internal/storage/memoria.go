package storage

import (
	"RideUleam/internal/models"
	"sync"
)

// Memoria es un almacén unificado de la cafetería.
type Memoria struct {
	rutas      []models.Ruta
	nextRutaID int

	estados      []models.Estado
	nextEstadoID int

	suscripciones     []models.Suscripcion
	nextSuscripcionID int

	mu sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{
		rutas:      []models.Ruta{},
		nextRutaID: 1,

		estados:      []models.Estado{},
		nextEstadoID: 1,

		suscripciones:     []models.Suscripcion{},
		nextSuscripcionID: 1,
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

// =========================================================
// Estados
// =========================================================

// SeedEstados carga rutas iniciales en memoria.

func (m *Memoria) SeedEstados() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.estados = []models.Estado{
		{
			ID: 1, RutaID: 1, Estado: "Activa", Motivo: "", FechaInicio: "2026-06-01", FechaFin: "",
		},
		{
			ID: 2, RutaID: 2, Estado: "Suspendida", Motivo: "Mantenimiento del vehículo", FechaInicio: "2026-06-05", FechaFin: "2026-06-10",
		},
		{
			ID: 3, RutaID: 3, Estado: "Activa", Motivo: "", FechaInicio: "2026-06-01", FechaFin: "",
		},
		{
			ID: 4, RutaID: 4, Estado: "Cancelada", Motivo: "Baja demanda", FechaInicio: "2026-06-08", FechaFin: "2026-06-08",
		},
		{
			ID: 5, RutaID: 5, Estado: "Activa", Motivo: "", FechaInicio: "2026-06-01", FechaFin: "",
		},
		{
			ID: 6, RutaID: 6, Estado: "Suspendida", Motivo: "Condiciones climáticas", FechaInicio: "2026-06-11", FechaFin: "2026-06-12",
		},
	}
	m.nextEstadoID = 7
}

func (m *Memoria) ListarEstados() []models.Estado {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Estado, len(m.estados))
	copy(copia, m.estados)
	return copia
}

// BuscarCategoriaPorID devuelve la categoría con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarEstadoPorID(id int) (models.Estado, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, e := range m.estados {
		if e.ID == id {
			return e, true
		}
	}
	return models.Estado{}, false
}

// CrearCategoria agrega una categoría nueva y devuelve la categoría con ID asignado.
func (m *Memoria) CrearEstado(e models.Estado) models.Estado {
	m.mu.Lock()
	defer m.mu.Unlock()

	e.ID = m.nextEstadoID
	m.nextEstadoID++
	m.estados = append(m.estados, e)
	return e
}

// ActualizarCategoria reemplaza la categoría con el ID dado.
func (m *Memoria) ActualizarEstado(id int, datos models.Estado) (models.Estado, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, e := range m.estados {
		if e.ID == id {
			datos.ID = id
			m.estados[i] = datos
			return datos, true
		}
	}
	return models.Estado{}, false
}

// BorrarCategoria elimina la categoría con el ID dado.
func (m *Memoria) BorrarEstado(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.estados {
		if c.ID == id {
			m.estados = append(m.estados[:i], m.estados[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// Suscripciones
// =========================================================

// SeedEstados carga rutas iniciales en memoria.

func (m *Memoria) SeedSuscripciones() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.suscripciones = []models.Suscripcion{
		{
			ID: 1, UsuarioID: 1, RutaID: 1, FechaInicio: "2026-06-01", Estado: "Activa",
		},
		{
			ID: 2, UsuarioID: 2, RutaID: 2, FechaInicio: "2026-06-03", Estado: "Activa",
		},
		{
			ID: 3, UsuarioID: 3, RutaID: 1, FechaInicio: "2026-06-05", Estado: "Pendiente",
		},
		{
			ID: 4, UsuarioID: 4, RutaID: 3, FechaInicio: "2026-06-07", Estado: "Cancelada",
		},
		{
			ID: 5, UsuarioID: 5, RutaID: 4, FechaInicio: "2026-06-08", Estado: "Activa",
		},
		{
			ID: 6, UsuarioID: 6, RutaID: 2, FechaInicio: "2026-06-10", Estado: "Activa",
		},
	}
	m.nextEstadoID = 7
}

func (m *Memoria) ListarSuscripciones() []models.Suscripcion {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Suscripcion, len(m.suscripciones))
	copy(copia, m.suscripciones)
	return copia
}

// BuscarCategoriaPorID devuelve la categoría con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarSuscripcionPorID(id int) (models.Suscripcion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, s := range m.suscripciones {
		if s.ID == id {
			return s, true
		}
	}
	return models.Suscripcion{}, false
}

// CrearCategoria agrega una categoría nueva y devuelve la categoría con ID asignado.
func (m *Memoria) CrearSuscripcion(s models.Suscripcion) models.Suscripcion {
	m.mu.Lock()
	defer m.mu.Unlock()

	s.ID = m.nextSuscripcionID
	m.nextSuscripcionID++
	m.suscripciones = append(m.suscripciones, s)
	return s
}

// ActualizarCategoria reemplaza la categoría con el ID dado.
func (m *Memoria) ActualizarSuscripcion(id int, datos models.Suscripcion) (models.Suscripcion, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, s := range m.suscripciones {
		if s.ID == id {
			datos.ID = id
			m.suscripciones[i] = datos
			return datos, true
		}
	}
	return models.Suscripcion{}, false
}

// BorrarCategoria elimina la categoría con el ID dado.
func (m *Memoria) BorrarSuscripcion(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, s := range m.estados {
		if s.ID == id {
			m.suscripciones = append(m.suscripciones[:i], m.suscripciones[i+1:]...)
			return true
		}
	}
	return false
}
