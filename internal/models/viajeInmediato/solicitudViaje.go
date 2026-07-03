package models

type SolicitudViaje struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	ViajeID    int    `json:"viaje_id" gorm:"not null"`
	PasajeroID int    `json:"pasajero_id" `
	Estado     string `json:"estado" `
}
