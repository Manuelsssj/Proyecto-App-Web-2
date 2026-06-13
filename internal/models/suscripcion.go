package models

type Suscripcion struct {
	ID          int    `json:"id"`
	UsuarioID   int    `json:"usuario_id"`
	RutaID      int    `json:"ruta_id"`
	FechaInicio string `json:"fecha_inicio"`
	Estado      string `json:"estado"`
}
