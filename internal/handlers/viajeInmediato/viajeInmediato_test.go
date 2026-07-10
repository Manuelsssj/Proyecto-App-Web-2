package handlers_test

import (
	models "RideUleam/internal/models/viajeInmediato"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ejecutar corre una peticion contra el handler y devuelve el recorder.
func ejecutar(h http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestListarViajeInmediatos_OK(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/viajes-inmediatos", "", token))

	require.Equal(t, http.StatusOK, rec.Code)
	var lista []models.ViajeInmediato
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.Len(t, lista, 1) // el producto sembrado
}

func TestObtenerViajeInmediato(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("existe -> 200", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/viajes-inmediatos/1", "", token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/viajes-inmediatos/9999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
	t.Run("id no numerico -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/viajes-inmediatos/abc", "", token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestCrearViajeInmediato(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 201", func(t *testing.T) {
		body := `{"conductor_id":1,"origen":"ULEAM","destino":"Centro","hora_salida":"08:00","cupos":4,"estado":"Disponible"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/viajes-inmediatos", body, token))
		require.Equal(t, http.StatusCreated, rec.Code)
		var creado models.ViajeInmediato
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
		assert.NotZero(t, creado.ID)
	})
	t.Run("conductor_id invalido -> 400", func(t *testing.T) {
		body := `{"conductor_id":0,"origen":"ULEAM","destino":"Centro"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/viajes-inmediatos", body, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("JSON malformado -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/viajes-inmediatos", `{roto`, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestActualizarViajeInmediato(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 200", func(t *testing.T) {
		body := `{"conductor_id":2,"origen":"Terminal","destino":"ULEAM","hora_salida":"10:00","cupos":3,"estado":"Disponible"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/viajes-inmediatos/1", body, token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		body := `{"conductor_id":1,"origen":"A","destino":"B"}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/viajes-inmediatos/9999", body, token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestBorrarViajeInmediato(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("existe -> 204", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/viajes-inmediatos/1", "", token))
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/viajes-inmediatos/9999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

// El corazon de la seguridad: el middleware corta ANTES del handler.
func TestViajeInmediato_SinToken(t *testing.T) {
	h, _, _ := construirEntorno()
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/viajes-inmediatos", "", "")) // sin Bearer
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestViajeInmediato_TokenInvalido(t *testing.T) {
	h, _, _ := construirEntorno()
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/viajes-inmediatos", "", "token.falso.123"))
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
