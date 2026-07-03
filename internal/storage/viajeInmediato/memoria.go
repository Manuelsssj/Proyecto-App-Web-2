// Package storage gestiona el almacenamiento en memoria de la cafetería.
//
// El tipo Memoria mantiene en un solo lugar todos los datos del dominio:
// Productos y Categorías.
package storage

import (
	models "cmd/rideUleam/internal/models/viajeInmediato"
	"sync"
)

// Memoria es un almacén unificado de la cafetería.
type Memoria struct {
	viajeInmediatos      []models.ViajeInmediato
	nextViajeInmediatoID int

	solicitudViajes      []models.SolicitudViaje
	nextSolicitudViajeID int

	participanteViajes      []models.ParticipanteViaje
	nextParticipanteViajeID int

	mu sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{
		viajeInmediatos:      []models.ViajeInmediato{},
		nextViajeInmediatoID: 1,

		solicitudViajes:      []models.SolicitudViaje{},
		nextSolicitudViajeID: 1,

		participanteViajes:      []models.ParticipanteViaje{},
		nextParticipanteViajeID: 1,
	}
}

// =========================================================
// ViajeInmediato
// =========================================================

// SeedProductos carga productos iniciales en memoria.
func (m *Memoria) SeedViajeInmediatos() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.viajeInmediatos = []models.ViajeInmediato{
		{ID: 1, ConductorID: 1, Origen: "Los Esteros", Destino: "Terminal Terrestre", HoraSalida: "07:00", Cupos: 30, Estado: "Disponible"},
		{ID: 2, ConductorID: 2, Origen: "Tarqui", Destino: "Universidad", HoraSalida: "08:00", Cupos: 25, Estado: "Disponible"},
		{ID: 3, ConductorID: 3, Origen: "Centro", Destino: "Hospital General", HoraSalida: "09:30", Cupos: 20, Estado: "Completa"},
		{ID: 4, ConductorID: 4, Origen: "La Pradera", Destino: "Aeropuerto", HoraSalida: "11:00", Cupos: 35, Estado: "Disponible"},
		{ID: 5, ConductorID: 5, Origen: "Manta 2000", Destino: "Centro Comercial", HoraSalida: "14:00", Cupos: 28, Estado: "Disponible"},
		{ID: 6, ConductorID: 6, Origen: "Jocay", Destino: "Terminal Terrestre", HoraSalida: "16:30", Cupos: 22, Estado: "En ruta"},
	}
	m.nextViajeInmediatoID = 7
}

// ListarProductos devuelve todos los productos en memoria.
func (m *Memoria) ListarViajeInmediatos() []models.ViajeInmediato {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.ViajeInmediato, len(m.viajeInmediatos))
	copy(copia, m.viajeInmediatos)
	return copia
}

// BuscarProductoPorID devuelve el producto con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarViajeInmediatoPorID(id int) (models.ViajeInmediato, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, vi := range m.viajeInmediatos {
		if vi.ID == id {
			return vi, true
		}
	}
	return models.ViajeInmediato{}, false
}

// CrearProducto agrega un producto nuevo y devuelve el producto con ID asignado.
func (m *Memoria) CrearViajeInmediato(vi models.ViajeInmediato) models.ViajeInmediato {
	m.mu.Lock()
	defer m.mu.Unlock()

	vi.ID = m.nextViajeInmediatoID
	m.nextViajeInmediatoID++
	m.viajeInmediatos = append(m.viajeInmediatos, vi)
	return vi
}

// ActualizarProducto reemplaza el producto con el ID dado.
func (m *Memoria) ActualizarViajeInmediato(id int, datos models.ViajeInmediato) (models.ViajeInmediato, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, vi := range m.viajeInmediatos {
		if vi.ID == id {
			datos.ID = id
			m.viajeInmediatos[i] = datos
			return datos, true
		}
	}
	return models.ViajeInmediato{}, false
}

// BorrarProducto elimina el producto con el ID dado.
func (m *Memoria) BorrarViajeInmediato(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, vi := range m.viajeInmediatos {
		if vi.ID == id {
			m.viajeInmediatos = append(m.viajeInmediatos[:i], m.viajeInmediatos[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// SolicitudViaje
// =========================================================

// SeedCategorias carga categorías iniciales que coinciden con CategoriaID de los productos pre-cargados.
func (m *Memoria) SeedSolicitudViajes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.solicitudViajes = []models.SolicitudViaje{
		{ID: 1, ViajeID: 1, PasajeroID: 2, Estado: "Pendiente"},
		{ID: 2, ViajeID: 1, PasajeroID: 3, Estado: "Aceptada"},
		{ID: 3, ViajeID: 2, PasajeroID: 4, Estado: "Aceptada"},
		{ID: 4, ViajeID: 3, PasajeroID: 5, Estado: "Rechazada"},
		{ID: 5, ViajeID: 4, PasajeroID: 6, Estado: "Pendiente"},
		{ID: 6, ViajeID: 5, PasajeroID: 1, Estado: "Aceptada"},
	}

	m.nextSolicitudViajeID = 7
}

// ListarCategorias devuelve todas las categorías en memoria.
func (m *Memoria) ListarSolicitudViajes() []models.SolicitudViaje {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.SolicitudViaje, len(m.solicitudViajes))
	copy(copia, m.solicitudViajes)
	return copia
}

// BuscarCategoriaPorID devuelve la categoría con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarSolicitudViajePorID(id int) (models.SolicitudViaje, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, sv := range m.solicitudViajes {
		if sv.ID == id {
			return sv, true
		}
	}
	return models.SolicitudViaje{}, false
}

// CrearCategoria agrega una categoría nueva y devuelve la categoría con ID asignado.
func (m *Memoria) CrearSolicitudViaje(sv models.SolicitudViaje) models.SolicitudViaje {
	m.mu.Lock()
	defer m.mu.Unlock()

	sv.ID = m.nextSolicitudViajeID
	m.nextSolicitudViajeID++
	m.solicitudViajes = append(m.solicitudViajes, sv)
	return sv
}

// ActualizarCategoria reemplaza la categoría con el ID dado.
func (m *Memoria) ActualizarSolicitudViaje(id int, datos models.SolicitudViaje) (models.SolicitudViaje, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, sv := range m.solicitudViajes {
		if sv.ID == id {
			datos.ID = id
			m.solicitudViajes[i] = datos
			return datos, true
		}
	}
	return models.SolicitudViaje{}, false
}

// BorrarCategoria elimina la categoría con el ID dado.
func (m *Memoria) BorrarSolicitudViaje(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, sv := range m.solicitudViajes {
		if sv.ID == id {
			m.solicitudViajes = append(m.solicitudViajes[:i], m.solicitudViajes[i+1:]...)
			return true
		}
	}
	return false
}

// =========================================================
// ParticipanteViaje
// =========================================================

// SeedCategorias carga categorías iniciales que coinciden con CategoriaID de los productos pre-cargados.
func (m *Memoria) SeedParticipanteViajes() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.participanteViajes = []models.ParticipanteViaje{
		{ID: 1, ViajeID: 1, UsuarioID: 2},
		{ID: 2, ViajeID: 1, UsuarioID: 3},
		{ID: 3, ViajeID: 2, UsuarioID: 4},
		{ID: 4, ViajeID: 3, UsuarioID: 5},
		{ID: 5, ViajeID: 4, UsuarioID: 6},
		{ID: 6, ViajeID: 5, UsuarioID: 1},
	}
	m.nextParticipanteViajeID = 7
}

// ListarCategorias devuelve todas las categorías en memoria.
func (m *Memoria) ListarParticipanteViajes() []models.ParticipanteViaje {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.ParticipanteViaje, len(m.participanteViajes))
	copy(copia, m.participanteViajes)
	return copia
}

// BuscarCategoriaPorID devuelve la categoría con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarParticipanteViajePorID(id int) (models.ParticipanteViaje, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, pv := range m.participanteViajes {
		if pv.ID == id {
			return pv, true
		}
	}
	return models.ParticipanteViaje{}, false
}

// CrearCategoria agrega una categoría nueva y devuelve la categoría con ID asignado.
func (m *Memoria) CrearParticipanteViaje(pv models.ParticipanteViaje) models.ParticipanteViaje {
	m.mu.Lock()
	defer m.mu.Unlock()

	pv.ID = m.nextParticipanteViajeID
	m.nextParticipanteViajeID++
	m.participanteViajes = append(m.participanteViajes, pv)
	return pv
}

// ActualizarCategoria reemplaza la categoría con el ID dado.
func (m *Memoria) ActualizarParticipanteViaje(id int, datos models.ParticipanteViaje) (models.ParticipanteViaje, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, pv := range m.participanteViajes {
		if pv.ID == id {
			datos.ID = id
			m.participanteViajes[i] = datos
			return datos, true
		}
	}
	return models.ParticipanteViaje{}, false
}

// BorrarCategoria elimina la categoría con el ID dado.
func (m *Memoria) BorrarParticipanteViaje(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, pv := range m.participanteViajes {
		if pv.ID == id {
			m.participanteViajes = append(m.participanteViajes[:i], m.participanteViajes[i+1:]...)
			return true
		}
	}
	return false
}
