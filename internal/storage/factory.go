package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite"
	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// MODELOS
	modelsRP "RideUleam/internal/models/rutaProgramada"
	modelsUV "RideUleam/internal/models/usuarioVehiculo"
	modelsVI "RideUleam/internal/models/viajeInmediato"

	// STORAGE DE CADA MÓDULO
	storageRP "RideUleam/internal/storage/rutaProgramada"
	storageUV "RideUleam/internal/storage/usuarioVehiculo"
	storageVI "RideUleam/internal/storage/viajeInmediato"
)

type Recursos struct {
	// UsuarioVehiculo
	AlmacenUV storageUV.Almacen
	Usuarios  storageUV.UsuarioRepository

	// ViajeInmediato
	AlmacenVI storageVI.Almacen

	// RutaProgramada
	AlmacenRP storageRP.Almacen

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
		&modelsRP.RutaProgramada{},
		&modelsRP.HorarioRuta{},
		&modelsRP.MantenimientoVehiculo{},
	); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}
	if err := sembrarUsuarios(gdb); err != nil {
		return nil, fmt.Errorf("sembrar usuarios: %w", err)
	}
	//-------------------------------------------------
	// GORM
	//-------------------------------------------------

	almacenUVGorm := storageUV.NuevoAlmacenSQLite(gdb)
	almacenUVGorm.SembrarSiVacio()

	almacenVIGorm := storageVI.NuevoAlmacenSQLite(gdb)
	almacenVIGorm.SembrarSiVacio()

	almacenRPGorm := storageRP.NuevoAlmacenSQLite(gdb)
	almacenRPGorm.SembrarSiVacio()

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
		AlmacenRP:    almacenRPGorm,
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

// //////////////
func sembrarUsuarios(db *gorm.DB) error {
	var cantidad int64

	if err := db.Model(&modelsUV.Usuario{}).Count(&cantidad).Error; err != nil {
		return err
	}

	if cantidad > 0 {
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte("123456"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	usuarios := []modelsUV.Usuario{
		{
			Email:        "conductor1@uleam.com",
			PasswordHash: string(hash),
			Rol:          "conductor",
		},
		{
			Email:        "conductor2@uleam.com",
			PasswordHash: string(hash),
			Rol:          "conductor",
		},
		{
			Email:        "conductor3@uleam.com",
			PasswordHash: string(hash),
			Rol:          "conductor",
		},
		{
			Email:        "conductor4@uleam.com",
			PasswordHash: string(hash),
			Rol:          "conductor",
		},
		{
			Email:        "conductor5@uleam.com",
			PasswordHash: string(hash),
			Rol:          "conductor",
		},
		{
			Email:        "admin@uleam.com",
			PasswordHash: string(hash),
			Rol:          "admin",
		},
	}

	return db.Create(&usuarios).Error
}
