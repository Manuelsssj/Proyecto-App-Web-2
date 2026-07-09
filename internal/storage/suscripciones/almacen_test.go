package storage

import (
	models "suscripciones-api/internal/models/suscripciones"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestAlmacenGORM_CrearYListarSuscripcion(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir sqlite en memoria: %v", err)
	}

	err = db.AutoMigrate(&models.SuscripcionRuta{})
	if err != nil {
		t.Fatalf("falló AutoMigrate: %v", err)
	}

	almacen := NuevoAlmacenGORM(db)

	creada, err := almacen.CrearSuscripcion(models.SuscripcionRuta{
		RutaID:    1,
		UsuarioID: 1,
	})
	if err != nil {
		t.Fatalf("no se pudo crear la suscripción: %v", err)
	}

	if creada.ID == 0 {
		t.Fatal("se esperaba que GORM asigne un ID")
	}

	lista, err := almacen.ListarSuscripciones()
	if err != nil {
		t.Fatalf("no se pudo listar suscripciones: %v", err)
	}

	if len(lista) != 1 {
		t.Fatalf("se esperaba 1 suscripción, se obtuvo %d", len(lista))
	}

	if lista[0].RutaID != 1 || lista[0].UsuarioID != 1 {
		t.Fatalf("la suscripción listada no coincide: %+v", lista[0])
	}
}

//go test ./internal/storage -v
