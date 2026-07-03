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
type solicitudViajeRepoMock struct {
	mock.Mock
}

func (m *solicitudViajeRepoMock) ListarSolicitudViajes() []models.SolicitudViaje {
	args := m.Called()
	return args.Get(0).([]models.SolicitudViaje)
}

func (m *solicitudViajeRepoMock) BuscarSolicitudViajePorID(id int) (models.SolicitudViaje, bool) {
	args := m.Called(id)
	return args.Get(0).(models.SolicitudViaje), args.Bool(1)
}

func (m *solicitudViajeRepoMock) CrearSolicitudViaje(sv models.SolicitudViaje) models.SolicitudViaje {
	args := m.Called(sv)
	return args.Get(0).(models.SolicitudViaje)
}

func (m *solicitudViajeRepoMock) ActualizarSolicitudViaje(id int, datos models.SolicitudViaje) (models.SolicitudViaje, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.SolicitudViaje), args.Bool(1)
}

func (m *solicitudViajeRepoMock) BorrarSolicitudViaje(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// Red de seguridad en tiempo de compilacion: el mock DEBE cumplir el contrato.
var _ storage.SolicitudViajeRepository = (*solicitudViajeRepoMock)(nil)

func TestSolicitudViajeService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.SolicitudViaje
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre: "viaje_id invalido -> ErrViajeIDInvalido",
			entrada: models.SolicitudViaje{
				ViajeID:    0,
				PasajeroID: 1,
				Estado:     "Pendiente",
			},
			errEsperado:   service.ErrViajeIDInvalido,
			debePersistir: false,
		},
		{
			nombre: "solicitud valida -> sin error y se persiste",
			entrada: models.SolicitudViaje{
				ViajeID:    1,
				PasajeroID: 1,
				Estado:     "Pendiente",
			},
			errEsperado:   nil,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(solicitudViajeRepoMock)

			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 42

				repo.On("CrearSolicitudViaje", c.entrada).Return(guardado)
			}

			svc := service.NewSolicitudViajeService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)

				repo.AssertNotCalled(t, "CrearSolicitudViaje")
			} else {
				require.NoError(t, err)

				assert.Equal(t, 42, creado.ID)

				repo.AssertCalled(t, "CrearSolicitudViaje", c.entrada)
			}
		})
	}
}

// TestProductoService_Obtener_NoEncontrado muestra como el service traduce el
// comma-ok del repositorio (false) en un error de dominio (ErrNoEncontrado).
func TestSolicitudViajeService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(solicitudViajeRepoMock)
	repo.On("BuscarSolicitudViajePorID", 999).Return(models.SolicitudViaje{}, false)
	svc := service.NewSolicitudViajeService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrSolicitudViajeNoEncontrado)
	repo.AssertExpectations(t)
}
