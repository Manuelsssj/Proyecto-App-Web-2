package handlers

import (
	authService "RideUleam/internal/service"
	rutaService "RideUleam/internal/service/rutaProgramada"
	rutaStorage "RideUleam/internal/storage/rutaProgramada"
	usuarioStorage "RideUleam/internal/storage/usuario"
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
