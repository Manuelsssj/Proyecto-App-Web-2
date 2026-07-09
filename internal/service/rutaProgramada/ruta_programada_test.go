package service

import (
	"testing"

	models "RideUleam/internal/models/rutaProgramada"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ======================================================
// MOCK DEL REPOSITORY / ALMACÉN
// ======================================================
// Este mock simula el repository usando testify.
// Así probamos el service sin depender de la base de datos real.

type mockAlmacenRuta struct {
	mock.Mock
}

// =====================
// MÉTODOS MOCK: RUTAS PROGRAMADAS
// =====================

func (m *mockAlmacenRuta) ListarRutasProgramadas() []models.RutaProgramada {
	args := m.Called()
	return args.Get(0).([]models.RutaProgramada)
}

func (m *mockAlmacenRuta) BuscarRutaProgramadaPorID(id int) (models.RutaProgramada, bool) {
	args := m.Called(id)
	return args.Get(0).(models.RutaProgramada), args.Bool(1)
}

func (m *mockAlmacenRuta) CrearRutaProgramada(ruta models.RutaProgramada) models.RutaProgramada {
	args := m.Called(ruta)
	return args.Get(0).(models.RutaProgramada)
}

func (m *mockAlmacenRuta) ActualizarRutaProgramada(id int, datos models.RutaProgramada) (models.RutaProgramada, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.RutaProgramada), args.Bool(1)
}

func (m *mockAlmacenRuta) BorrarRutaProgramada(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// =====================
// MÉTODOS MOCK: HORARIOS DE RUTA
// =====================

func (m *mockAlmacenRuta) ListarHorariosRuta() []models.HorarioRuta {
	args := m.Called()
	return args.Get(0).([]models.HorarioRuta)
}

func (m *mockAlmacenRuta) ListarHorariosPorRutaID(rutaID int) []models.HorarioRuta {
	args := m.Called(rutaID)
	return args.Get(0).([]models.HorarioRuta)
}

func (m *mockAlmacenRuta) BuscarHorarioRutaPorID(id int) (models.HorarioRuta, bool) {
	args := m.Called(id)
	return args.Get(0).(models.HorarioRuta), args.Bool(1)
}

func (m *mockAlmacenRuta) CrearHorarioRuta(horario models.HorarioRuta) models.HorarioRuta {
	args := m.Called(horario)
	return args.Get(0).(models.HorarioRuta)
}

func (m *mockAlmacenRuta) ActualizarHorarioRuta(id int, datos models.HorarioRuta) (models.HorarioRuta, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.HorarioRuta), args.Bool(1)
}

func (m *mockAlmacenRuta) BorrarHorarioRuta(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// =====================
// MÉTODOS MOCK: MANTENIMIENTOS VEHÍCULO
// =====================

func (m *mockAlmacenRuta) ListarMantenimientosVehiculo() []models.MantenimientoVehiculo {
	args := m.Called()
	return args.Get(0).([]models.MantenimientoVehiculo)
}

func (m *mockAlmacenRuta) ListarMantenimientosPorVehiculoID(vehiculoID int) []models.MantenimientoVehiculo {
	args := m.Called(vehiculoID)
	return args.Get(0).([]models.MantenimientoVehiculo)
}

func (m *mockAlmacenRuta) BuscarMantenimientoVehiculoPorID(id int) (models.MantenimientoVehiculo, bool) {
	args := m.Called(id)
	return args.Get(0).(models.MantenimientoVehiculo), args.Bool(1)
}

func (m *mockAlmacenRuta) CrearMantenimientoVehiculo(mantenimiento models.MantenimientoVehiculo) models.MantenimientoVehiculo {
	args := m.Called(mantenimiento)
	return args.Get(0).(models.MantenimientoVehiculo)
}

func (m *mockAlmacenRuta) ActualizarMantenimientoVehiculo(id int, datos models.MantenimientoVehiculo) (models.MantenimientoVehiculo, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.MantenimientoVehiculo), args.Bool(1)
}

func (m *mockAlmacenRuta) BorrarMantenimientoVehiculo(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// ======================================================
// TEST 1: CREAR RUTA PROGRAMADA VÁLIDA
// ======================================================
// Objetivo:
// Verificar que el service acepta una ruta correcta y llama al repository.

func TestCrearRutaProgramada_Valida(t *testing.T) {
	mockRepo := new(mockAlmacenRuta)
	servicio := NewRutaProgramadaService(mockRepo)

	ruta := models.RutaProgramada{
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "ULEAM",
		Costo:       0.75,
	}

	rutaCreada := ruta
	rutaCreada.ID = 1

	mockRepo.On("CrearRutaProgramada", ruta).Return(rutaCreada)

	resultado, err := servicio.CrearRutaProgramada(ruta)

	assert.NoError(t, err)
	assert.Equal(t, 1, resultado.ID)
	assert.Equal(t, "Los Esteros", resultado.Origen)
	assert.Equal(t, "ULEAM", resultado.Destino)
	assert.Equal(t, 0.75, resultado.Costo)

	mockRepo.AssertExpectations(t)
}

// ======================================================
// TEST 2: RECHAZAR RUTA CON COSTO NEGATIVO
// ======================================================
// Objetivo:
// Verificar que el service bloquea una ruta inválida.
// Además confirma que el repository NO se llama.

func TestCrearRutaProgramada_RechazaCostoNegativoYNoLlegaAlAlmacen(t *testing.T) {
	mockRepo := new(mockAlmacenRuta)
	servicio := NewRutaProgramadaService(mockRepo)

	ruta := models.RutaProgramada{
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "ULEAM",
		Costo:       -1,
	}

	_, err := servicio.CrearRutaProgramada(ruta)

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "CrearRutaProgramada")
}

// ======================================================
// TEST 3: OBTENER RUTA PROGRAMADA
// ======================================================
// Objetivo:
// En un solo test se revisan dos casos:
// 1. Cuando la ruta existe.
// 2. Cuando la ruta no existe y debe devolver ErrNoEncontrado.

func TestObtenerRutaProgramada_EncontradaYNoEncontrada(t *testing.T) {
	mockRepo := new(mockAlmacenRuta)
	servicio := NewRutaProgramadaService(mockRepo)

	rutaEsperada := models.RutaProgramada{
		ID:          1,
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "ULEAM",
		Costo:       0.75,
	}

	mockRepo.On("BuscarRutaProgramadaPorID", 1).Return(rutaEsperada, true)
	mockRepo.On("BuscarRutaProgramadaPorID", 99).Return(models.RutaProgramada{}, false)

	// Caso 1: ruta encontrada
	resultado, err := servicio.ObtenerRutaProgramada(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, resultado.ID)
	assert.Equal(t, "Los Esteros", resultado.Origen)

	// Caso 2: ruta no encontrada
	_, err = servicio.ObtenerRutaProgramada(99)

	assert.ErrorIs(t, err, ErrNoEncontrado)

	mockRepo.AssertExpectations(t)
}

// ======================================================
// TEST 4: ACTUALIZAR Y BORRAR RUTA PROGRAMADA
// ======================================================
// Objetivo:
// En un solo test se cubren dos operaciones del service:
// actualizar una ruta válida y borrar una ruta existente.

func TestRutaProgramada_ActualizarYBorrar(t *testing.T) {
	mockRepo := new(mockAlmacenRuta)
	servicio := NewRutaProgramadaService(mockRepo)

	datos := models.RutaProgramada{
		ConductorID: 1,
		Origen:      "Tarqui",
		Destino:     "ULEAM",
		Costo:       1.00,
	}

	actualizada := datos
	actualizada.ID = 1

	mockRepo.On("ActualizarRutaProgramada", 1, datos).Return(actualizada, true)
	mockRepo.On("BorrarRutaProgramada", 1).Return(true)

	// Caso 1: actualizar ruta
	resultado, err := servicio.ActualizarRutaProgramada(1, datos)

	assert.NoError(t, err)
	assert.Equal(t, 1, resultado.ID)
	assert.Equal(t, "Tarqui", resultado.Origen)
	assert.Equal(t, 1.00, resultado.Costo)

	// Caso 2: borrar ruta
	err = servicio.BorrarRutaProgramada(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

// ======================================================
// TEST 5: CREAR HORARIO DE RUTA
// ======================================================
// Objetivo:
// En un solo test se revisan dos casos:
// 1. Crear horario válido cuando la ruta existe.
// 2. Rechazar horario inválido cuando la hora está vacía.

func TestCrearHorarioRuta_ValidoYHoraVacia(t *testing.T) {
	mockRepo := new(mockAlmacenRuta)
	servicio := NewRutaProgramadaService(mockRepo)

	horarioValido := models.HorarioRuta{
		RutaID: 1,
		Dia:    "Lunes",
		Hora:   "07:00",
	}

	horarioCreado := horarioValido
	horarioCreado.ID = 1

	rutaExistente := models.RutaProgramada{
		ID:          1,
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "ULEAM",
		Costo:       0.75,
	}

	mockRepo.On("BuscarRutaProgramadaPorID", 1).Return(rutaExistente, true)
	mockRepo.On("CrearHorarioRuta", horarioValido).Return(horarioCreado)

	// Caso 1: horario válido
	resultado, err := servicio.CrearHorarioRuta(horarioValido)

	assert.NoError(t, err)
	assert.Equal(t, 1, resultado.ID)
	assert.Equal(t, "Lunes", resultado.Dia)
	assert.Equal(t, "07:00", resultado.Hora)

	// Caso 2: horario con hora vacía
	horarioInvalido := models.HorarioRuta{
		RutaID: 1,
		Dia:    "Lunes",
		Hora:   "",
	}

	_, err = servicio.CrearHorarioRuta(horarioInvalido)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "CrearHorarioRuta", horarioInvalido)
}

// ======================================================
// TEST 6: CREAR MANTENIMIENTO VEHÍCULO
// ======================================================
// Objetivo:
// En un solo test se revisan dos casos:
// 1. Crear mantenimiento válido.
// 2. Rechazar mantenimiento inválido cuando el motivo está vacío.

func TestCrearMantenimientoVehiculo_ValidoYMotivoVacio(t *testing.T) {
	mockRepo := new(mockAlmacenRuta)
	servicio := NewRutaProgramadaService(mockRepo)

	mantenimientoValido := models.MantenimientoVehiculo{
		VehiculoID:  1,
		FechaInicio: "2026-07-08",
		FechaFin:    "2026-07-10",
		Motivo:      "Cambio de aceite",
	}

	mantenimientoCreado := mantenimientoValido
	mantenimientoCreado.ID = 1

	mockRepo.On("CrearMantenimientoVehiculo", mantenimientoValido).Return(mantenimientoCreado)

	// Caso 1: mantenimiento válido
	resultado, err := servicio.CrearMantenimientoVehiculo(mantenimientoValido)

	assert.NoError(t, err)
	assert.Equal(t, 1, resultado.ID)
	assert.Equal(t, "Cambio de aceite", resultado.Motivo)

	// Caso 2: mantenimiento con motivo vacío
	mantenimientoInvalido := models.MantenimientoVehiculo{
		VehiculoID:  1,
		FechaInicio: "2026-07-08",
		FechaFin:    "2026-07-10",
		Motivo:      "",
	}

	_, err = servicio.CrearMantenimientoVehiculo(mantenimientoInvalido)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "CrearMantenimientoVehiculo", mantenimientoInvalido)
}

// ======================================================
// TEST 7: LISTAR, OBTENER, ACTUALIZAR Y BORRAR HORARIO
// ======================================================
// Objetivo:
// Cubrir varias funciones del service relacionadas con horarios:
// listar horarios, obtener un horario existente, actualizarlo y borrarlo.

func TestHorarioRuta_ListarObtenerActualizarYBorrar(t *testing.T) {
	mockRepo := new(mockAlmacenRuta)
	servicio := NewRutaProgramadaService(mockRepo)

	horario := models.HorarioRuta{
		ID:     1,
		RutaID: 1,
		Dia:    "Lunes",
		Hora:   "07:00",
	}

	horarioActualizado := models.HorarioRuta{
		ID:     1,
		RutaID: 1,
		Dia:    "Martes",
		Hora:   "08:00",
	}

	mockRepo.On("ListarHorariosRuta").Return([]models.HorarioRuta{horario})
	mockRepo.On("BuscarHorarioRutaPorID", 1).Return(horario, true)

	mockRepo.On("BuscarRutaProgramadaPorID", 1).Return(models.RutaProgramada{
		ID:          1,
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "ULEAM",
		Costo:       0.75,
	}, true)

	mockRepo.On("ActualizarHorarioRuta", 1, horarioActualizado).Return(horarioActualizado, true)
	mockRepo.On("BorrarHorarioRuta", 1).Return(true)

	// Caso 1: listar horarios
	lista := servicio.ListarHorariosRuta()

	assert.Len(t, lista, 1)
	assert.Equal(t, "Lunes", lista[0].Dia)

	// Caso 2: obtener horario existente
	obtenido, err := servicio.ObtenerHorarioRuta(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, obtenido.ID)
	assert.Equal(t, "07:00", obtenido.Hora)

	// Caso 3: actualizar horario
	actualizado, err := servicio.ActualizarHorarioRuta(1, horarioActualizado)

	assert.NoError(t, err)
	assert.Equal(t, "Martes", actualizado.Dia)
	assert.Equal(t, "08:00", actualizado.Hora)

	// Caso 4: borrar horario
	err = servicio.BorrarHorarioRuta(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
