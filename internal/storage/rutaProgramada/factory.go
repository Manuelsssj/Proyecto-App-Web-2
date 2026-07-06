package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AbrirGorm(driver, dsn, rutaDB string) (*gorm.DB, error) {
	switch driver {
	case "postgres":
		var db *gorm.DB
		var err error

		for intento := 1; intento <= 10; intento++ {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				return db, nil
			}

			log.Printf("PostgreSQL no está listo, intento %d/10: %v", intento, err)
			time.Sleep(2 * time.Second)
		}

		return nil, fmt.Errorf("no se pudo conectar a PostgreSQL: %w", err)

	default:
		db, err := gorm.Open(sqlite.Open(rutaDB), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("no se pudo abrir SQLite: %w", err)
		}

		return db, nil
	}
}
