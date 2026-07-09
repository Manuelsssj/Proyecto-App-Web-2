package service

import (
	"suscripciones-api/internal/models/suscripciones"
	"suscripciones-api/internal/storage/suscripciones"
)

// PlanService contiene la lógica de negocio de los planes de pago.
type PlanService struct {
	almacen storage.Almacen
}

// NewPlanService crea un nuevo servicio de planes de pago.
func NewPlanService(a storage.Almacen) *PlanService {
	return &PlanService{almacen: a}
}

// Listar devuelve todos los planes de pago.
func (s *PlanService) Listar() ([]models.PlanPago, error) {
	return s.almacen.ListarPlanes()
}

// Obtener devuelve un plan por su ID.
func (s *PlanService) Obtener(id int) (models.PlanPago, error) {
	if id <= 0 {
		return models.PlanPago{}, ErrDatosInvalidos
	}
	return s.almacen.ObtenerPlan(id)
}

// Crear valida y registra un nuevo plan de pago.
func (s *PlanService) Crear(p models.PlanPago) (models.PlanPago, error) {
	if p.RutaID <= 0 || p.ValorSemanal < 0 {
		return models.PlanPago{}, ErrDatosInvalidos
	}
	return s.almacen.CrearPlan(p)
}

// Actualizar valida y modifica un plan existente.
func (s *PlanService) Actualizar(p models.PlanPago) (models.PlanPago, error) {
	if p.ID <= 0 || p.RutaID <= 0 || p.ValorSemanal < 0 {
		return models.PlanPago{}, ErrDatosInvalidos
	}
	return s.almacen.ActualizarPlan(p)
}

// Eliminar borra un plan por su ID.
func (s *PlanService) Eliminar(id int) error {
	if id <= 0 {
		return ErrDatosInvalidos
	}
	return s.almacen.EliminarPlan(id)
}
