package storage

import models "RideUleam/internal/models/suscripciones"

// Almacen define las operaciones de persistencia que necesita la aplicación.
type Almacen interface {
	// SuscripcionRuta
	ListarSuscripciones() ([]models.SuscripcionRuta, error)
	ObtenerSuscripcion(id int) (models.SuscripcionRuta, error)
	CrearSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error)
	ActualizarSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error)
	EliminarSuscripcion(id int) error

	// PlanPago
	ListarPlanes() ([]models.PlanPago, error)
	ObtenerPlan(id int) (models.PlanPago, error)
	CrearPlan(p models.PlanPago) (models.PlanPago, error)
	ActualizarPlan(p models.PlanPago) (models.PlanPago, error)
	EliminarPlan(id int) error

	// HistorialSuscripcion
	ListarHistorial() ([]models.HistorialSuscripcion, error)
	ObtenerHistorial(id int) (models.HistorialSuscripcion, error)
	CrearHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error)
	ActualizarHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error)
	EliminarHistorial(id int) error
}

// UserRepository define las operaciones necesarias para autenticación.
type UserRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

var _ Almacen = (*AlmacenMemoria)(nil)
var _ Almacen = (*AlmacenGORM)(nil)
