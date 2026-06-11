package models

type Estado struct {
	ID          int    `json:"id"`
	RutaID      int    `json:"ruta_id"`
	Estado      string `json:"estado"`
	Motivo      string `json:"motivo"`
	FechaInicio string `json:"fecha_inicio"`
	FechaFin    string `json:"fecha_fin"`
}
