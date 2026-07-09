package models

import usuarioModels "RideUleam/internal/models/usuarioVehiculo"

type SolicitudViaje struct {
	ID         int                   `json:"id" gorm:"primaryKey"`
	ViajeID    int                   `json:"viaje_id" gorm:"not null"`
	Viaje      ViajeInmediato        `json:"viaje,omitempty" gorm:"foreignKey:ViajeID"`
	PasajeroID int                   `json:"pasajero_id" gorm:"not null"`
	Pasajero   usuarioModels.Usuario `json:"pasajero,omitempty" gorm:"foreignKey:PasajeroID"`
	Estado     string                `json:"estado"`
}
