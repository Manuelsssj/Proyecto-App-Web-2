package handlers

import (
	service "cmd/rideUleam/internal/service/usuarioVehiculo"
)

type Server struct {
	Vehiculos *service.VehiculoService
	Auth      *service.AuthService
}

func NewServer(vehiculos *service.VehiculoService, auth *service.AuthService) *Server {
	return &Server{
		Vehiculos: vehiculos,

		Auth: auth,
	}
}
