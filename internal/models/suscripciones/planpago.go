package models

// PlanPago representa el plan de pago asociado a una ruta.
type PlanPago struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	RutaID       uint    `json:"ruta_id"`
	ValorSemanal float64 `json:"valor_semanal"`
}
