package models

import usuarioModels "RideUleam/internal/models/usuarioVehiculo"

type ParticipanteViaje struct {
	ID        int                   `json:"id" gorm:"primaryKey"`
	ViajeID   int                   `json:"viaje_id" gorm:"not null"`
	Viaje     ViajeInmediato        `json:"viaje,omitempty" gorm:"foreignKey:ViajeID"`
	UsuarioID int                   `json:"usuario_id" gorm:"not null"`
	Usuario   usuarioModels.Usuario `json:"usuario,omitempty" gorm:"foreignKey:UsuarioID"`
}
