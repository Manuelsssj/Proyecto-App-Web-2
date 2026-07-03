package models

type ViajeInmediato struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ConductorID int    `json:"conductor_id" gorm:"not null"`
	Origen      string `json:"origen" gorm:"not null"`
	Destino     string `json:"destino"`
	HoraSalida  string `json:"hora_salida" `
	Cupos       int    `json:"cupos" `
	Estado      string `json:"estado" `
}
