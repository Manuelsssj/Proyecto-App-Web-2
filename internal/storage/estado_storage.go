package storage

import "RideUleam/internal/models"

func (m *Memoria) CrearEstado(estado models.Estado) models.Estado {
	m.mu.Lock()
	defer m.mu.Unlock()

	estado.ID = m.nextEstadoID
	m.nextEstadoID++

	m.estados = append(m.estados, estado)
	return estado
}

func (m *Memoria) ListarEstados() []models.Estado {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Estado, len(m.estados))
	copy(copia, m.estados)
	return copia
}

func (m *Memoria) BuscarEstadoPorID(id int) (models.Estado, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, estado := range m.estados {
		if estado.ID == id {
			return estado, true
		}
	}
	return models.Estado{}, false
}

func (m *Memoria) ActualizarEstado(id int, datos models.Estado) (models.Estado, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, estado := range m.estados {
		if estado.ID == id {
			datos.ID = id
			m.estados[i] = datos
			return datos, true
		}
	}
	return models.Estado{}, false
}

func (m *Memoria) BorrarEstado(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, estado := range m.estados {
		if estado.ID == id {
			m.estados = append(m.estados[:i], m.estados[i+1:]...)
			return true
		}
	}
	return false
}
