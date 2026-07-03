package handlers

import (
	authService "cmd/rideUleam/internal/service"
	rutaService "cmd/rideUleam/internal/service/rutaProgramada"
	rutaStorage "cmd/rideUleam/internal/storage/rutaProgramada"
	usuarioStorage "cmd/rideUleam/internal/storage/usuario"
)

type Server struct {
	RutaProgramadaService *rutaService.RutaProgramadaService
	Auth                  *authService.AuthService
}

func NewServer(almacen rutaStorage.Almacen, userRepo usuarioStorage.UserRepository) *Server {
	return &Server{
		RutaProgramadaService: rutaService.NewRutaProgramadaService(almacen),
		Auth:                  authService.NewAuthService(userRepo),
	}
}
