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

func TestCrearViajeInmediatoExitoso(t *testing.T) {
	h, token := construirEntorno(t)

	body := `{
		"conductor_id":1,
		"origen":"Lima",
		"destino":"Callao",
		"hora_salida":"08:00",
		"cupos":4,
		"estado":"activo"
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/viajes-inmediatos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	require.Equal(t, http.StatusCreated, rec.Code)

	var creado models.ViajeInmediato
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))

	assert.NotZero(t, creado.ID)
	assert.Equal(t, 1, creado.ConductorID)
	assert.Equal(t, "Lima", creado.Origen)
	assert.Equal(t, "Callao", creado.Destino)
	assert.Equal(t, "08:00", creado.HoraSalida)
	assert.Equal(t, 4, creado.Cupos)
	assert.Equal(t, "activo", creado.Estado)
}

// TestObtenerProducto_NoEncontrado: id inexistente -> 404 Not Found.
func TestObtenerViajeInmediato_NoEncontrado(t *testing.T) {
	h, token := construirEntorno(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/viajes-inmediatos/9999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCrearViajeInmediato_Invalido(t *testing.T) {
	h, token := construirEntorno(t)
	body := `{
		"conductor_id":0,
		"origen":"Lima",
		"destino":"Callao",
		"hora_salida":"08:00",
		"cupos":4,
		"estado":"activo"
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/viajes-inmediatos", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// TestRutaProtegida_SinToken: sin header Authorization, el middleware corta
// antes de llegar al handler -> 401 Unauthorized.
func TestRutaViajeInmediatoProtegida_SinToken(t *testing.T) {
	h, _ := construirEntorno(t)
	body := `{"conductor_id":1,"origen":"Universidad","destino":"Centro","hora_salida":"08:00","cupos":4,"estado":"activo"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/viajes-inmediatos", strings.NewReader(body))
	// A proposito: NO seteamos Authorization.
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
