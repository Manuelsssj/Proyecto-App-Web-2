package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	models "RideUleam/internal/models/viajeInmediato"
	service "RideUleam/internal/service/viajeInmediato"
	storage "RideUleam/internal/storage/viajeInmediato"
)

// productoRepoMock es un doble de prueba de storage.ProductoRepository.
//
// Gracias al ISP, ProductoService depende SOLO de esta interfaz estrecha (5
// metodos), no del Almacen completo (10). Por eso el mock implementa 5 metodos
// y no toca para nada categorias. Si manana ProductoService necesitara mas,
// el compilador nos obligaria a anadirlos aqui: ese es el valor de la asercion
// de abajo.
type viajeInmediatoRepoMock struct {
	mock.Mock
}

func (m *viajeInmediatoRepoMock) ListarViajeInmediatos() []models.ViajeInmediato {
	args := m.Called()
	return args.Get(0).([]models.ViajeInmediato)
}

func (m *viajeInmediatoRepoMock) BuscarViajeInmediatoPorID(id int) (models.ViajeInmediato, bool) {
	args := m.Called(id)
	return args.Get(0).(models.ViajeInmediato), args.Bool(1)
}

func (m *viajeInmediatoRepoMock) CrearViajeInmediato(vi models.ViajeInmediato) models.ViajeInmediato {
	args := m.Called(vi)
	return args.Get(0).(models.ViajeInmediato)
}

func (m *viajeInmediatoRepoMock) ActualizarViajeInmediato(id int, datos models.ViajeInmediato) (models.ViajeInmediato, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.ViajeInmediato), args.Bool(1)
}

func (m *viajeInmediatoRepoMock) BorrarViajeInmediato(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// Red de seguridad en tiempo de compilacion: el mock DEBE cumplir el contrato.
var _ storage.ViajeInmediatoRepository = (*viajeInmediatoRepoMock)(nil)

func TestViajeInmediatoService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.ViajeInmediato
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre: "conductor_id invalido -> ErrConductorIDInvalido",
			entrada: models.ViajeInmediato{
				ConductorID: 0,
				Origen:      "ULEAM",
				Destino:     "Terminal",
				HoraSalida:  "08:00",
				Cupos:       4,
				Estado:      "Disponible",
			},
			errEsperado:   service.ErrConductorIDInvalido,
			debePersistir: false,
		},
		{
			nombre: "viaje valido -> sin error y se persiste",
			entrada: models.ViajeInmediato{
				ConductorID: 1,
				Origen:      "ULEAM",
				Destino:     "Terminal",
				HoraSalida:  "08:00",
				Cupos:       4,
				Estado:      "Disponible",
			},
			errEsperado:   nil,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(viajeInmediatoRepoMock)

			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 42

				repo.On("CrearViajeInmediato", c.entrada).Return(guardado)
			}

			svc := service.NewViajeInmediatoService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)

				repo.AssertNotCalled(t, "CrearViajeInmediato")
			} else {
				require.NoError(t, err)

				assert.Equal(t, 42, creado.ID)

				repo.AssertCalled(t, "CrearViajeInmediato", c.entrada)
			}
		})
	}
}

// TestProductoService_Obtener_NoEncontrado muestra como el service traduce el
// comma-ok del repositorio (false) en un error de dominio (ErrNoEncontrado).
func TestViajeInmediatoService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(viajeInmediatoRepoMock)
	repo.On("BuscarViajeInmediatoPorID", 999).Return(models.ViajeInmediato{}, false)
	svc := service.NewViajeInmediatoService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrViajeInmediatoNoEncontrado)
	repo.AssertExpectations(t)
}
