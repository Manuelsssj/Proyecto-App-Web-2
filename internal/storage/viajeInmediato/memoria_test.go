// Tests del almacen en memoria con la libreria ESTANDAR (testing), sin testify.
// Es el estilo mas basico de Go y el que conviene mostrar primero: t.Fatalf
// para abortar, t.Errorf para seguir. Sin dependencias externas.
package storage

import (
	models "RideUleam/internal/models/viajeInmediato"
	"testing"
)

func TestMemoria_CrearYBuscarViajeInmediato(t *testing.T) {
	m := NuevaMemoria()

	creado := m.CrearViajeInmediato(models.ViajeInmediato{
		ConductorID: 1,
		Origen:      "Universidad ULEAM",
		Destino:     "Terminal Terrestre",
		HoraSalida:  "08:30",
		Cupos:       4,
		Estado:      "Disponible",
	})

	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrado, ok := m.BuscarViajeInmediatoPorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontró el viaje recién creado (id=%d)", creado.ID)
	}

	if encontrado.Origen != "Universidad ULEAM" {
		t.Errorf("origen = %q; esperaba %q", encontrado.Origen, "Universidad ULEAM")
	}

}

func TestMemoria_BuscarInexistenteViajeInmediato(t *testing.T) {
	m := NuevaMemoria()

	// El patron comma-ok: ok debe ser false para un id que no existe.
	if _, ok := m.BuscarViajeInmediatoPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoria_ActualizarYBorrarViajeInmediato(t *testing.T) {
	m := NuevaMemoria()
	creado := m.CrearViajeInmediato(models.ViajeInmediato{
		ConductorID: 1,
		Origen:      "Universidad ULEAM",
		Destino:     "Terminal Terrestre",
		HoraSalida:  "08:30",
		Cupos:       4,
		Estado:      "Disponible",
	})

	if _, ok := m.ActualizarViajeInmediato(creado.ID, models.ViajeInmediato{
		ConductorID: 1,
		Origen:      "Universidad ULEAM23",
		Destino:     "Terminal Terrestre",
		HoraSalida:  "08:30",
		Cupos:       4,
		Estado:      "Disponible",
	}); !ok {
		t.Fatalf("no se pudo actualizar el ViajeInmediato id=%d", creado.ID)
	}

	if !m.BorrarViajeInmediato(creado.ID) {
		t.Errorf("esperaba poder borrar el ViajeInmediato id=%d", creado.ID)
	}
	if _, ok := m.BuscarViajeInmediatoPorID(creado.ID); ok {
		t.Errorf("el ViajeInmediato id=%d deberia haber sido borrado", creado.ID)
	}
}
