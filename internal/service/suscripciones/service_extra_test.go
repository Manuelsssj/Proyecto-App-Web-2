package service

import (
	"testing"

	"github.com/stretchr/testify/require"

	models "RideUleam/internal/models/suscripciones"
)

type fakeAlmacenServiceExtra struct {
	suscripciones []models.SuscripcionRuta
	planes        []models.PlanPago
	historial     []models.HistorialSuscripcion
}

func (f *fakeAlmacenServiceExtra) ListarSuscripciones() ([]models.SuscripcionRuta, error) {
	return f.suscripciones, nil
}
func (f *fakeAlmacenServiceExtra) ObtenerSuscripcion(id int) (models.SuscripcionRuta, error) {
	return models.SuscripcionRuta{ID: uint(id), RutaID: 1, UsuarioID: 1}, nil
}
func (f *fakeAlmacenServiceExtra) CrearSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	s.ID = 1
	f.suscripciones = append(f.suscripciones, s)
	return s, nil
}
func (f *fakeAlmacenServiceExtra) ActualizarSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	return s, nil
}
func (f *fakeAlmacenServiceExtra) EliminarSuscripcion(id int) error {
	return nil
}

func (f *fakeAlmacenServiceExtra) ListarPlanes() ([]models.PlanPago, error) {
	return f.planes, nil
}
func (f *fakeAlmacenServiceExtra) ObtenerPlan(id int) (models.PlanPago, error) {
	return models.PlanPago{ID: uint(id), RutaID: 1, ValorSemanal: 5}, nil
}
func (f *fakeAlmacenServiceExtra) CrearPlan(p models.PlanPago) (models.PlanPago, error) {
	p.ID = 1
	f.planes = append(f.planes, p)
	return p, nil
}
func (f *fakeAlmacenServiceExtra) ActualizarPlan(p models.PlanPago) (models.PlanPago, error) {
	return p, nil
}
func (f *fakeAlmacenServiceExtra) EliminarPlan(id int) error {
	return nil
}

func (f *fakeAlmacenServiceExtra) ListarHistorial() ([]models.HistorialSuscripcion, error) {
	return f.historial, nil
}
func (f *fakeAlmacenServiceExtra) ObtenerHistorial(id int) (models.HistorialSuscripcion, error) {
	return models.HistorialSuscripcion{ID: uint(id), SuscripcionID: 1, Estado: "activa"}, nil
}
func (f *fakeAlmacenServiceExtra) CrearHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	h.ID = 1
	f.historial = append(f.historial, h)
	return h, nil
}
func (f *fakeAlmacenServiceExtra) ActualizarHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	return h, nil
}
func (f *fakeAlmacenServiceExtra) EliminarHistorial(id int) error {
	return nil
}

func TestPlanService_CrearValido(t *testing.T) {
	srv := NewPlanService(&fakeAlmacenServiceExtra{})

	plan, err := srv.Crear(models.PlanPago{
		RutaID:       1,
		ValorSemanal: 10.50,
	})

	require.NoError(t, err)
	require.Equal(t, uint(1), plan.ID)
	require.Equal(t, 10.50, plan.ValorSemanal)
}

func TestPlanService_CrearInvalido(t *testing.T) {
	srv := NewPlanService(&fakeAlmacenServiceExtra{})

	_, err := srv.Crear(models.PlanPago{
		RutaID:       0,
		ValorSemanal: 10.50,
	})

	require.ErrorIs(t, err, ErrDatosInvalidos)
}

func TestPlanService_ActualizarInvalido(t *testing.T) {
	srv := NewPlanService(&fakeAlmacenServiceExtra{})

	_, err := srv.Actualizar(models.PlanPago{
		ID:           0,
		RutaID:       1,
		ValorSemanal: 5,
	})

	require.ErrorIs(t, err, ErrDatosInvalidos)
}

func TestSuscripcionService_CrearValida(t *testing.T) {
	srv := NewSuscripcionService(&fakeAlmacenServiceExtra{})

	sub, err := srv.Crear(models.SuscripcionRuta{
		RutaID:    1,
		UsuarioID: 1,
	})

	require.NoError(t, err)
	require.Equal(t, uint(1), sub.ID)
}

func TestSuscripcionService_CrearInvalida(t *testing.T) {
	srv := NewSuscripcionService(&fakeAlmacenServiceExtra{})

	_, err := srv.Crear(models.SuscripcionRuta{
		RutaID:    0,
		UsuarioID: 1,
	})

	require.ErrorIs(t, err, ErrDatosInvalidos)
}

func TestSuscripcionService_EliminarInvalido(t *testing.T) {
	srv := NewSuscripcionService(&fakeAlmacenServiceExtra{})

	err := srv.Eliminar(0)

	require.ErrorIs(t, err, ErrDatosInvalidos)
}

func TestHistorialService_CrearValidoConFechaAutomatica(t *testing.T) {
	srv := NewHistorialService(&fakeAlmacenServiceExtra{})

	historial, err := srv.Crear(models.HistorialSuscripcion{
		SuscripcionID: 1,
		Estado:        "activa",
	})

	require.NoError(t, err)
	require.Equal(t, uint(1), historial.ID)
	require.NotEmpty(t, historial.FechaRegistro)
}

func TestHistorialService_CrearInvalido(t *testing.T) {
	srv := NewHistorialService(&fakeAlmacenServiceExtra{})

	_, err := srv.Crear(models.HistorialSuscripcion{
		SuscripcionID: 0,
		Estado:        "activa",
	})

	require.ErrorIs(t, err, ErrDatosInvalidos)
}

func TestHistorialService_ActualizarInvalido(t *testing.T) {
	srv := NewHistorialService(&fakeAlmacenServiceExtra{})

	_, err := srv.Actualizar(models.HistorialSuscripcion{
		ID:            0,
		SuscripcionID: 1,
		Estado:        "activa",
	})

	require.ErrorIs(t, err, ErrDatosInvalidos)
}

///
