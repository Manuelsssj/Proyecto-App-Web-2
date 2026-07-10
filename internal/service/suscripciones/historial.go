package service

import (
	"time"

	models "suscripciones-api/internal/models/suscripciones"
	storage "suscripciones-api/internal/storage/suscripciones"
)

// HistorialService contiene la lógica de negocio del historial de suscripciones.
type HistorialService struct {
	almacen storage.Almacen
}

// NewHistorialService crea un nuevo servicio de historial.
func NewHistorialService(a storage.Almacen) *HistorialService {
	return &HistorialService{almacen: a}
}

// Listar devuelve todo el historial.
func (s *HistorialService) Listar() ([]models.HistorialSuscripcion, error) {
	return s.almacen.ListarHistorial()
}

// Obtener devuelve un registro de historial por su ID.
func (s *HistorialService) Obtener(id int) (models.HistorialSuscripcion, error) {
	if id <= 0 {
		return models.HistorialSuscripcion{}, ErrDatosInvalidos
	}
	return s.almacen.ObtenerHistorial(id)
}

// Crear valida y registra un nuevo historial. Si no se envía fecha, usa la actual.
func (s *HistorialService) Crear(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	if h.SuscripcionID <= 0 || h.Estado == "" {
		return models.HistorialSuscripcion{}, ErrDatosInvalidos
	}
	if h.FechaRegistro == "" {
		h.FechaRegistro = time.Now().Format("2006-01-02")
	}
	return s.almacen.CrearHistorial(h)
}

// Actualizar valida y modifica un registro de historial existente.
func (s *HistorialService) Actualizar(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	if h.ID <= 0 || h.SuscripcionID <= 0 || h.Estado == "" {
		return models.HistorialSuscripcion{}, ErrDatosInvalidos
	}
	return s.almacen.ActualizarHistorial(h)
}

// Eliminar borra un registro de historial por su ID.
func (s *HistorialService) Eliminar(id int) error {
	if id <= 0 {
		return ErrDatosInvalidos
	}
	return s.almacen.EliminarHistorial(id)
}
