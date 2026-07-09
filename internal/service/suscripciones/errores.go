package service

import (
	"errors"

	"suscripciones-api/internal/storage/suscripciones"
)

// Reutilizamos los errores definidos en la capa de almacenamiento
// para mantener una sola fuente de verdad.
var (
	ErrNoEncontrado          = storage.ErrNoEncontrado
	ErrDatosInvalidos        = storage.ErrDatosInvalidos
	ErrCredencialesInvalidas = errors.New("credenciales inválidas")
	ErrEmailEnUso            = errors.New("el email ya está en uso")
)
