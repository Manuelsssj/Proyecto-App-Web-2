package models

// SuscripcionRuta representa la suscripción de un usuario a una ruta.
type SuscripcionRuta struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	RutaID    uint `json:"ruta_id"`
	UsuarioID uint `json:"usuario_id"`
}
