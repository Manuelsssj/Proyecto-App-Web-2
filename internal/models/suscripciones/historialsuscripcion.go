package models

// HistorialSuscripcion representa un registro histórico del estado de una suscripción.
type HistorialSuscripcion struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	SuscripcionID uint   `json:"suscripcion_id"`
	FechaRegistro string `json:"fecha_registro"`
	Estado        string `json:"estado"`
}
