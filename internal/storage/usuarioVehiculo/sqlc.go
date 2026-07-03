package storage

import (
	"context"
	"database/sql"

	models "RideUleam/internal/models/usuarioVehiculo"
	"RideUleam/internal/storage/sqlcdb"
)

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

// --- mapeo sqlc -> dominio (la capa que traduce) ---
func aVehiculoDominio(v sqlcdb.Vehiculo) models.Vehiculo {
	return models.Vehiculo{
		ID:          int(v.ID),
		ConductorID: int(v.ConductorID),
		Placa:       v.Placa,
		Marca:       v.Marca,
		Modelo:      v.Modelo,
		Capacidad:   int(v.Capacidad),
	}
}

func (a *AlmacenSQLC) ListarVehiculos() []models.Vehiculo {
	filas, err := a.q.ListarVehiculos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.Vehiculo, 0, len(filas))
	for _, f := range filas {
		out = append(out, aVehiculoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarVehiculoPorID(id int) (models.Vehiculo, bool) {
	f, err := a.q.BuscarVehiculoPorID(context.Background(), int64(id))
	if err != nil {
		return models.Vehiculo{}, false // absorbe sql.ErrNoRows
	}
	return aVehiculoDominio(f), true
}

func (a *AlmacenSQLC) CrearVehiculo(v models.Vehiculo) models.Vehiculo {
	f, err := a.q.CrearVehiculo(context.Background(), sqlcdb.CrearVehiculoParams{
		ConductorID: int64(v.ConductorID),
		Placa:       v.Placa,
		Marca:       v.Marca,
		Modelo:      v.Modelo,
		Capacidad:   int64(v.Capacidad),
	})
	if err != nil {
		return models.Vehiculo{}
	}
	return aVehiculoDominio(f)
}

func (a *AlmacenSQLC) ActualizarVehiculo(id int, datos models.Vehiculo) (models.Vehiculo, bool) {
	f, err := a.q.ActualizarVehiculo(context.Background(), sqlcdb.ActualizarVehiculoParams{

		ConductorID: int64(datos.ConductorID),
		Placa:       datos.Placa,
		Marca:       datos.Marca,
		Modelo:      datos.Modelo,
		Capacidad:   int64(datos.Capacidad),
		ID:          int64(id),
	})
	if err != nil {
		return models.Vehiculo{}, false
	}
	return aVehiculoDominio(f), true
}

func (a *AlmacenSQLC) BorrarVehiculo(id int) bool {
	filas, err := a.q.BorrarVehiculo(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

var _ Almacen = (*AlmacenSQLC)(nil)
