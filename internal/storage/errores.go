package storage

import "errors"

// Errores comunes de la capa de almacenamiento.
var (
	ErrNoEncontrado   = errors.New("recurso no encontrado")
	ErrDatosInvalidos = errors.New("datos inválidos")
)
