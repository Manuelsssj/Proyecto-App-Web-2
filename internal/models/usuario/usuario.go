package Usuario

import "time"

type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Rol          string    `json:"rol" gorm:"not null;default:conductor"`
	CreadoEn     time.Time `json:"creado_en"`
}
