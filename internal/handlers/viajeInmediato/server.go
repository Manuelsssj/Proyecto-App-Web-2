package handlers

import (
	service "RideUleam/internal/service/viajeInmediato"
)

type Server struct {
	ViajeInmediatos    *service.ViajeInmediatoService
	SolicitudViajes    *service.SolicitudViajeService
	ParticipanteViajes *service.ParticipanteViajeService
}

func NewServer(viajeInmediatos *service.ViajeInmediatoService,
	solicitudViajes *service.SolicitudViajeService,
	participanteViajes *service.ParticipanteViajeService) *Server {
	return &Server{
		ViajeInmediatos:    viajeInmediatos,
		SolicitudViajes:    solicitudViajes,
		ParticipanteViajes: participanteViajes,
	}
}
