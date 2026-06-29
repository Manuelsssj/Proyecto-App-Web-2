package service

import (
	"suscripciones-api/internal/models"
	"suscripciones-api/internal/storage"
)

// SuscripcionService contiene la lógica de negocio de las suscripciones a rutas.
type SuscripcionService struct {
	almacen storage.Almacen
}

// NewSuscripcionService crea un nuevo servicio de suscripciones.
func NewSuscripcionService(a storage.Almacen) *SuscripcionService {
	return &SuscripcionService{almacen: a}
}

// Listar devuelve todas las suscripciones.
func (s *SuscripcionService) Listar() ([]models.SuscripcionRuta, error) {
	return s.almacen.ListarSuscripciones()
}

// Obtener devuelve una suscripción por su ID.
func (s *SuscripcionService) Obtener(id int) (models.SuscripcionRuta, error) {
	if id <= 0 {
		return models.SuscripcionRuta{}, ErrDatosInvalidos
	}
	return s.almacen.ObtenerSuscripcion(id)
}

// Crear valida y registra una nueva suscripción.
func (s *SuscripcionService) Crear(sub models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	if sub.RutaID <= 0 || sub.UsuarioID <= 0 {
		return models.SuscripcionRuta{}, ErrDatosInvalidos
	}
	return s.almacen.CrearSuscripcion(sub)
}

//go test ./internal/service -cover

// Actualizar valida y modifica una suscripción existente.
func (s *SuscripcionService) Actualizar(sub models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	if sub.ID <= 0 || sub.RutaID <= 0 || sub.UsuarioID <= 0 {
		return models.SuscripcionRuta{}, ErrDatosInvalidos
	}
	return s.almacen.ActualizarSuscripcion(sub)
}

// Eliminar borra una suscripción por su ID.
func (s *SuscripcionService) Eliminar(id int) error {
	if id <= 0 {
		return ErrDatosInvalidos
	}
	return s.almacen.EliminarSuscripcion(id)
}
