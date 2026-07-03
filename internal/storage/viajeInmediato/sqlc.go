package storage

import (
	"context"
	"database/sql"

	models "RideUleam/internal/models/viajeInmediato"
	"RideUleam/internal/storage/sqlcdb"
)

type AlmacenSQLC struct {
	q *sqlcdb.Queries
}

func NuevoAlmacenSQLC(db *sql.DB) *AlmacenSQLC {
	return &AlmacenSQLC{q: sqlcdb.New(db)}
}

// --- mapeo sqlc -> dominio (la capa que traduce) ---
func aViajeInmediatoDominio(vi sqlcdb.ViajeInmediato) models.ViajeInmediato {
	return models.ViajeInmediato{
		ID:          int(vi.ID),
		ConductorID: int(vi.ConductorID),
		Origen:      vi.Origen,
		Destino:     vi.Destino,
		HoraSalida:  vi.HoraSalida,
		Cupos:       int(vi.Cupos),
		Estado:      vi.Estado,
	}
}

func (a *AlmacenSQLC) ListarViajeInmediatos() []models.ViajeInmediato {
	filas, err := a.q.ListarViajesInmediatos(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.ViajeInmediato, 0, len(filas))
	for _, f := range filas {
		out = append(out, aViajeInmediatoDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarViajeInmediatoPorID(id int) (models.ViajeInmediato, bool) {
	f, err := a.q.BuscarViajeInmediatoPorID(context.Background(), int64(id))
	if err != nil {
		return models.ViajeInmediato{}, false // absorbe sql.ErrNoRows
	}
	return aViajeInmediatoDominio(f), true
}

func (a *AlmacenSQLC) CrearViajeInmediato(vi models.ViajeInmediato) models.ViajeInmediato {
	f, err := a.q.CrearViajeInmediato(context.Background(), sqlcdb.CrearViajeInmediatoParams{
		ConductorID: int64(vi.ConductorID),
		Origen:      vi.Origen,
		Destino:     vi.Destino,
		HoraSalida:  vi.HoraSalida,
		Cupos:       int64(vi.Cupos),
		Estado:      vi.Estado,
	})
	if err != nil {
		return models.ViajeInmediato{}
	}
	return aViajeInmediatoDominio(f)
}

func (a *AlmacenSQLC) ActualizarViajeInmediato(id int, datos models.ViajeInmediato) (models.ViajeInmediato, bool) {
	f, err := a.q.ActualizarViajeInmediato(context.Background(), sqlcdb.ActualizarViajeInmediatoParams{

		ConductorID: int64(datos.ConductorID),
		Origen:      datos.Origen,
		Destino:     datos.Destino,
		HoraSalida:  datos.HoraSalida,
		Cupos:       int64(datos.Cupos),
		Estado:      datos.Estado,
		ID:          int64(id),
	})
	if err != nil {
		return models.ViajeInmediato{}, false
	}
	return aViajeInmediatoDominio(f), true
}

func (a *AlmacenSQLC) BorrarViajeInmediato(id int) bool {
	filas, err := a.q.BorrarViajeInmediato(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// --- mapeo sqlc -> dominio (la capa que traduce) ---
func aSolicitudViajeDominio(sv sqlcdb.SolicitudViaje) models.SolicitudViaje {
	return models.SolicitudViaje{
		ID:         int(sv.ID),
		ViajeID:    int(sv.ViajeID),
		PasajeroID: int(sv.PasajeroID),
		Estado:     sv.Estado,
	}
}

func (a *AlmacenSQLC) ListarSolicitudViajes() []models.SolicitudViaje {
	filas, err := a.q.ListarSolicitudesViajes(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.SolicitudViaje, 0, len(filas))
	for _, f := range filas {
		out = append(out, aSolicitudViajeDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarSolicitudViajePorID(id int) (models.SolicitudViaje, bool) {
	f, err := a.q.BuscarSolicitudViajePorID(context.Background(), int64(id))
	if err != nil {
		return models.SolicitudViaje{}, false // absorbe sql.ErrNoRows
	}
	return aSolicitudViajeDominio(f), true
}

func (a *AlmacenSQLC) CrearSolicitudViaje(sv models.SolicitudViaje) models.SolicitudViaje {
	f, err := a.q.CrearSolicitudViaje(context.Background(), sqlcdb.CrearSolicitudViajeParams{
		ViajeID:    int64(sv.ViajeID),
		PasajeroID: int64(sv.PasajeroID),
		Estado:     sv.Estado,
	})
	if err != nil {
		return models.SolicitudViaje{}
	}
	return aSolicitudViajeDominio(f)
}

func (a *AlmacenSQLC) ActualizarSolicitudViaje(id int, datos models.SolicitudViaje) (models.SolicitudViaje, bool) {
	f, err := a.q.ActualizarSolicitudViaje(context.Background(), sqlcdb.ActualizarSolicitudViajeParams{

		ViajeID:    int64(datos.ViajeID),
		PasajeroID: int64(datos.PasajeroID),
		Estado:     datos.Estado,
		ID:         int64(id),
	})
	if err != nil {
		return models.SolicitudViaje{}, false
	}
	return aSolicitudViajeDominio(f), true
}

func (a *AlmacenSQLC) BorrarSolicitudViaje(id int) bool {
	filas, err := a.q.BorrarSolicitudViaje(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

// --- mapeo sqlc -> dominio (la capa que traduce) ---
func aParticipanteViajeDominio(pv sqlcdb.ParticipanteViaje) models.ParticipanteViaje {
	return models.ParticipanteViaje{
		ID:        int(pv.ID),
		ViajeID:   int(pv.ViajeID),
		UsuarioID: int(pv.UsuarioID),
	}
}

func (a *AlmacenSQLC) ListarParticipanteViajes() []models.ParticipanteViaje {
	filas, err := a.q.ListarParticipantesViajes(context.Background())
	if err != nil {
		return nil
	}
	out := make([]models.ParticipanteViaje, 0, len(filas))
	for _, f := range filas {
		out = append(out, aParticipanteViajeDominio(f))
	}
	return out
}

func (a *AlmacenSQLC) BuscarParticipanteViajePorID(id int) (models.ParticipanteViaje, bool) {
	f, err := a.q.BuscarParticipanteViajePorID(context.Background(), int64(id))
	if err != nil {
		return models.ParticipanteViaje{}, false // absorbe sql.ErrNoRows
	}
	return aParticipanteViajeDominio(f), true
}

func (a *AlmacenSQLC) CrearParticipanteViaje(pv models.ParticipanteViaje) models.ParticipanteViaje {
	f, err := a.q.CrearParticipanteViaje(context.Background(), sqlcdb.CrearParticipanteViajeParams{
		ViajeID:   int64(pv.ViajeID),
		UsuarioID: int64(pv.UsuarioID),
	})
	if err != nil {
		return models.ParticipanteViaje{}
	}
	return aParticipanteViajeDominio(f)
}

func (a *AlmacenSQLC) ActualizarParticipanteViaje(id int, datos models.ParticipanteViaje) (models.ParticipanteViaje, bool) {
	f, err := a.q.ActualizarParticipanteViaje(context.Background(), sqlcdb.ActualizarParticipanteViajeParams{

		ViajeID:   int64(datos.ViajeID),
		UsuarioID: int64(datos.UsuarioID),

		ID: int64(id),
	})
	if err != nil {
		return models.ParticipanteViaje{}, false
	}
	return aParticipanteViajeDominio(f), true
}

func (a *AlmacenSQLC) BorrarParticipanteViaje(id int) bool {
	filas, err := a.q.BorrarParticipanteViaje(context.Background(), int64(id))
	if err != nil {
		return false
	}
	return filas > 0
}

var _ Almacen = (*AlmacenSQLC)(nil)
