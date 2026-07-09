package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// MODELOS
	modelsUV "RideUleam/internal/models/usuarioVehiculo"
	modelsVI "RideUleam/internal/models/viajeInmediato"

	// STORAGE DE CADA MÓDULO
	storageUV "RideUleam/internal/storage/usuarioVehiculo"
	storageVI "RideUleam/internal/storage/viajeInmediato"
)

type Recursos struct {
	// UsuarioVehiculo
	AlmacenUV storageUV.Almacen
	Usuarios  storageUV.UsuarioRepository

	// ViajeInmediato
	AlmacenVI storageVI.Almacen

	BackendUsado string

	Cerrar func() error
}

func Inicializar(driver, dsn, rutaDB, backend string) (*Recursos, error) {

	gdb, err := abrirGorm(driver, dsn, rutaDB)
	if err != nil {
		return nil, err
	}

	// MIGRAR TODAS LAS TABLAS
	if err := gdb.AutoMigrate(
		&modelsUV.Usuario{},
		&modelsUV.Vehiculo{},
		&modelsVI.ViajeInmediato{},
		&modelsVI.SolicitudViaje{},
		&modelsVI.ParticipanteViaje{},
	); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}

	//-------------------------------------------------
	// GORM
	//-------------------------------------------------

	almacenUVGorm := storageUV.NuevoAlmacenSQLite(gdb)
	almacenUVGorm.SembrarSiVacio()

	almacenVIGorm := storageVI.NuevoAlmacenSQLite(gdb)
	almacenVIGorm.SembrarSiVacio()

	var almacenUV storageUV.Almacen
	var almacenVI storageVI.Almacen

	var sdb *sql.DB

	backendUsado := "gorm"

	if backend == "sqlc" && driver != "postgres" {

		sdb, err = sql.Open("sqlite", rutaDB)
		if err != nil {
			return nil, fmt.Errorf("abrir sql.DB para sqlc: %w", err)
		}

		almacenUV = storageUV.NuevoAlmacenSQLC(sdb)
		almacenVI = storageVI.NuevoAlmacenSQLC(sdb)

		backendUsado = "sqlc"

	} else {

		almacenUV = almacenUVGorm
		almacenVI = almacenVIGorm
	}

	usuarios := storageUV.NuevoUsuarioGORM(gdb)

	cerrar := func() error {

		if sdb != nil {
			if err := sdb.Close(); err != nil {
				return err
			}
		}

		sqlDB, err := gdb.DB()
		if err != nil {
			return err
		}

		return sqlDB.Close()
	}

	return &Recursos{
		AlmacenUV:    almacenUV,
		AlmacenVI:    almacenVI,
		Usuarios:     usuarios,
		BackendUsado: backendUsado,
		Cerrar:       cerrar,
	}, nil
}

func abrirGorm(driver, dsn, rutaDB string) (*gorm.DB, error) {

	switch driver {

	case "postgres":

		var gdb *gorm.DB
		var err error

		for intento := 1; intento <= 10; intento++ {
			gdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				return gdb, nil
			}
			log.Printf("PostgreSQL no esta listo (intento %d/10): %v", intento, err)
			time.Sleep(2 * time.Second)
		}

		return nil, fmt.Errorf("conectar a PostgreSQL tras reintentos: %w", err)
	default: // "sqlite"
		gdb, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("abrir SQLite: %w", err)
		}
		return gdb, nil
	}
}
