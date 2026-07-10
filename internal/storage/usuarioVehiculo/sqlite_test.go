package storage

import (
	models "RideUleam/internal/models/usuarioVehiculo"
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
	if err := gdb.AutoMigrate(&models.Vehiculo{}, &models.Usuario{}); err != nil {
		t.Fatalf("fallo AutoMigrate: %v", err)
	}
	return gdb
}

func TestSQLite_VehiculoCRUD(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))

	// Crear: GORM debe asignar el ID autogenerado.
	creado := alm.CrearVehiculo(models.Vehiculo{
		ConductorID: 1,
		Placa:       "MBT-456",
		Marca:       "Chevrolet",
		Modelo:      "Spark",
		Capacidad:   4,
	})
	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado por la base, obtuve 0")
	}

	// Buscar el recien creado.
	encontrado, ok := alm.BuscarVehiculoPorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontro el Vehiculo id=%d", creado.ID)
	}
	if encontrado.Placa != "MBT-456" {
		t.Errorf("placa = %q; esperaba %q", encontrado.Placa, "MBT-456")
	}

	// Actualizar.
	if _, ok := alm.ActualizarVehiculo(creado.ID, models.Vehiculo{
		ConductorID: 2,
		Placa:       "ABC-123",
		Marca:       "Toyota",
		Modelo:      "Corolla",
		Capacidad:   5,
	}); !ok {
		t.Fatalf("no se pudo actualizar el Vehiculo id=%d", creado.ID)
	}

	// Borrar y confirmar que ya no esta.
	if !alm.BorrarVehiculo(creado.ID) {
		t.Errorf("esperaba poder borrar el Vehiculo id=%d", creado.ID)
	}
	if _, ok := alm.BuscarVehiculoPorID(creado.ID); ok {
		t.Errorf("el Vehiculo id=%d deberia haber sido borrado", creado.ID)
	}
}

func TestSQLite_BuscarInexistente(t *testing.T) {
	alm := NuevoAlmacenSQLite(nuevaDBPrueba(t))
	// El error de GORM (registro no encontrado) se traduce a comma-ok = false.
	if _, ok := alm.BuscarVehiculoPorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

// TestSQLite_UsuarioEmailUnico prueba una garantia que SOLO la base puede dar:
// el indice unico de email impide dos usuarios con el mismo correo.
func TestSQLite_UsuarioEmailUnico(t *testing.T) {
	repo := NuevoUsuarioGORM(nuevaDBPrueba(t))

	if _, err := repo.CrearUsuario(models.Usuario{Email: "ana@uleam.edu.ec", PasswordHash: "hash1"}); err != nil {
		t.Fatalf("el primer usuario deberia crearse sin error: %v", err)
	}
	if _, err := repo.CrearUsuario(models.Usuario{Email: "ana@uleam.edu.ec", PasswordHash: "hash2"}); err == nil {
		t.Errorf("esperaba error por email duplicado (indice unico), no lo hubo")
	}
}
