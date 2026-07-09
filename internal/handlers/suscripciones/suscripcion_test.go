package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"suscripciones-api/internal/middleware"
	"suscripciones-api/internal/models/suscripciones"
	"suscripciones-api/internal/service/suscripciones"
)

type fakeAlmacenHandler struct {
	suscripciones []models.SuscripcionRuta
}

func (f *fakeAlmacenHandler) ListarSuscripciones() ([]models.SuscripcionRuta, error) {
	return f.suscripciones, nil
}

func (f *fakeAlmacenHandler) ObtenerSuscripcion(id int) (models.SuscripcionRuta, error) {
	for _, s := range f.suscripciones {
		if s.ID == uint(id) {
			return s, nil
		}
	}
	return models.SuscripcionRuta{}, service.ErrNoEncontrado
}

func (f *fakeAlmacenHandler) CrearSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	s.ID = uint(len(f.suscripciones) + 1)
	f.suscripciones = append(f.suscripciones, s)
	return s, nil
}

func (f *fakeAlmacenHandler) ActualizarSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	for i := range f.suscripciones {
		if f.suscripciones[i].ID == s.ID {
			f.suscripciones[i] = s
			return s, nil
		}
	}
	return models.SuscripcionRuta{}, service.ErrNoEncontrado
}

func (f *fakeAlmacenHandler) EliminarSuscripcion(id int) error {
	for i := range f.suscripciones {
		if f.suscripciones[i].ID == uint(id) {
			f.suscripciones = append(f.suscripciones[:i], f.suscripciones[i+1:]...)
			return nil
		}
	}
	return service.ErrNoEncontrado
}

func (f *fakeAlmacenHandler) ListarPlanes() ([]models.PlanPago, error) {
	return nil, nil
}

func (f *fakeAlmacenHandler) ObtenerPlan(id int) (models.PlanPago, error) {
	return models.PlanPago{}, nil
}

func (f *fakeAlmacenHandler) CrearPlan(p models.PlanPago) (models.PlanPago, error) {
	return p, nil
}

func (f *fakeAlmacenHandler) ActualizarPlan(p models.PlanPago) (models.PlanPago, error) {
	return p, nil
}

func (f *fakeAlmacenHandler) EliminarPlan(id int) error {
	return nil
}

func (f *fakeAlmacenHandler) ListarHistorial() ([]models.HistorialSuscripcion, error) {
	return nil, nil
}

func (f *fakeAlmacenHandler) ObtenerHistorial(id int) (models.HistorialSuscripcion, error) {
	return models.HistorialSuscripcion{}, nil
}

func (f *fakeAlmacenHandler) CrearHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	return h, nil
}

func (f *fakeAlmacenHandler) ActualizarHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	return h, nil
}

func (f *fakeAlmacenHandler) EliminarHistorial(id int) error {
	return nil
}

type fakeUserRepoHandler struct{}

func (f *fakeUserRepoHandler) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	return u, nil
}

func (f *fakeUserRepoHandler) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	return models.Usuario{}, false
}

func TestSuscripcionHandler_CrearSuscripcion(t *testing.T) {
	fake := &fakeAlmacenHandler{}
	srv := service.NewSuscripcionService(fake)
	h := NewSuscripcionHandler(srv)

	r := chi.NewRouter()
	h.Registrar(r)

	body := bytes.NewBufferString(`{"ruta_id":1,"usuario_id":1}`)
	req := httptest.NewRequest(http.MethodPost, "/suscripciones", body)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("se esperaba status 201, se obtuvo %d. body: %s", rr.Code, rr.Body.String())
	}

	if len(fake.suscripciones) != 1 {
		t.Fatalf("se esperaba 1 suscripción guardada, se obtuvo %d", len(fake.suscripciones))
	}
}

func TestSuscripcionHandler_RutaProtegidaSinTokenRetorna401(t *testing.T) {
	fake := &fakeAlmacenHandler{}
	srv := service.NewSuscripcionService(fake)
	h := NewSuscripcionHandler(srv)

	r := chi.NewRouter()
	authSrv := service.NewAuthService(&fakeUserRepoHandler{})

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(authSrv))
		h.Registrar(r)
	})

	req := httptest.NewRequest(http.MethodGet, "/suscripciones", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba status 401, se obtuvo %d. body: %s", rr.Code, rr.Body.String())
	}
}

//go test ./internal/handlers -v
