package storage

import (
	models "RideUleam/internal/models/viajeInmediato"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// nuevaDBPrueba abre una SQLite en memoria, migra el esquema y la devuelve.
// SetMaxOpenConns(1) garantiza que migracion y consultas usen la MISMA conexion
// (con ":memory:" cada conexion tendria su propia base, vacia).
func nuevaDBPrueba(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir la base de prueba: %v", err)
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		t.Fatalf("no se pudo obtener *sql.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(1)
	if err := gdb.AutoMigrate(&models.ViajeInmediato{}, &models.SolicitudViaje{}, &models.ParticipanteViaje{}); err != nil {
		t.Fatalf("fallo AutoMigrate: %v", err)
	}
	return gdb
}

func TestSQLite_ParticipanteViajeCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// Crear: GORM debe asignar el ID autogenerado.
	creado := alm.CrearParticipanteViaje(models.ParticipanteViaje{
		ViajeID:   1,
		UsuarioID: 2,
	})

	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado por la base, obtuve 0")
	}

	// Buscar el recién creado.
	encontrado, ok := alm.BuscarParticipanteViajePorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontro el ParticipanteViaje id=%d", creado.ID)
	}

	if encontrado.ViajeID != 1 {
		t.Errorf("viaje_id = %d; esperaba %d", encontrado.ViajeID, 1)
	}

	// Actualizar.
	if _, ok := alm.ActualizarParticipanteViaje(creado.ID, models.ParticipanteViaje{
		ViajeID:   2,
		UsuarioID: 5,
	}); !ok {
		t.Fatalf("no se pudo actualizar el ParticipanteViaje id=%d", creado.ID)
	}

	// Borrar y confirmar que ya no está.
	if !alm.BorrarParticipanteViaje(creado.ID) {
		t.Errorf("esperaba poder borrar el ParticipanteViaje id=%d", creado.ID)
	}

	if _, ok := alm.BuscarParticipanteViajePorID(creado.ID); ok {
		t.Errorf("el ParticipanteViaje id=%d deberia haber sido borrado", creado.ID)
	}
}

func TestSQLite_BuscarParticipanteViajeInexistente(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// El error de GORM (registro no encontrado) se traduce a comma-ok = false.
	if _, ok := alm.BuscarParticipanteViajePorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestSQLite_SolicitudViajeCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// Crear: GORM debe asignar el ID autogenerado.
	creado := alm.CrearSolicitudViaje(models.SolicitudViaje{
		ViajeID:    1,
		PasajeroID: 2,
		Estado:     "Pendiente",
	})

	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado por la base, obtuve 0")
	}

	// Buscar el recién creado.
	encontrado, ok := alm.BuscarSolicitudViajePorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontro la SolicitudViaje id=%d", creado.ID)
	}

	if encontrado.Estado != "Pendiente" {
		t.Errorf("estado = %q; esperaba %q", encontrado.Estado, "Pendiente")
	}

	// Actualizar.
	if _, ok := alm.ActualizarSolicitudViaje(creado.ID, models.SolicitudViaje{
		ViajeID:    2,
		PasajeroID: 5,
		Estado:     "Aceptada",
	}); !ok {
		t.Fatalf("no se pudo actualizar la SolicitudViaje id=%d", creado.ID)
	}

	// Borrar y confirmar que ya no está.
	if !alm.BorrarSolicitudViaje(creado.ID) {
		t.Errorf("esperaba poder borrar la SolicitudViaje id=%d", creado.ID)
	}

	if _, ok := alm.BuscarSolicitudViajePorID(creado.ID); ok {
		t.Errorf("la SolicitudViaje id=%d deberia haber sido borrada", creado.ID)
	}
}

func TestSQLite_BuscarSolicitudViajeInexistente(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// El error de GORM (registro no encontrado) se traduce a comma-ok = false.
	if _, ok := alm.BuscarSolicitudViajePorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestSQLite_ViajeInmediatoCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// Crear: GORM debe asignar el ID autogenerado.
	creado := alm.CrearViajeInmediato(models.ViajeInmediato{
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "Universidad",
		HoraSalida:  "07:30",
		Cupos:       4,
		Estado:      "Disponible",
	})

	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado por la base, obtuve 0")
	}

	// Buscar el recién creado.
	encontrado, ok := alm.BuscarViajeInmediatoPorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontro el ViajeInmediato id=%d", creado.ID)
	}

	if encontrado.Origen != "Los Esteros" {
		t.Errorf("origen = %q; esperaba %q", encontrado.Origen, "Los Esteros")
	}

	// Actualizar.
	if _, ok := alm.ActualizarViajeInmediato(creado.ID, models.ViajeInmediato{
		ConductorID: 2,
		Origen:      "Tarqui",
		Destino:     "Terminal",
		HoraSalida:  "08:00",
		Cupos:       5,
		Estado:      "En ruta",
	}); !ok {
		t.Fatalf("no se pudo actualizar el ViajeInmediato id=%d", creado.ID)
	}

	// Borrar y confirmar que ya no está.
	if !alm.BorrarViajeInmediato(creado.ID) {
		t.Errorf("esperaba poder borrar el ViajeInmediato id=%d", creado.ID)
	}

	if _, ok := alm.BuscarViajeInmediatoPorID(creado.ID); ok {
		t.Errorf("el ViajeInmediato id=%d deberia haber sido borrado", creado.ID)
	}
}

func TestSQLite_BuscarViajeInmediatoInexistente(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// El error de GORM (registro no encontrado) se traduce a comma-ok = false.
	if _, ok := alm.BuscarViajeInmediatoPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}
