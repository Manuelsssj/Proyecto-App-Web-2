package service

import "errors"

// Lista de errores posibles
var (
	ErrVehiculoNoEncontrado = errors.New("Vehiculo no encontrado")
	ErrConductorIDInvalido  = errors.New("conductor_id es requerido")

	ErrEmailEnUso            = errors.New("email ya en uso")
	ErrCredencialesInvalidas = errors.New("email o contraseña")
)
