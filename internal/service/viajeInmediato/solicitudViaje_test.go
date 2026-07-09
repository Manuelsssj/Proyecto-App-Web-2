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
	return m.Called().Get(0).([]models.SolicitudViaje)
}
func (m *solicitudViajeRepoMock) BuscarSolicitudViajePorID(id int) (models.SolicitudViaje, bool) {
	a := m.Called(id)
	return a.Get(0).(models.SolicitudViaje), a.Bool(1)
}
func (m *solicitudViajeRepoMock) CrearSolicitudViaje(sv models.SolicitudViaje) models.SolicitudViaje {
	return m.Called(sv).Get(0).(models.SolicitudViaje)
}
func (m *solicitudViajeRepoMock) ActualizarSolicitudViaje(id int, datos models.SolicitudViaje) (models.SolicitudViaje, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.SolicitudViaje), a.Bool(1)
}
func (m *solicitudViajeRepoMock) BorrarSolicitudViaje(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.SolicitudViajeRepository = (*solicitudViajeRepoMock)(nil)

func TestSolicitudViajeService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.SolicitudViaje
		errEsperado   error // nil = exito
		debePersistir bool
	}{
		{"viaje_id invalido rechazado", models.SolicitudViaje{ViajeID: 0}, service.ErrViajeIDInvalido, false},
		{"solicitud valida se persiste", models.SolicitudViaje{ViajeID: 1, PasajeroID: 2, Estado: "Pendiente"}, nil, true},
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
				repo.AssertNotCalled(t, "CrearSolicitudViaje") // la validacion corto antes
			} else {
				require.NoError(t, err)
				assert.Equal(t, 42, creado.ID)
				repo.AssertCalled(t, "CrearSolicitudViaje", c.entrada)
			}
		})
	}
}

// --- Obtener: comma-ok del repo traducido a error de dominio ---

func TestSolicitudViajeService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(solicitudViajeRepoMock)
		repo.On("BuscarSolicitudViajePorID", 1).Return(models.SolicitudViaje{ID: 1, ViajeID: 1, PasajeroID: 2, Estado: "Pendiente"}, true)
		s, err := service.NewSolicitudViajeService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, 1, s.ViajeID)
	})
	t.Run("no existe -> ErrSolicitudViajeNoEncontrado", func(t *testing.T) {
		repo := new(solicitudViajeRepoMock)
		repo.On("BuscarSolicitudViajePorID", 999).Return(models.SolicitudViaje{}, false)
		_, err := service.NewSolicitudViajeService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrSolicitudViajeNoEncontrado)
	})
}

// --- Actualizar: valida ANTES de tocar el repo, y mapea el no encontrado ---

func TestSolicitudViajeService_Actualizar(t *testing.T) {
	datos := models.SolicitudViaje{ViajeID: 2, PasajeroID: 3, Estado: "Aceptada"}

	t.Run("valido", func(t *testing.T) {
		repo := new(solicitudViajeRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarSolicitudViaje", 1, datos).Return(actualizado, true)
		s, err := service.NewSolicitudViajeService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, s.ID)
	})
	t.Run("no existe -> ErrSolicitudViajeNoEncontrado", func(t *testing.T) {
		repo := new(solicitudViajeRepoMock)
		repo.On("ActualizarSolicitudViaje", 999, datos).Return(models.SolicitudViaje{}, false)
		_, err := service.NewSolicitudViajeService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrSolicitudViajeNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(solicitudViajeRepoMock)
		_, err := service.NewSolicitudViajeService(repo).Actualizar(1, models.SolicitudViaje{ViajeID: 0})
		require.ErrorIs(t, err, service.ErrViajeIDInvalido)
		repo.AssertNotCalled(t, "ActualizarSolicitudViaje")
	})
}

// --- Borrar ---

func TestSolicitudViajeService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(solicitudViajeRepoMock)
		repo.On("BorrarSolicitudViaje", 1).Return(true)
		require.NoError(t, service.NewSolicitudViajeService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrSolicitudViajeNoEncontrado", func(t *testing.T) {
		repo := new(solicitudViajeRepoMock)
		repo.On("BorrarSolicitudViaje", 999).Return(false)
		require.ErrorIs(t, service.NewSolicitudViajeService(repo).Borrar(999), service.ErrSolicitudViajeNoEncontrado)
	})
}

// --- Listar: el service solo delega ---

func TestSolicitudViajeService_Listar(t *testing.T) {
	repo := new(solicitudViajeRepoMock)
	repo.On("ListarSolicitudViajes").Return([]models.SolicitudViaje{{ID: 1}, {ID: 2}})
	lista := service.NewSolicitudViajeService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
