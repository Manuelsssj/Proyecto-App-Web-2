package models

import usuarioModels "RideUleam/internal/models/usuarioVehiculo"

type ViajeInmediato struct {
	ID            int                   `json:"id" gorm:"primaryKey"`
	ConductorID   int                   `json:"conductor_id" gorm:"not null"`
	Conductor     usuarioModels.Usuario `json:"conductor,omitempty" gorm:"foreignKey:ConductorID"`
	Origen        string                `json:"origen" gorm:"not null"`
	Destino       string                `json:"destino"`
	HoraSalida    string                `json:"hora_salida"`
	Cupos         int                   `json:"cupos"`
	Estado        string                `json:"estado"`
	Solicitudes   []SolicitudViaje      `json:"solicitudes,omitempty" gorm:"foreignKey:ViajeID"`
	Participantes []ParticipanteViaje   `json:"participantes,omitempty" gorm:"foreignKey:ViajeID"`
}
