// Tests del almacen en memoria con la libreria ESTANDAR (testing), sin testify.
// Es el estilo mas basico de Go y el que conviene mostrar primero: t.Fatalf
// para abortar, t.Errorf para seguir. Sin dependencias externas.
package storage

import (
	models "RideUleam/internal/models/usuarioVehiculo"
	"testing"
)

func TestMemoria_CrearYBuscar(t *testing.T) {
	m := NuevaMemoria()

	creado := m.CrearVehiculo(models.Vehiculo{
		ConductorID: 1,
		Placa:       "ABC123",
		Marca:       "Toyota",
		Modelo:      "Corolla",
		Capacidad:   5,
	})

	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrado, ok := m.BuscarVehiculoPorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontro el vehiculo recien creado (id=%d)", creado.ID)
	}

	if encontrado.Placa != "ABC123" {
		t.Errorf("placa = %q; esperaba %q", encontrado.Placa, "ABC123")
	}
}

func TestMemoria_BuscarInexistente(t *testing.T) {
	m := NuevaMemoria()

	// El patron comma-ok: ok debe ser false para un id que no existe.
	if _, ok := m.BuscarVehiculoPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrar(t *testing.T) {
	m := NuevaMemoria()
	creado := m.CrearVehiculo(models.Vehiculo{
		ConductorID: 1,
		Placa:       "ABC123",
		Marca:       "Toyota",
		Modelo:      "Corolla",
		Capacidad:   5,
	})

	if _, ok := m.ActualizarVehiculo(creado.ID, models.Vehiculo{
		ConductorID: 1,
		Placa:       "ABC123",
		Marca:       "Toyota23",
		Modelo:      "Corolla",
		Capacidad:   5,
	}); !ok {
		t.Fatalf("no se pudo actualizar el Vehiculo id=%d", creado.ID)
	}

	if !m.BorrarVehiculo(creado.ID) {
		t.Errorf("esperaba poder borrar el Vehiculo id=%d", creado.ID)
	}
	if _, ok := m.BuscarVehiculoPorID(creado.ID); ok {
		t.Errorf("el Vehiculo id=%d deberia haber sido borrado", creado.ID)
	}
}
