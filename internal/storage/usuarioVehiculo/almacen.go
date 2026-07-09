package storage

import models "RideUleam/internal/models/usuarioVehiculo"

// Almacen define QUÉ sabe hacer un almacén de la cafetería, sin decir CÓMO.
//
// Memoria (slices) ya cumple esta interfaz sin cambios — por el duck typing
// que vimos en S3 — y AlmacenSQLite (GORM) la cumple igual. El Server depende
// de esta interfaz, no de una implementación concreta: por eso podemos cambiar
// el backend de almacenamiento sin tocar un solo handler.

// Vehículos
type VehiculoRepository interface {
	ListarVehiculos() []models.Vehiculo
	BuscarVehiculoPorID(id int) (models.Vehiculo, bool)
	CrearVehiculo(v models.Vehiculo) models.Vehiculo
	ActualizarVehiculo(id int, datos models.Vehiculo) (models.Vehiculo, bool)
	BorrarVehiculo(id int) bool
}

type Almacen interface {
	VehiculoRepository
}

type UsuarioRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila. Red de seguridad opcional.
var _ Almacen = (*Memoria)(nil)
