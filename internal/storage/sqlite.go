package storage

import (
	"errors"

	"gorm.io/gorm"

	"suscripciones-api/internal/models"
)

// AlmacenSQLite es una implementación de Almacen usando GORM y SQLite.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite crea un almacén SQLite usando una conexión GORM existente.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

func (a *AlmacenSQLite) ListarSuscripciones() ([]models.SuscripcionRuta, error) {
	var lista []models.SuscripcionRuta
	if err := a.db.Order("id").Find(&lista).Error; err != nil {
		return nil, err
	}
	return lista, nil
}

func (a *AlmacenSQLite) ObtenerSuscripcion(id int) (models.SuscripcionRuta, error) {
	var s models.SuscripcionRuta
	if err := a.db.First(&s, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.SuscripcionRuta{}, ErrNoEncontrado
		}
		return models.SuscripcionRuta{}, err
	}
	return s, nil
}

func (a *AlmacenSQLite) CrearSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	if err := a.db.Create(&s).Error; err != nil {
		return models.SuscripcionRuta{}, err
	}
	return s, nil
}

func (a *AlmacenSQLite) ActualizarSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	var actual models.SuscripcionRuta
	if err := a.db.First(&actual, s.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.SuscripcionRuta{}, ErrNoEncontrado
		}
		return models.SuscripcionRuta{}, err
	}

	actual.RutaID = s.RutaID
	actual.UsuarioID = s.UsuarioID

	if err := a.db.Save(&actual).Error; err != nil {
		return models.SuscripcionRuta{}, err
	}
	return actual, nil
}

func (a *AlmacenSQLite) EliminarSuscripcion(id int) error {
	res := a.db.Delete(&models.SuscripcionRuta{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNoEncontrado
	}
	return nil
}

func (a *AlmacenSQLite) ListarPlanes() ([]models.PlanPago, error) {
	var lista []models.PlanPago
	if err := a.db.Order("id").Find(&lista).Error; err != nil {
		return nil, err
	}
	return lista, nil
}

func (a *AlmacenSQLite) ObtenerPlan(id int) (models.PlanPago, error) {
	var p models.PlanPago
	if err := a.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.PlanPago{}, ErrNoEncontrado
		}
		return models.PlanPago{}, err
	}
	return p, nil
}

func (a *AlmacenSQLite) CrearPlan(p models.PlanPago) (models.PlanPago, error) {
	if err := a.db.Create(&p).Error; err != nil {
		return models.PlanPago{}, err
	}
	return p, nil
}

func (a *AlmacenSQLite) ActualizarPlan(p models.PlanPago) (models.PlanPago, error) {
	var actual models.PlanPago
	if err := a.db.First(&actual, p.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.PlanPago{}, ErrNoEncontrado
		}
		return models.PlanPago{}, err
	}

	actual.RutaID = p.RutaID
	actual.ValorSemanal = p.ValorSemanal

	if err := a.db.Save(&actual).Error; err != nil {
		return models.PlanPago{}, err
	}
	return actual, nil
}

func (a *AlmacenSQLite) EliminarPlan(id int) error {
	res := a.db.Delete(&models.PlanPago{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNoEncontrado
	}
	return nil
}

func (a *AlmacenSQLite) ListarHistorial() ([]models.HistorialSuscripcion, error) {
	var lista []models.HistorialSuscripcion
	if err := a.db.Order("id").Find(&lista).Error; err != nil {
		return nil, err
	}
	return lista, nil
}

func (a *AlmacenSQLite) ObtenerHistorial(id int) (models.HistorialSuscripcion, error) {
	var h models.HistorialSuscripcion
	if err := a.db.First(&h, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.HistorialSuscripcion{}, ErrNoEncontrado
		}
		return models.HistorialSuscripcion{}, err
	}
	return h, nil
}

func (a *AlmacenSQLite) CrearHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	if err := a.db.Create(&h).Error; err != nil {
		return models.HistorialSuscripcion{}, err
	}
	return h, nil
}

func (a *AlmacenSQLite) ActualizarHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	var actual models.HistorialSuscripcion
	if err := a.db.First(&actual, h.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.HistorialSuscripcion{}, ErrNoEncontrado
		}
		return models.HistorialSuscripcion{}, err
	}

	actual.SuscripcionID = h.SuscripcionID
	actual.FechaRegistro = h.FechaRegistro
	actual.Estado = h.Estado

	if err := a.db.Save(&actual).Error; err != nil {
		return models.HistorialSuscripcion{}, err
	}
	return actual, nil
}

func (a *AlmacenSQLite) EliminarHistorial(id int) error {
	res := a.db.Delete(&models.HistorialSuscripcion{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNoEncontrado
	}
	return nil
}
