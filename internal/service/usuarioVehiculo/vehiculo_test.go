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
	args := m.Called()
	return args.Get(0).([]models.Vehiculo)
}

func (m *vehiculoRepoMock) BuscarVehiculoPorID(id int) (models.Vehiculo, bool) {
	args := m.Called(id)
	return args.Get(0).(models.Vehiculo), args.Bool(1)
}

func (m *vehiculoRepoMock) CrearVehiculo(v models.Vehiculo) models.Vehiculo {
	args := m.Called(v)
	return args.Get(0).(models.Vehiculo)
}

func (m *vehiculoRepoMock) ActualizarVehiculo(id int, datos models.Vehiculo) (models.Vehiculo, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.Vehiculo), args.Bool(1)
}

func (m *vehiculoRepoMock) BorrarVehiculo(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// Red de seguridad en tiempo de compilacion: el mock DEBE cumplir el contrato.
var _ storage.VehiculoRepository = (*vehiculoRepoMock)(nil)

func TestVehiculoService_Crear(t *testing.T) {

	casos := []struct {
		nombre        string
		entrada       models.Vehiculo
		errEsperado   error // nil = se espera exito
		debePersistir bool
	}{
		{
			nombre:        "conductor_id invalido -> ErrConductorIDInvalido",
			entrada:       models.Vehiculo{ConductorID: 0, Placa: "ABC123", Capacidad: 5},
			errEsperado:   service.ErrConductorIDInvalido,
			debePersistir: false,
		},
		{
			nombre:        "vehiculo valido -> sin error y se persiste",
			entrada:       models.Vehiculo{ConductorID: 1, Placa: "ABC123", Marca: "Toyota", Modelo: "Corolla", Capacidad: 5},
			errEsperado:   nil,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			// Preparar: un mock nuevo por caso para no arrastrar estado.
			repo := new(vehiculoRepoMock)

			if c.debePersistir {
				// El repo devuelve el vehiculo con un ID asignado.
				guardado := c.entrada
				guardado.ID = 42

				repo.On("CrearVehiculo", c.entrada).Return(guardado)
			}

			svc := service.NuevoVehiculoService(repo)

			// Ejecutar.
			creado, err := svc.Crear(c.entrada)

			// Verificar.
			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)

				// Si la validacion falla, el repo NO debe haberse tocado.
				repo.AssertNotCalled(t, "CrearVehiculo")
			} else {
				require.NoError(t, err)

				assert.Equal(t, 42, creado.ID, "el service debe devolver el vehiculo que entrego el repo")

				repo.AssertCalled(t, "CrearVehiculo", c.entrada)
			}
		})
	}
}

// TestProductoService_Obtener_NoEncontrado muestra como el service traduce el
// comma-ok del repositorio (false) en un error de dominio (ErrNoEncontrado).
func TestVehiculoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(vehiculoRepoMock)
	repo.On("BuscarVehiculoPorID", 999).Return(models.Vehiculo{}, false)
	svc := service.NuevoVehiculoService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrVehiculoNoEncontrado)
	repo.AssertExpectations(t)
}
