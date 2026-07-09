package storage

import (
	"testing"

	models "RideUleam/internal/models/rutaProgramada"
	usuarioModels "RideUleam/internal/models/usuario"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func prepararBDGORMDePrueba(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir SQLite en memoria: %v", err)
	}

	err = db.AutoMigrate(
		&models.RutaProgramada{},
		&models.HorarioRuta{},
		&models.MantenimientoVehiculo{},
		&usuarioModels.Usuario{},
	)
	if err != nil {
		t.Fatalf("no se pudo ejecutar AutoMigrate: %v", err)
	}

	return db
}

func TestAlmacenSQLite_CrearYBuscarRutaProgramada(t *testing.T) {
	db := prepararBDGORMDePrueba(t)
	almacen := NuevoAlmacenSQLite(db)

	ruta := models.RutaProgramada{
		ConductorID: 1,
		Origen:      "Los Esteros",
		Destino:     "ULEAM",
		Costo:       0.75,
	}

	creada := almacen.CrearRutaProgramada(ruta)

	if creada.ID == 0 {
		t.Fatal("se esperaba que la ruta creada tenga ID")
	}

	encontrada, ok := almacen.BuscarRutaProgramadaPorID(creada.ID)
	if !ok {
		t.Fatal("se esperaba encontrar la ruta creada")
	}

	if encontrada.Origen != "Los Esteros" {
		t.Errorf("origen esperado Los Esteros, obtenido %s", encontrada.Origen)
	}

	if encontrada.Destino != "ULEAM" {
		t.Errorf("destino esperado ULEAM, obtenido %s", encontrada.Destino)
	}
}

func TestAlmacenSQLite_CrearYListarRutaProgramada(t *testing.T) {
	db := prepararBDGORMDePrueba(t)
	almacen := NuevoAlmacenSQLite(db)

	almacen.CrearRutaProgramada(models.RutaProgramada{
		ConductorID: 2,
		Origen:      "Tarqui",
		Destino:     "ULEAM",
		Costo:       1.00,
	})

	rutas := almacen.ListarRutasProgramadas()

	if len(rutas) != 1 {
		t.Fatalf("se esperaba 1 ruta, se obtuvieron %d", len(rutas))
	}

	if rutas[0].Origen != "Tarqui" {
		t.Errorf("origen esperado Tarqui, obtenido %s", rutas[0].Origen)
	}
}
