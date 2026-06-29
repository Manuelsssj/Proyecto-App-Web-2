package models

type ParticipanteViaje struct {
	ID        int `json:"id" gorm:"primaryKey"`
	ViajeID   int `json:"viaje_id" gorm:"not null"`
	UsuarioID int `json:"usuario_id" gorm:"not null"`
}
