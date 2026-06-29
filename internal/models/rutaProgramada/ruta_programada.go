package models

type RutaProgramada struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	ConductorID int     `json:"conductor_id" gorm:"not null"`
	Origen      string  `json:"origen" gorm:"not null"`
	Destino     string  `json:"destino" gorm:"not null"`
	Costo       float64 `json:"costo" gorm:"not null"`
}

func (RutaProgramada) TableName() string {
	return "rutas_programadas"
}

type HorarioRuta struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	RutaID int    `json:"ruta_id" gorm:"not null"`
	Dia    string `json:"dia" gorm:"not null"`
	Hora   string `json:"hora" gorm:"not null"`
}

func (HorarioRuta) TableName() string {
	return "horarios_ruta"
}

type MantenimientoVehiculo struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	VehiculoID  int    `json:"vehiculo_id" gorm:"not null"`
	FechaInicio string `json:"fecha_inicio" gorm:"not null"`
	FechaFin    string `json:"fecha_fin" gorm:"not null"`
	Motivo      string `json:"motivo" gorm:"not null"`
}

func (MantenimientoVehiculo) TableName() string {
	return "mantenimientos_vehiculo"
}
