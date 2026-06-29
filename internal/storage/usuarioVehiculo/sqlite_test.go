package storage

import (
	models "RideUleam/internal/models/usuarioVehiculo"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestSQLite_CrearYBuscarVehiculo(t *testing.T) {
	// Base de datos SQLite en memoria
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir sqlite en memoria: %v", err)
	}

	// Crear la tabla
	if err := db.AutoMigrate(&models.Vehiculo{}); err != nil {
		t.Fatalf("falló AutoMigrate: %v", err)
	}

	// Repositorio real
	repo := NuevoAlmacenSQLite(db)

	// Crear un vehículo
	creado := repo.CrearVehiculo(models.Vehiculo{
		ConductorID: 1,
		Placa:       "MBA-1234",
		Marca:       "Toyota",
		Modelo:      "Hiace",
		Capacidad:   30,
	})

	if creado.ID == 0 {
		t.Fatalf("esperaba que GORM asignara un ID")
	}

	// Buscar por ID
	encontrado, ok := repo.BuscarVehiculoPorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontró el vehículo recién creado")
	}

	if encontrado.Placa != "MBA-1234" {
		t.Errorf("placa = %q; esperaba %q", encontrado.Placa, "MBA-1234")
	}

	// Verificar que aparece al listar
	lista := repo.ListarVehiculos()

	if len(lista) != 1 {
		t.Fatalf("esperaba 1 vehículo, obtuvo %d", len(lista))
	}
}
