package storage

import (
	"time"

	"gorm.io/gorm"

	models "RideUleam/internal/models/suscripciones"
)

// SembrarDatosIniciales inserta datos base si la base está vacía.
func SembrarDatosIniciales(db *gorm.DB) error {
	var totalPlanes int64
	if err := db.Model(&models.PlanPago{}).Count(&totalPlanes).Error; err != nil {
		return err
	}

	if totalPlanes == 0 {
		planes := []models.PlanPago{
			{
				RutaID:       1,
				ValorSemanal: 5.00,
			},
			{
				RutaID:       2,
				ValorSemanal: 7.50,
			},
		}

		if err := db.Create(&planes).Error; err != nil {
			return err
		}
	}

	var totalSuscripciones int64
	if err := db.Model(&models.SuscripcionRuta{}).Count(&totalSuscripciones).Error; err != nil {
		return err
	}

	if totalSuscripciones == 0 {
		suscripcion := models.SuscripcionRuta{
			RutaID:    1,
			UsuarioID: 2,
		}

		if err := db.Create(&suscripcion).Error; err != nil {
			return err
		}

		historial := models.HistorialSuscripcion{
			SuscripcionID: suscripcion.ID,
			FechaRegistro: time.Now().Format("2006-01-02"),
			Estado:        "activa",
		}

		if err := db.Create(&historial).Error; err != nil {
			return err
		}
	}

	return nil
}
