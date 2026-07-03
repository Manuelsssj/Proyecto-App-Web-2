package service

import "errors"

//Lista de errores posibles
var (
	ErrViajeIDInvalido               = errors.New("viaje_id es requerido")
	ErrUsuarioIDInvalido             = errors.New("usuario_id es requerido")
	ErrParticipanteViajeNoEncontrado = errors.New("participanteViaje no encontrado")
	ErrSolicitudViajeNoEncontrado    = errors.New("solicitudViaje no encontrado")
	ErrPasajeroIDInvalido            = errors.New("pasajero_id es requerido")
	ErrViajeInmediatoNoEncontrado    = errors.New("viajeInmediato no encontrado")
	ErrConductorIDInvalido           = errors.New("conductor_id es requerido")
)
