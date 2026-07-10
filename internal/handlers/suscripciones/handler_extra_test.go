package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	models "RideUleam/internal/models/suscripciones"
	service "RideUleam/internal/service/suscripciones"
)

type fakeAlmacenHandlerExtra struct{}

func (f *fakeAlmacenHandlerExtra) ListarSuscripciones() ([]models.SuscripcionRuta, error) {
	return []models.SuscripcionRuta{{ID: 1, RutaID: 1, UsuarioID: 1}}, nil
}
func (f *fakeAlmacenHandlerExtra) ObtenerSuscripcion(id int) (models.SuscripcionRuta, error) {
	return models.SuscripcionRuta{ID: uint(id), RutaID: 1, UsuarioID: 1}, nil
}
func (f *fakeAlmacenHandlerExtra) CrearSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	s.ID = 1
	return s, nil
}
func (f *fakeAlmacenHandlerExtra) ActualizarSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	return s, nil
}
func (f *fakeAlmacenHandlerExtra) EliminarSuscripcion(id int) error { return nil }

func (f *fakeAlmacenHandlerExtra) ListarPlanes() ([]models.PlanPago, error) {
	return []models.PlanPago{{ID: 1, RutaID: 1, ValorSemanal: 5}}, nil
}
func (f *fakeAlmacenHandlerExtra) ObtenerPlan(id int) (models.PlanPago, error) {
	return models.PlanPago{ID: uint(id), RutaID: 1, ValorSemanal: 5}, nil
}
func (f *fakeAlmacenHandlerExtra) CrearPlan(p models.PlanPago) (models.PlanPago, error) {
	p.ID = 1
	return p, nil
}
func (f *fakeAlmacenHandlerExtra) ActualizarPlan(p models.PlanPago) (models.PlanPago, error) {
	return p, nil
}
func (f *fakeAlmacenHandlerExtra) EliminarPlan(id int) error { return nil }

func (f *fakeAlmacenHandlerExtra) ListarHistorial() ([]models.HistorialSuscripcion, error) {
	return []models.HistorialSuscripcion{{ID: 1, SuscripcionID: 1, Estado: "activa"}}, nil
}
func (f *fakeAlmacenHandlerExtra) ObtenerHistorial(id int) (models.HistorialSuscripcion, error) {
	return models.HistorialSuscripcion{ID: uint(id), SuscripcionID: 1, Estado: "activa"}, nil
}
func (f *fakeAlmacenHandlerExtra) CrearHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	h.ID = 1
	return h, nil
}
func (f *fakeAlmacenHandlerExtra) ActualizarHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	return h, nil
}
func (f *fakeAlmacenHandlerExtra) EliminarHistorial(id int) error { return nil }

func TestPlanHandler_CRUD(t *testing.T) {
	srv := service.NewPlanService(&fakeAlmacenHandlerExtra{})
	h := NewPlanHandler(srv)

	r := chi.NewRouter()
	h.Registrar(r)

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		code   int
	}{
		{"listar", http.MethodGet, "/planes", "", http.StatusOK},
		{"obtener", http.MethodGet, "/planes/1", "", http.StatusOK},
		{"crear", http.MethodPost, "/planes", `{"ruta_id":1,"valor_semanal":5}`, http.StatusCreated},
		{"actualizar", http.MethodPut, "/planes/1", `{"ruta_id":2,"valor_semanal":7}`, http.StatusOK},
		{"eliminar", http.MethodDelete, "/planes/1", "", http.StatusNoContent},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(tc.method, tc.path, bytes.NewBufferString(tc.body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, tc.code, rr.Code, tc.name)
	}
}

func TestHistorialHandler_CRUD(t *testing.T) {
	srv := service.NewHistorialService(&fakeAlmacenHandlerExtra{})
	h := NewHistorialHandler(srv)

	r := chi.NewRouter()
	h.Registrar(r)

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		code   int
	}{
		{"listar", http.MethodGet, "/historial", "", http.StatusOK},
		{"obtener", http.MethodGet, "/historial/1", "", http.StatusOK},
		{"crear", http.MethodPost, "/historial", `{"suscripcion_id":1,"estado":"activa"}`, http.StatusCreated},
		{"actualizar", http.MethodPut, "/historial/1", `{"suscripcion_id":1,"estado":"cancelada"}`, http.StatusOK},
		{"eliminar", http.MethodDelete, "/historial/1", "", http.StatusNoContent},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(tc.method, tc.path, bytes.NewBufferString(tc.body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, tc.code, rr.Code, tc.name)
	}
}

func TestSuscripcionHandler_JSONInvalido(t *testing.T) {
	srv := service.NewSuscripcionService(&fakeAlmacenHandlerExtra{})
	h := NewSuscripcionHandler(srv)

	r := chi.NewRouter()
	h.Registrar(r)

	req := httptest.NewRequest(http.MethodPost, "/suscripciones", bytes.NewBufferString(`{json malo`))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPlanHandler_IDInvalido(t *testing.T) {
	srv := service.NewPlanService(&fakeAlmacenHandlerExtra{})
	h := NewPlanHandler(srv)

	r := chi.NewRouter()
	h.Registrar(r)

	req := httptest.NewRequest(http.MethodGet, "/planes/abc", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
}
