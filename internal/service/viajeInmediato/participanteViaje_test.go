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
type participanteViajeRepoMock struct {
	mock.Mock
}

func (m *participanteViajeRepoMock) ListarParticipanteViajes() []models.ParticipanteViaje {
	args := m.Called()
	return args.Get(0).([]models.ParticipanteViaje)
}

func (m *participanteViajeRepoMock) BuscarParticipanteViajePorID(id int) (models.ParticipanteViaje, bool) {
	args := m.Called(id)
	return args.Get(0).(models.ParticipanteViaje), args.Bool(1)
}

func (m *participanteViajeRepoMock) CrearParticipanteViaje(pv models.ParticipanteViaje) models.ParticipanteViaje {
	args := m.Called(pv)
	return args.Get(0).(models.ParticipanteViaje)
}

func (m *participanteViajeRepoMock) ActualizarParticipanteViaje(id int, datos models.ParticipanteViaje) (models.ParticipanteViaje, bool) {
	args := m.Called(id, datos)
	return args.Get(0).(models.ParticipanteViaje), args.Bool(1)
}

func (m *participanteViajeRepoMock) BorrarParticipanteViaje(id int) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// Red de seguridad en tiempo de compilacion: el mock DEBE cumplir el contrato.
var _ storage.SolicitudViajeRepository = (*solicitudViajeRepoMock)(nil)

func TestParticipanteViajeService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.ParticipanteViaje
		errEsperado   error
		debePersistir bool
	}{
		{
			nombre: "viaje_id invalido -> ErrViajeIDInvalido",
			entrada: models.ParticipanteViaje{
				ViajeID:   0,
				UsuarioID: 1,
			},
			errEsperado:   service.ErrViajeIDInvalido,
			debePersistir: false,
		},
		{
			nombre: "participante valido -> sin error y se persiste",
			entrada: models.ParticipanteViaje{
				ViajeID:   1,
				UsuarioID: 2,
			},
			errEsperado:   nil,
			debePersistir: true,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := new(participanteViajeRepoMock)

			if c.debePersistir {
				guardado := c.entrada
				guardado.ID = 42

				repo.On("CrearParticipanteViaje", c.entrada).Return(guardado)
			}

			svc := service.NewParticipanteViajeService(repo)

			creado, err := svc.Crear(c.entrada)

			if c.errEsperado != nil {
				require.ErrorIs(t, err, c.errEsperado)

				repo.AssertNotCalled(t, "CrearParticipanteViaje")
			} else {
				require.NoError(t, err)

				assert.Equal(t, 42, creado.ID)

				repo.AssertCalled(t, "CrearParticipanteViaje", c.entrada)
			}
		})
	}
}

// TestProductoService_Obtener_NoEncontrado muestra como el service traduce el
// comma-ok del repositorio (false) en un error de dominio (ErrNoEncontrado).
func TestParticipanteViajeService_Obtener_NoEncontrado(t *testing.T) {
	repo := new(participanteViajeRepoMock)
	repo.On("BuscarParticipanteViajePorID", 999).Return(models.ParticipanteViaje{}, false)
	svc := service.NewParticipanteViajeService(repo)

	_, err := svc.Obtener(999)

	require.ErrorIs(t, err, service.ErrParticipanteViajeNoEncontrado)
	repo.AssertExpectations(t)
}
