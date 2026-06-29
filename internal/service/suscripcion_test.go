package service

import (
	"errors"
	"testing"

	"suscripciones-api/internal/models"
)

type mockAlmacenSuscripcion struct {
	crearLlamado bool
}

func (m *mockAlmacenSuscripcion) ListarSuscripciones() ([]models.SuscripcionRuta, error) {
	return nil, nil
}

func (m *mockAlmacenSuscripcion) ObtenerSuscripcion(id int) (models.SuscripcionRuta, error) {
	return models.SuscripcionRuta{}, nil
}

func (m *mockAlmacenSuscripcion) CrearSuscripcion(sub models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	m.crearLlamado = true
	sub.ID = 1
	return sub, nil
}

func (m *mockAlmacenSuscripcion) ActualizarSuscripcion(sub models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	return sub, nil
}

func (m *mockAlmacenSuscripcion) EliminarSuscripcion(id int) error {
	return nil
}

func (m *mockAlmacenSuscripcion) ListarPlanes() ([]models.PlanPago, error) {
	return nil, nil
}

func (m *mockAlmacenSuscripcion) ObtenerPlan(id int) (models.PlanPago, error) {
	return models.PlanPago{}, nil
}

func (m *mockAlmacenSuscripcion) CrearPlan(p models.PlanPago) (models.PlanPago, error) {
	return p, nil
}

func (m *mockAlmacenSuscripcion) ActualizarPlan(p models.PlanPago) (models.PlanPago, error) {
	return p, nil
}

func (m *mockAlmacenSuscripcion) EliminarPlan(id int) error {
	return nil
}

func (m *mockAlmacenSuscripcion) ListarHistorial() ([]models.HistorialSuscripcion, error) {
	return nil, nil
}

func (m *mockAlmacenSuscripcion) ObtenerHistorial(id int) (models.HistorialSuscripcion, error) {
	return models.HistorialSuscripcion{}, nil
}

func (m *mockAlmacenSuscripcion) CrearHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	return h, nil
}

func (m *mockAlmacenSuscripcion) ActualizarHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	return h, nil
}

func (m *mockAlmacenSuscripcion) EliminarHistorial(id int) error {
	return nil
}

func TestSuscripcionService_CrearRechazaDatosInvalidos(t *testing.T) {
	mock := &mockAlmacenSuscripcion{}
	srv := NewSuscripcionService(mock)

	_, err := srv.Crear(models.SuscripcionRuta{
		RutaID:    0,
		UsuarioID: 1,
	})

	if !errors.Is(err, ErrDatosInvalidos) {
		t.Fatalf("se esperaba ErrDatosInvalidos, se obtuvo: %v", err)
	}

	if mock.crearLlamado {
		t.Fatal("no debía llamarse al repositorio cuando los datos son inválidos")
	}
}

func TestSuscripcionService_CrearValidaGuardaSuscripcion(t *testing.T) {
	mock := &mockAlmacenSuscripcion{}
	srv := NewSuscripcionService(mock)

	sub, err := srv.Crear(models.SuscripcionRuta{
		RutaID:    1,
		UsuarioID: 1,
	})

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}

	if !mock.crearLlamado {
		t.Fatal("se esperaba que el repositorio fuera llamado")
	}

	if sub.ID != 1 {
		t.Fatalf("se esperaba ID 1, se obtuvo: %d", sub.ID)
	}
}
