package service

import (
	"testing"

	models "cmd/rideUleam/internal/models/rutaProgramada"
)

type mockAlmacenRuta struct {
	crearRutaFueLlamado bool
}

// =====================
// RUTAS PROGRAMADAS
// =====================

func (m *mockAlmacenRuta) ListarRutasProgramadas() []models.RutaProgramada {
	return []models.RutaProgramada{}
}

func (m *mockAlmacenRuta) BuscarRutaProgramadaPorID(id int) (models.RutaProgramada, bool) {
	return models.RutaProgramada{}, false
}

func (m *mockAlmacenRuta) CrearRutaProgramada(ruta models.RutaProgramada) models.RutaProgramada {
	m.crearRutaFueLlamado = true
	ruta.ID = 1
	return ruta
}

func (m *mockAlmacenRuta) ActualizarRutaProgramada(id int, datos models.RutaProgramada) (models.RutaProgramada, bool) {
	return models.RutaProgramada{}, false
}

func (m *mockAlmacenRuta) BorrarRutaProgramada(id int) bool {
	return false
}

// =====================
// HORARIOS DE RUTA
// =====================

func (m *mockAlmacenRuta) ListarHorariosRuta() []models.HorarioRuta {
	return []models.HorarioRuta{}
}

func (m *mockAlmacenRuta) ListarHorariosPorRutaID(rutaID int) []models.HorarioRuta {
	return []models.HorarioRuta{}
}

func (m *mockAlmacenRuta) BuscarHorarioRutaPorID(id int) (models.HorarioRuta, bool) {
	return models.HorarioRuta{}, false
}

func (m *mockAlmacenRuta) CrearHorarioRuta(horario models.HorarioRuta) models.HorarioRuta {
	return models.HorarioRuta{}
}

func (m *mockAlmacenRuta) ActualizarHorarioRuta(id int, datos models.HorarioRuta) (models.HorarioRuta, bool) {
	return models.HorarioRuta{}, false
}

func (m *mockAlmacenRuta) BorrarHorarioRuta(id int) bool {
	return false
}

// =====================
// MANTENIMIENTOS VEHÍCULO
// =====================

func (m *mockAlmacenRuta) ListarMantenimientosVehiculo() []models.MantenimientoVehiculo {
	return []models.MantenimientoVehiculo{}
}

func (m *mockAlmacenRuta) ListarMantenimientosPorVehiculoID(vehiculoID int) []models.MantenimientoVehiculo {
	return []models.MantenimientoVehiculo{}
}

func (m *mockAlmacenRuta) BuscarMantenimientoVehiculoPorID(id int) (models.MantenimientoVehiculo, bool) {
	return models.MantenimientoVehiculo{}, false
}

func (m *mockAlmacenRuta) CrearMantenimientoVehiculo(mantenimiento models.MantenimientoVehiculo) models.MantenimientoVehiculo {
	return models.MantenimientoVehiculo{}
}

func (m *mockAlmacenRuta) ActualizarMantenimientoVehiculo(id int, datos models.MantenimientoVehiculo) (models.MantenimientoVehiculo, bool) {
	return models.MantenimientoVehiculo{}, false
}

func (m *mockAlmacenRuta) BorrarMantenimientoVehiculo(id int) bool {
	return false
}

// =====================
// TEST SERVICE
// =====================

func TestCrearRutaProgramada_RechazaCostoNegativoYNoLlegaAlAlmacen(t *testing.T) {
	mock := &mockAlmacenRuta{}
	servicio := NewRutaProgramadaService(mock)

	ruta := models.RutaProgramada{
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "ULEAM",
		Costo:       -1,
	}

	_, err := servicio.CrearRutaProgramada(ruta)

	if err == nil {
		t.Fatal("se esperaba error cuando el costo es negativo")
	}

	if mock.crearRutaFueLlamado {
		t.Fatal("no se esperaba que la ruta llegue al almacen si tiene costo negativo")
	}
}
