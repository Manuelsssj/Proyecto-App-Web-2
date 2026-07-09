package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Puerto     string
	DBDriver   string
	DBDsn      string
	RutaDB     string
	JWTSecreto string
}

func Cargar() Config {
	_ = godotenv.Load()

	return Config{
		Puerto:     getEnv("PUERTO", ":8080"),
		DBDriver:   getEnv("DB_DRIVER", "sqlite"),
		DBDsn:      getEnv("DB_DSN", ""),
		RutaDB:     getEnv("RUTA_DB", "rideuleam.db"),
		JWTSecreto: getEnv("JWT_SECRETO", "rideuleam-secreto-dev"),
	}
}

func getEnv(clave, defecto string) string {
	if valor := os.Getenv(clave); valor != "" {
		return valor
	}
	return defecto
}
