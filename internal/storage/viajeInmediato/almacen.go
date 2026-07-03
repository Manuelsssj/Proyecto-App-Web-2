package storage

import models "cmd/rideUleam/internal/models/viajeInmediato"

// Almacen define QUÉ sabe hacer un almacén de la cafetería, sin decir CÓMO.
//
// Memoria (slices) ya cumple esta interfaz sin cambios — por el duck typing
// que vimos en S3 — y AlmacenSQLite (GORM) la cumple igual. El Server depende
// de esta interfaz, no de una implementación concreta: por eso podemos cambiar
// el backend de almacenamiento sin tocar un solo handler.

type ViajeInmediatoRepository interface {

	// ViajeInmediato
	ListarViajeInmediatos() []models.ViajeInmediato
	BuscarViajeInmediatoPorID(id int) (models.ViajeInmediato, bool)
	CrearViajeInmediato(vi models.ViajeInmediato) models.ViajeInmediato
	ActualizarViajeInmediato(id int, datos models.ViajeInmediato) (models.ViajeInmediato, bool)
	BorrarViajeInmediato(id int) bool
}
type SolicitudViajeRepository interface {

	// SolicitudViaje
	ListarSolicitudViajes() []models.SolicitudViaje
	BuscarSolicitudViajePorID(id int) (models.SolicitudViaje, bool)
	CrearSolicitudViaje(sv models.SolicitudViaje) models.SolicitudViaje
	ActualizarSolicitudViaje(id int, datos models.SolicitudViaje) (models.SolicitudViaje, bool)
	BorrarSolicitudViaje(id int) bool
}
type ParticipanteViajeRepository interface {

	//ParticipanteViaje
	ListarParticipanteViajes() []models.ParticipanteViaje
	BuscarParticipanteViajePorID(id int) (models.ParticipanteViaje, bool)
	CrearParticipanteViaje(pv models.ParticipanteViaje) models.ParticipanteViaje
	ActualizarParticipanteViaje(id int, datos models.ParticipanteViaje) (models.ParticipanteViaje, bool)
	BorrarParticipanteViaje(id int) bool
}

type Almacen interface {
	ParticipanteViajeRepository
	SolicitudViajeRepository
	ViajeInmediatoRepository
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila. Red de seguridad opcional.
var _ Almacen = (*Memoria)(nil)
