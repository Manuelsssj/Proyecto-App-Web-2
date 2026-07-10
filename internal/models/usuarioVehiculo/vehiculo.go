package models

type Vehiculo struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	ConductorID int     `json:"conductor_id" gorm:"not null"`
	Conductor   Usuario `json:"conductor,omitempty" gorm:"foreignKey:ConductorID"`
	Placa       string  `json:"placa" gorm:"not null"`
	Marca       string  `json:"marca"`
	Modelo      string  `json:"modelo"`
	Capacidad   int     `json:"capacidad" gorm:"not null"`
}
