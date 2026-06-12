// update
package models

type Ruta struct {
	ID         int    `json:"id"`
	Sector     string `json:"sector"`
	Destino    string `json:"destino"`
	HoraSalida string `json:"hora_salida"`
	Cupos      int    `json:"cupos"`
	Estado     string `json:"estado"`
	Conductor  string `json:"conductor"`
}
