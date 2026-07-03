package handlers_test

import (
	models "RideUleam/internal/models/viajeInmediato"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// construirEntorno arma el MISMO router que main.go (mismas rutas, mismo
// middleware.Auth real) pero con almacen en memoria y repo de usuarios fake.
// Devuelve el handler listo para httptest y un token valido ya emitido.
//
// Clave pedagogica: probamos a traves del middleware REAL, no de uno simplificado.
// Si el wiring de la ruta protegida se rompe, este test se entera.

// registrarYObtenerToken hace register + login contra el propio router para
// conseguir un JWT valido, igual que lo haria un cliente real.

// TestCrearProducto_Exitoso: POST con token y cuerpo valido -> 201 Created.

func TestCrearSolicitudViajeExitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"viaje_id":1,
		"pasajero_id":2,
		"estado":"pendiente"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/solicitudes-viajes", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creada models.SolicitudViaje
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creada))

	assert.NotZero(t, creada.ID)
	assert.Equal(t, 1, creada.ViajeID)
	assert.Equal(t, 2, creada.PasajeroID)
	assert.Equal(t, "pendiente", creada.Estado)
}

// TestCrearProducto_Exitoso: POST con token y cuerpo valido -> 201 Created.

// TestObtenerProducto_NoEncontrado: id inexistente -> 404 Not Found.
func TestObtenerSolicitudViaje_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/solicitudes-viajes/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

// TestCrearProducto_Invalido: cuerpo que viola la regla de negocio -> 400.
func TestCrearSolicitudViaje_Invalido(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{"viaje_id":0,"pasajero_id":1,"estado":"pendiente"}` // viaje_id invalido

	req := httptest.NewRequest(http.MethodPost, "/api/v1/solicitudes-viajes", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// TestRutaProtegida_SinToken: sin header Authorization, el middleware corta
// antes de llegar al handler -> 401 Unauthorized.
func TestRutaSolicitudViajeProtegida_SinToken(t *testing.T) {
	h, _ := construirEntorno(t)
	body := `{"viaje_id":1,"pasajero_id":1,"estado":"pendiente"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/solicitudes-viajes", strings.NewReader(body))
	// A proposito: NO seteamos Authorization.
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
