package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	models "RideUleam/internal/models/usuarioVehiculo"
	service "RideUleam/internal/service/usuarioVehiculo"
	storage "RideUleam/internal/storage/usuarioVehiculo"
)

// productoRepoMock es un doble de prueba de storage.ProductoRepository.
//
// Gracias al ISP, ProductoService depende SOLO de esta interfaz estrecha (5
// metodos), no del Almacen completo (10). Por eso el mock implementa 5 metodos
// y no toca para nada categorias. Si manana ProductoService necesitara mas,
// el compilador nos obligaria a anadirlos aqui: ese es el valor de la asercion
// de abajo.
type vehiculoRepoMock struct {
	mock.Mock
}

func (m *vehiculoRepoMock) ListarVehiculos() []models.Vehiculo {
	return m.Called().Get(0).([]models.Vehiculo)
}
func (m *vehiculoRepoMock) BuscarVehiculoPorID(id int) (models.Vehiculo, bool) {
	a := m.Called(id)
	return a.Get(0).(models.Vehiculo), a.Bool(1)
}
func (m *vehiculoRepoMock) CrearVehiculo(v models.Vehiculo) models.Vehiculo {
	return m.Called(v).Get(0).(models.Vehiculo)
}
func (m *vehiculoRepoMock) ActualizarVehiculo(id int, datos models.Vehiculo) (models.Vehiculo, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.Vehiculo), a.Bool(1)
}
func (m *vehiculoRepoMock) BorrarVehiculo(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.VehiculoRepository = (*vehiculoRepoMock)(nil)

func TestVehiculoService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.Vehiculo
		errEsperado   error // nil = exito
		debePersistir bool
	}{
		{"conductor_id invalido rechazado", models.Vehiculo{ConductorID: 0}, service.ErrConductorIDInvalido, false},
		{"vehiculo valido se persiste", models.Vehiculo{ConductorID: 1, Placa: "ABC-123", Marca: "Toyota", Modelo: "Corolla", Capacidad: 5}, nil, true},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(vehiculoRepoMock)
			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 42
				repo.On("CrearVehiculo", c.entrada).Return(guardado)
			}
			svc := service.NuevoVehiculoService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)
				repo.AssertNotCalled(t, "CrearVehiculo") // la validacion corto antes
			} else {
				require.NoError(t, err)
				assert.Equal(t, 42, creado.ID)
				repo.AssertCalled(t, "CrearVehiculo", c.entrada)
			}
		})
	}
}

// --- Obtener: comma-ok del repo traducido a error de dominio ---

func TestVehiculoService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(vehiculoRepoMock)
		repo.On("BuscarVehiculoPorID", 1).Return(models.Vehiculo{ID: 1, Placa: "ABC-123"}, true)
		v, err := service.NuevoVehiculoService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "ABC-123", v.Placa)
	})
	t.Run("no existe -> ErrVehiculoNoEncontrado", func(t *testing.T) {
		repo := new(vehiculoRepoMock)
		repo.On("BuscarVehiculoPorID", 999).Return(models.Vehiculo{}, false)
		_, err := service.NuevoVehiculoService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrVehiculoNoEncontrado)
	})
}

// --- Actualizar: valida ANTES de tocar el repo, y mapea el no encontrado ---

func TestVehiculoService_Actualizar(t *testing.T) {
	datos := models.Vehiculo{ConductorID: 1, Placa: "XYZ-999", Marca: "Kia", Modelo: "Rio", Capacidad: 4}

	t.Run("valido", func(t *testing.T) {
		repo := new(vehiculoRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarVehiculo", 1, datos).Return(actualizado, true)
		v, err := service.NuevoVehiculoService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, v.ID)
	})
	t.Run("no existe -> ErrVehiculoNoEncontrado", func(t *testing.T) {
		repo := new(vehiculoRepoMock)
		repo.On("ActualizarVehiculo", 999, datos).Return(models.Vehiculo{}, false)
		_, err := service.NuevoVehiculoService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrVehiculoNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(vehiculoRepoMock)
		_, err := service.NuevoVehiculoService(repo).Actualizar(1, models.Vehiculo{ConductorID: 0})
		require.ErrorIs(t, err, service.ErrConductorIDInvalido)
		repo.AssertNotCalled(t, "ActualizarVehiculo")
	})
}

// --- Borrar ---

func TestVehiculoService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(vehiculoRepoMock)
		repo.On("BorrarVehiculo", 1).Return(true)
		require.NoError(t, service.NuevoVehiculoService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrVehiculoNoEncontrado", func(t *testing.T) {
		repo := new(vehiculoRepoMock)
		repo.On("BorrarVehiculo", 999).Return(false)
		require.ErrorIs(t, service.NuevoVehiculoService(repo).Borrar(999), service.ErrVehiculoNoEncontrado)
	})
}

// --- Listar: el service solo delega ---

func TestVehiculoService_Listar(t *testing.T) {
	repo := new(vehiculoRepoMock)
	repo.On("ListarVehiculos").Return([]models.Vehiculo{{ID: 1}, {ID: 2}})
	lista := service.NuevoVehiculoService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
