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
	return m.Called().Get(0).([]models.ViajeInmediato)
}
func (m *viajeInmediatoRepoMock) BuscarViajeInmediatoPorID(id int) (models.ViajeInmediato, bool) {
	a := m.Called(id)
	return a.Get(0).(models.ViajeInmediato), a.Bool(1)
}
func (m *viajeInmediatoRepoMock) CrearViajeInmediato(vi models.ViajeInmediato) models.ViajeInmediato {
	return m.Called(vi).Get(0).(models.ViajeInmediato)
}
func (m *viajeInmediatoRepoMock) ActualizarViajeInmediato(id int, datos models.ViajeInmediato) (models.ViajeInmediato, bool) {
	a := m.Called(id, datos)
	return a.Get(0).(models.ViajeInmediato), a.Bool(1)
}
func (m *viajeInmediatoRepoMock) BorrarViajeInmediato(id int) bool {
	return m.Called(id).Bool(0)
}

var _ storage.ViajeInmediatoRepository = (*viajeInmediatoRepoMock)(nil)

func TestViajeInmediatoService_Crear(t *testing.T) {
	casos := []struct {
		nombre        string
		entrada       models.ViajeInmediato
		errEsperado   error // nil = exito
		debePersistir bool
	}{
		{"conductor_id invalido rechazado", models.ViajeInmediato{ConductorID: 0, Origen: "ULEAM"}, service.ErrConductorIDInvalido, false},
		{"viaje valido se persiste", models.ViajeInmediato{ConductorID: 1, Origen: "ULEAM"}, nil, true},
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

// --- Obtener: comma-ok del repo traducido a error de dominio ---

func TestViajeInmediatoService_Obtener(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(viajeInmediatoRepoMock)
		repo.On("BuscarViajeInmediatoPorID", 1).Return(models.ViajeInmediato{ID: 1, Origen: "ULEAM"}, true)
		v, err := service.NewViajeInmediatoService(repo).Obtener(1)
		require.NoError(t, err)
		assert.Equal(t, "ULEAM", v.Origen)
	})
	t.Run("no existe -> ErrViajeInmediatoNoEncontrado", func(t *testing.T) {
		repo := new(viajeInmediatoRepoMock)
		repo.On("BuscarViajeInmediatoPorID", 999).Return(models.ViajeInmediato{}, false)
		_, err := service.NewViajeInmediatoService(repo).Obtener(999)
		require.ErrorIs(t, err, service.ErrViajeInmediatoNoEncontrado)
	})
}

// --- Actualizar: valida ANTES de tocar el repo, y mapea el no encontrado ---

func TestViajeInmediatoService_Actualizar(t *testing.T) {
	datos := models.ViajeInmediato{ConductorID: 1, Origen: "Terminal"}

	t.Run("valido", func(t *testing.T) {
		repo := new(viajeInmediatoRepoMock)
		actualizado := datos
		actualizado.ID = 1
		repo.On("ActualizarViajeInmediato", 1, datos).Return(actualizado, true)
		v, err := service.NewViajeInmediatoService(repo).Actualizar(1, datos)
		require.NoError(t, err)
		assert.Equal(t, 1, v.ID)
	})
	t.Run("no existe -> ErrViajeInmediatoNoEncontrado", func(t *testing.T) {
		repo := new(viajeInmediatoRepoMock)
		repo.On("ActualizarViajeInmediato", 999, datos).Return(models.ViajeInmediato{}, false)
		_, err := service.NewViajeInmediatoService(repo).Actualizar(999, datos)
		require.ErrorIs(t, err, service.ErrViajeInmediatoNoEncontrado)
	})
	t.Run("invalido no toca el repo", func(t *testing.T) {
		repo := new(viajeInmediatoRepoMock)
		_, err := service.NewViajeInmediatoService(repo).Actualizar(1, models.ViajeInmediato{})
		require.ErrorIs(t, err, service.ErrConductorIDInvalido)
		repo.AssertNotCalled(t, "ActualizarViajeInmediato")
	})
}

// --- Borrar ---

func TestViajeInmediatoService_Borrar(t *testing.T) {
	t.Run("existe", func(t *testing.T) {
		repo := new(viajeInmediatoRepoMock)
		repo.On("BorrarViajeInmediato", 1).Return(true)
		require.NoError(t, service.NewViajeInmediatoService(repo).Borrar(1))
	})
	t.Run("no existe -> ErrViajeInmediatoNoEncontrado", func(t *testing.T) {
		repo := new(viajeInmediatoRepoMock)
		repo.On("BorrarViajeInmediato", 999).Return(false)
		require.ErrorIs(t, service.NewViajeInmediatoService(repo).Borrar(999), service.ErrViajeInmediatoNoEncontrado)
	})
}

// --- Listar: el service solo delega ---

func TestViajeInmediatoService_Listar(t *testing.T) {
	repo := new(viajeInmediatoRepoMock)
	repo.On("ListarViajeInmediatos").Return([]models.ViajeInmediato{{ID: 1}, {ID: 2}})
	lista := service.NewViajeInmediatoService(repo).Listar()
	assert.Len(t, lista, 2)
	repo.AssertExpectations(t)
}
