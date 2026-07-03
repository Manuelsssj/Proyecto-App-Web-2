package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mw "cmd/rideUleam/internal/middleware"
	models "cmd/rideUleam/internal/models/rutaProgramada"
	storage "cmd/rideUleam/internal/storage/rutaProgramada"

	"github.com/go-chi/chi/v5"
)

func prepararRouterDePrueba() *chi.Mux {
	almacen := storage.NuevaMemoria()
	userRepo := almacen

	servidor := NewServer(almacen, userRepo)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", servidor.Registrar)
		r.Post("/auth/login", servidor.Login)

		r.Group(func(r chi.Router) {
			r.Use(mw.AuthJWT(servidor.Auth))

			r.Get("/rutas-programadas", servidor.ListarRutasProgramadas)
			r.Post("/rutas-programadas", servidor.CrearRutaProgramada)
		})
	})

	return r
}

func obtenerTokenDePrueba(t *testing.T, r http.Handler) string {
	t.Helper()

	usuario := credenciales{
		Email:    "test@uleam.com",
		Password: "12345",
	}

	bodyRegistro, _ := json.Marshal(usuario)

	reqRegistro := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/register",
		bytes.NewBuffer(bodyRegistro),
	)
	reqRegistro.Header.Set("Content-Type", "application/json")

	recRegistro := httptest.NewRecorder()
	r.ServeHTTP(recRegistro, reqRegistro)

	if recRegistro.Code != http.StatusCreated {
		t.Fatalf("se esperaba status 201 en registro, se obtuvo %d. Respuesta: %s", recRegistro.Code, recRegistro.Body.String())
	}

	login := credenciales{
		Email:    "test@uleam.com",
		Password: "12345",
	}

	bodyLogin, _ := json.Marshal(login)

	reqLogin := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/login",
		bytes.NewBuffer(bodyLogin),
	)
	reqLogin.Header.Set("Content-Type", "application/json")

	recLogin := httptest.NewRecorder()
	r.ServeHTTP(recLogin, reqLogin)

	if recLogin.Code != http.StatusOK {
		t.Fatalf("se esperaba status 200 en login, se obtuvo %d. Respuesta: %s", recLogin.Code, recLogin.Body.String())
	}

	var respuesta map[string]string
	if err := json.NewDecoder(recLogin.Body).Decode(&respuesta); err != nil {
		t.Fatalf("no se pudo leer la respuesta del login: %v", err)
	}

	token := respuesta["token"]
	if token == "" {
		t.Fatal("se esperaba recibir un token en el login")
	}

	return token
}

func TestRutaProgramadaHandler_SinTokenResponde401(t *testing.T) {
	r := prepararRouterDePrueba()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/rutas-programadas",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("se esperaba status 401, se obtuvo %d", rec.Code)
	}
}

func TestRutaProgramadaHandler_CrearRutaConTokenResponde201(t *testing.T) {
	r := prepararRouterDePrueba()
	token := obtenerTokenDePrueba(t, r)

	ruta := models.RutaProgramada{
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "ULEAM",
		Costo:       0.75,
	}

	body, _ := json.Marshal(ruta)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/rutas-programadas",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("se esperaba status 201, se obtuvo %d. Respuesta: %s", rec.Code, rec.Body.String())
	}

	var respuesta models.RutaProgramada
	if err := json.NewDecoder(rec.Body).Decode(&respuesta); err != nil {
		t.Fatalf("no se pudo leer la respuesta: %v", err)
	}

	if respuesta.ID == 0 {
		t.Fatal("se esperaba que la ruta creada tenga ID")
	}

	if respuesta.Origen != "Los Esteros" {
		t.Errorf("origen esperado Los Esteros, obtenido %s", respuesta.Origen)
	}
}
