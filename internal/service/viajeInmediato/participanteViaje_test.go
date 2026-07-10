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
	return m.Called().Get(0).([]models.ParticipanteViaje)
}
func (m *participanteViajeRepoMock) BuscarParticipanteViajePorID(id int) (models.ParticipanteViaje, bool) {
	a := m.Called(id)
	return a.Get(0).(models.ParticipanteViaje), a.Bool(1)
}
func (m *participanteViajeRepoMock) CrearParticipanteViaje(pv models.ParticipanteViaje) models.ParticipanteViaje {
	return m.Called(pv).Get(0).(models.ParticipanteViaje)
}
func (m *participanteViajeRepoMock) ActualizarParticipanteViaje(id int, datos models.ParticipanteViaje) (models.ParticipanteViaje, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.ParticipanteViaje), a.Bool(1)
}
func (m *participanteViajeRepoMock) BorrarParticipanteViaje(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.ParticipanteViajeRepository = (*participanteViajeRepoMock)(nil)

func TestParticipanteViajeService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.ParticipanteViaje
		errEsperado   error // nil = exito
		debePersistir bool
	}{
		{"viaje_id invalido rechazado", models.ParticipanteViaje{ViajeID: 0}, service.ErrViajeIDInvalido, false},
		{"participante valido se persiste", models.ParticipanteViaje{ViajeID: 1, UsuarioID: 2}, nil, true},
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
				repo.AssertNotCalled(t, "CrearParticipanteViaje") // la validacion corto antes
			} else {
				require.NoError(t, err)
				assert.Equal(t, 42, creado.ID)
				repo.AssertCalled(t, "CrearParticipanteViaje", c.entrada)
			}
		})
	}
}

// --- Obtener: comma-ok del repo traducido a error de dominio ---

func TestParticipanteViajeService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(participanteViajeRepoMock)
		repo.On("BuscarParticipanteViajePorID", 1).Return(models.ParticipanteViaje{ID: 1, ViajeID: 1, UsuarioID: 2}, true)
		p, err := service.NewParticipanteViajeService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, 1, p.ViajeID)
	})
	t.Run("no existe -> ErrParticipanteViajeNoEncontrado", func(t *testing.T) {
		repo := new(participanteViajeRepoMock)
		repo.On("BuscarParticipanteViajePorID", 999).Return(models.ParticipanteViaje{}, false)
		_, err := service.NewParticipanteViajeService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrParticipanteViajeNoEncontrado)
	})
}

// --- Actualizar: valida ANTES de tocar el repo, y mapea el no encontrado ---

func TestParticipanteViajeService_Actualizar(t *testing.T) {
	datos := models.ParticipanteViaje{ViajeID: 2, UsuarioID: 3}

	t.Run("valido", func(t *testing.T) {
		repo := new(participanteViajeRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarParticipanteViaje", 1, datos).Return(actualizado, true)
		p, err := service.NewParticipanteViajeService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, p.ID)
	})
	t.Run("no existe -> ErrParticipanteViajeNoEncontrado", func(t *testing.T) {
		repo := new(participanteViajeRepoMock)
		repo.On("ActualizarParticipanteViaje", 999, datos).Return(models.ParticipanteViaje{}, false)
		_, err := service.NewParticipanteViajeService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrParticipanteViajeNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(participanteViajeRepoMock)
		_, err := service.NewParticipanteViajeService(repo).Actualizar(1, models.ParticipanteViaje{ViajeID: 0})
		require.ErrorIs(t, err, service.ErrViajeIDInvalido)
		repo.AssertNotCalled(t, "ActualizarParticipanteViaje")
	})
}

// --- Borrar ---

func TestParticipanteViajeService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(participanteViajeRepoMock)
		repo.On("BorrarParticipanteViaje", 1).Return(true)
		require.NoError(t, service.NewParticipanteViajeService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrParticipanteViajeNoEncontrado", func(t *testing.T) {
		repo := new(participanteViajeRepoMock)
		repo.On("BorrarParticipanteViaje", 999).Return(false)
		require.ErrorIs(t, service.NewParticipanteViajeService(repo).Borrar(999), service.ErrParticipanteViajeNoEncontrado)
	})
}

// --- Listar: el service solo delega ---

func TestParticipanteViajeService_Listar(t *testing.T) {
	repo := new(participanteViajeRepoMock)
	repo.On("ListarParticipanteViajes").Return([]models.ParticipanteViaje{{ID: 1}, {ID: 2}})
	lista := service.NewParticipanteViajeService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
