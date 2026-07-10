package handlers_test

import (
	models "RideUleam/internal/models/usuarioVehiculo"
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

func TestListarVehiculos_OK(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/vehiculos", "", token))

	require.Equal(t, http.StatusOK, rec.Code)
	var lista []models.Vehiculo
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&lista))
	assert.Len(t, lista, 1) // el producto sembrado
}

func TestObtenerVehiculo(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("existe -> 200", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/vehiculos/1", "", token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/vehiculos/9999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
	t.Run("id no numerico -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/vehiculos/abc", "", token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestCrearVehiculo(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 201", func(t *testing.T) {
		body := `{"conductor_id":1,"placa":"MBT-456","marca":"Chevrolet","modelo":"Spark","capacidad":4}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/vehiculos", body, token))
		require.Equal(t, http.StatusCreated, rec.Code)
		var creado models.Vehiculo
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&creado))
		assert.NotZero(t, creado.ID)
	})
	t.Run("conductor_id invalido -> 400", func(t *testing.T) {
		body := `{"conductor_id":0,"placa":"ABC-123","capacidad":4}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/vehiculos", body, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	t.Run("JSON malformado -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/vehiculos", `{roto`, token))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestActualizarVehiculo(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("valido -> 200", func(t *testing.T) {
		body := `{"conductor_id":2,"placa":"ABC-123","marca":"Toyota","modelo":"Corolla","capacidad":5}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/vehiculos/1", body, token))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		body := `{"conductor_id":1,"placa":"XYZ-999","capacidad":4}`
		rec := ejecutar(h, jsonReq(http.MethodPut, "/api/v1/vehiculos/9999", body, token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestBorrarVehiculo(t *testing.T) {
	h, _, _ := construirEntorno()
	token := tokenValido(t, h)

	t.Run("existe -> 204", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/vehiculos/1", "", token))
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
	t.Run("no existe -> 404", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodDelete, "/api/v1/vehiculos/9999", "", token))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

// El corazon de la seguridad: el middleware corta ANTES del handler.
func TestRutaProtegida_SinToken(t *testing.T) {
	h, _, _ := construirEntorno()
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/vehiculos", "", "")) // sin Bearer
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestRutaProtegida_TokenInvalido(t *testing.T) {
	h, _, _ := construirEntorno()
	rec := ejecutar(h, jsonReq(http.MethodGet, "/api/v1/vehiculos", "", "token.falso.123"))
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
