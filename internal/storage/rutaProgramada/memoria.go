package storage

import (
	"sync"
	"time"

	models "RideUleam/internal/models/rutaProgramada"
	usuarioModels "RideUleam/internal/models/usuario"
)

type Memoria struct {
	mu sync.Mutex

	rutasProgramadas       []models.RutaProgramada
	horariosRuta           []models.HorarioRuta
	mantenimientosVehiculo []models.MantenimientoVehiculo
	usuarios               []usuarioModels.Usuario

	nextRutaProgramadaID int
	nextHorarioRutaID    int
	nextMantenimientoID  int
	nextUsuarioID        int
}

func NuevaMemoria() *Memoria {
	return &Memoria{
		rutasProgramadas:       []models.RutaProgramada{},
		horariosRuta:           []models.HorarioRuta{},
		mantenimientosVehiculo: []models.MantenimientoVehiculo{},
		usuarios:               []usuarioModels.Usuario{},

		nextRutaProgramadaID: 1,
		nextHorarioRutaID:    1,
		nextMantenimientoID:  1,
		nextUsuarioID:        1,
	}
}

// CRUD Rutas Programadas

func (m *Memoria) ListarRutasProgramadas() []models.RutaProgramada {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.rutasProgramadas
}

func (m *Memoria) BuscarRutaProgramadaPorID(id int) (models.RutaProgramada, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, ruta := range m.rutasProgramadas {
		if ruta.ID == id {
			return ruta, true
		}
	}

	return models.RutaProgramada{}, false
}

func (m *Memoria) CrearRutaProgramada(ruta models.RutaProgramada) models.RutaProgramada {
	m.mu.Lock()
	defer m.mu.Unlock()

	ruta.ID = m.nextRutaProgramadaID
	m.nextRutaProgramadaID++

	m.rutasProgramadas = append(m.rutasProgramadas, ruta)

	return ruta
}

func (m *Memoria) ActualizarRutaProgramada(id int, datos models.RutaProgramada) (models.RutaProgramada, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, ruta := range m.rutasProgramadas {
		if ruta.ID == id {
			datos.ID = id
			m.rutasProgramadas[i] = datos
			return datos, true
		}
	}

	return models.RutaProgramada{}, false
}

func (m *Memoria) BorrarRutaProgramada(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, ruta := range m.rutasProgramadas {
		if ruta.ID == id {
			m.rutasProgramadas = append(m.rutasProgramadas[:i], m.rutasProgramadas[i+1:]...)
			return true
		}
	}

	return false
}

// CRUD Horarios de Ruta

func (m *Memoria) ListarHorariosRuta() []models.HorarioRuta {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.horariosRuta
}

func (m *Memoria) ListarHorariosPorRutaID(rutaID int) []models.HorarioRuta {
	m.mu.Lock()
	defer m.mu.Unlock()

	horariosFiltrados := []models.HorarioRuta{}

	for _, horario := range m.horariosRuta {
		if horario.RutaID == rutaID {
			horariosFiltrados = append(horariosFiltrados, horario)
		}
	}

	return horariosFiltrados
}

func (m *Memoria) BuscarHorarioRutaPorID(id int) (models.HorarioRuta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, horario := range m.horariosRuta {
		if horario.ID == id {
			return horario, true
		}
	}

	return models.HorarioRuta{}, false
}

func (m *Memoria) CrearHorarioRuta(horario models.HorarioRuta) models.HorarioRuta {
	m.mu.Lock()
	defer m.mu.Unlock()

	horario.ID = m.nextHorarioRutaID
	m.nextHorarioRutaID++

	m.horariosRuta = append(m.horariosRuta, horario)

	return horario
}

func (m *Memoria) ActualizarHorarioRuta(id int, datos models.HorarioRuta) (models.HorarioRuta, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, horario := range m.horariosRuta {
		if horario.ID == id {
			datos.ID = id
			m.horariosRuta[i] = datos
			return datos, true
		}
	}

	return models.HorarioRuta{}, false
}

func (m *Memoria) BorrarHorarioRuta(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, horario := range m.horariosRuta {
		if horario.ID == id {
			m.horariosRuta = append(m.horariosRuta[:i], m.horariosRuta[i+1:]...)
			return true
		}
	}

	return false
}

// CRUD Mantenimientos de Vehículo

func (m *Memoria) ListarMantenimientosVehiculo() []models.MantenimientoVehiculo {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.mantenimientosVehiculo
}

func (m *Memoria) ListarMantenimientosPorVehiculoID(vehiculoID int) []models.MantenimientoVehiculo {
	m.mu.Lock()
	defer m.mu.Unlock()

	mantenimientosFiltrados := []models.MantenimientoVehiculo{}

	for _, mantenimiento := range m.mantenimientosVehiculo {
		if mantenimiento.VehiculoID == vehiculoID {
			mantenimientosFiltrados = append(mantenimientosFiltrados, mantenimiento)
		}
	}

	return mantenimientosFiltrados
}

func (m *Memoria) BuscarMantenimientoVehiculoPorID(id int) (models.MantenimientoVehiculo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, mantenimiento := range m.mantenimientosVehiculo {
		if mantenimiento.ID == id {
			return mantenimiento, true
		}
	}

	return models.MantenimientoVehiculo{}, false
}

func (m *Memoria) CrearMantenimientoVehiculo(mantenimiento models.MantenimientoVehiculo) models.MantenimientoVehiculo {
	m.mu.Lock()
	defer m.mu.Unlock()

	mantenimiento.ID = m.nextMantenimientoID
	m.nextMantenimientoID++

	m.mantenimientosVehiculo = append(m.mantenimientosVehiculo, mantenimiento)

	return mantenimiento
}

func (m *Memoria) ActualizarMantenimientoVehiculo(id int, datos models.MantenimientoVehiculo) (models.MantenimientoVehiculo, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, mantenimiento := range m.mantenimientosVehiculo {
		if mantenimiento.ID == id {
			datos.ID = id
			m.mantenimientosVehiculo[i] = datos
			return datos, true
		}
	}

	return models.MantenimientoVehiculo{}, false
}

func (m *Memoria) BorrarMantenimientoVehiculo(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, mantenimiento := range m.mantenimientosVehiculo {
		if mantenimiento.ID == id {
			m.mantenimientosVehiculo = append(m.mantenimientosVehiculo[:i], m.mantenimientosVehiculo[i+1:]...)
			return true
		}
	}

	return false
}

func (m *Memoria) CargarDatosIniciales() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.rutasProgramadas = append(m.rutasProgramadas,
		models.RutaProgramada{
			ID:          m.nextRutaProgramadaID,
			ConductorID: 1,
			Origen:      "Los Esteros",
			Destino:     "ULEAM",
			Costo:       0.75,
		},
	)
	m.nextRutaProgramadaID++

	m.rutasProgramadas = append(m.rutasProgramadas,
		models.RutaProgramada{
			ID:          m.nextRutaProgramadaID,
			ConductorID: 2,
			Origen:      "Tarqui",
			Destino:     "ULEAM",
			Costo:       1.00,
		},
	)
	m.nextRutaProgramadaID++

	m.horariosRuta = append(m.horariosRuta,
		models.HorarioRuta{
			ID:     m.nextHorarioRutaID,
			RutaID: 1,
			Dia:    "Lunes",
			Hora:   "07:00",
		},
	)
	m.nextHorarioRutaID++

	m.horariosRuta = append(m.horariosRuta,
		models.HorarioRuta{
			ID:     m.nextHorarioRutaID,
			RutaID: 2,
			Dia:    "Martes",
			Hora:   "08:00",
		},
	)
	m.nextHorarioRutaID++

	m.mantenimientosVehiculo = append(m.mantenimientosVehiculo,
		models.MantenimientoVehiculo{
			ID:          m.nextMantenimientoID,
			VehiculoID:  1,
			FechaInicio: "2026-06-22",
			FechaFin:    "2026-06-25",
			Motivo:      "Cambio de aceite",
		},
	)
	m.nextMantenimientoID++
}

// =====================
// USUARIOS / AUTH
// =====================

func (m *Memoria) CrearUsuario(u usuarioModels.Usuario) (usuarioModels.Usuario, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	u.ID = m.nextUsuarioID
	m.nextUsuarioID++

	u.CreadoEn = time.Now()

	m.usuarios = append(m.usuarios, u)

	return u, nil
}

func (m *Memoria) BuscarUsuarioPorEmail(email string) (usuarioModels.Usuario, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, usuario := range m.usuarios {
		if usuario.Email == email {
			return usuario, true
		}
	}

	return usuarioModels.Usuario{}, false
}
