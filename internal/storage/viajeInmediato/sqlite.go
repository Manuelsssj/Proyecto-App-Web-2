package storage

import (
	models "cmd/rideUleam/internal/models/viajeInmediato"

	"gorm.io/gorm"
)

// AlmacenSQLite implementa la interfaz Almacen usando GORM sobre SQLite.
//
// Fíjense: los métodos tienen EXACTAMENTE las mismas firmas que los de Memoria.
// Por eso el Server y los handlers no se enteran de cuál de los dos reciben.
type AlmacenSQLite struct {
	db *gorm.DB
}

// NuevoAlmacenSQLite envuelve una conexión *gorm.DB ya abierta.
func NuevoAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

// =========================================================
// PRViajeInmediato
// =========================================================

func (a *AlmacenSQLite) ListarViajeInmediatos() []models.ViajeInmediato {
	var viajeInmediatos []models.ViajeInmediato
	a.db.Find(&viajeInmediatos)
	return viajeInmediatos
}

func (a *AlmacenSQLite) BuscarViajeInmediatoPorID(id int) (models.ViajeInmediato, bool) {
	var vi models.ViajeInmediato
	if err := a.db.First(&vi, id).Error; err != nil {
		// Absorbemos el error de la DB y conservamos la firma comma-ok.
		return models.ViajeInmediato{}, false
	}
	return vi, true
}

func (a *AlmacenSQLite) CrearViajeInmediato(vi models.ViajeInmediato) models.ViajeInmediato {
	a.db.Create(&vi) // GORM rellena el ID autogenerado en &p
	return vi
}

func (a *AlmacenSQLite) ActualizarViajeInmediato(id int, datos models.ViajeInmediato) (models.ViajeInmediato, bool) {
	var existente models.ViajeInmediato
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.ViajeInmediato{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarViajeInmediato(id int) bool {
	res := a.db.Delete(&models.ViajeInmediato{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SolicitudViaje
// =========================================================

func (a *AlmacenSQLite) ListarSolicitudViajes() []models.SolicitudViaje {
	var solicitudViajes []models.SolicitudViaje
	a.db.Find(&solicitudViajes)
	return solicitudViajes
}

func (a *AlmacenSQLite) BuscarSolicitudViajePorID(id int) (models.SolicitudViaje, bool) {
	var sv models.SolicitudViaje
	if err := a.db.First(&sv, id).Error; err != nil {
		return models.SolicitudViaje{}, false
	}
	return sv, true
}

func (a *AlmacenSQLite) CrearSolicitudViaje(sv models.SolicitudViaje) models.SolicitudViaje {
	a.db.Create(&sv)
	return sv
}

func (a *AlmacenSQLite) ActualizarSolicitudViaje(id int, datos models.SolicitudViaje) (models.SolicitudViaje, bool) {
	var existente models.SolicitudViaje
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.SolicitudViaje{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarSolicitudViaje(id int) bool {
	res := a.db.Delete(&models.SolicitudViaje{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// ParticipanteViaje
// =========================================================

func (a *AlmacenSQLite) ListarParticipanteViajes() []models.ParticipanteViaje {
	var participanteViajes []models.ParticipanteViaje
	a.db.Find(&participanteViajes)
	return participanteViajes
}

func (a *AlmacenSQLite) BuscarParticipanteViajePorID(id int) (models.ParticipanteViaje, bool) {
	var pv models.ParticipanteViaje
	if err := a.db.First(&pv, id).Error; err != nil {
		return models.ParticipanteViaje{}, false
	}
	return pv, true
}

func (a *AlmacenSQLite) CrearParticipanteViaje(pv models.ParticipanteViaje) models.ParticipanteViaje {
	a.db.Create(&pv)
	return pv
}

func (a *AlmacenSQLite) ActualizarParticipanteViaje(id int, datos models.ParticipanteViaje) (models.ParticipanteViaje, bool) {
	var existente models.ParticipanteViaje
	if err := a.db.First(&existente, id).Error; err != nil {
		return models.ParticipanteViaje{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarParticipanteViaje(id int) bool {
	res := a.db.Delete(&models.ParticipanteViaje{}, id)
	return res.RowsAffected > 0
}

// =========================================================
// SEEDS
// =========================================================

// SembrarSiVacio inserta datos iniciales solo si aún no hay categorías.
// Así no duplicamos datos en cada arranque del servidor.
func (a *AlmacenSQLite) SembrarSiVacio() {
	var n int64
	a.db.Model(&models.ViajeInmediato{}).Count(&n)
	if n > 0 {
		return
	}

	viajeInmediatos := []models.ViajeInmediato{
		{ID: 1, ConductorID: 1, Origen: "Los Esteros", Destino: "Terminal Terrestre", HoraSalida: "07:00", Cupos: 30, Estado: "Disponible"},
		{ID: 2, ConductorID: 2, Origen: "Tarqui", Destino: "Universidad", HoraSalida: "08:00", Cupos: 25, Estado: "Disponible"},
		{ID: 3, ConductorID: 3, Origen: "Centro", Destino: "Hospital General", HoraSalida: "09:30", Cupos: 20, Estado: "Completa"},
		{ID: 4, ConductorID: 4, Origen: "La Pradera", Destino: "Aeropuerto", HoraSalida: "11:00", Cupos: 35, Estado: "Disponible"},
		{ID: 5, ConductorID: 5, Origen: "Manta 2000", Destino: "Centro Comercial", HoraSalida: "14:00", Cupos: 28, Estado: "Disponible"},
		{ID: 6, ConductorID: 6, Origen: "Jocay", Destino: "Terminal Terrestre", HoraSalida: "16:30", Cupos: 22, Estado: "En ruta"},
	}
	a.db.Create(&viajeInmediatos)

	solicitudViajes := []models.SolicitudViaje{
		{ID: 1, ViajeID: 1, PasajeroID: 2, Estado: "Pendiente"},
		{ID: 2, ViajeID: 1, PasajeroID: 3, Estado: "Aceptada"},
		{ID: 3, ViajeID: 2, PasajeroID: 4, Estado: "Aceptada"},
		{ID: 4, ViajeID: 3, PasajeroID: 5, Estado: "Rechazada"},
		{ID: 5, ViajeID: 4, PasajeroID: 6, Estado: "Pendiente"},
		{ID: 6, ViajeID: 5, PasajeroID: 1, Estado: "Aceptada"},
	}
	a.db.Create(&solicitudViajes)

	participanteViajes := []models.ParticipanteViaje{
		{ID: 1, ViajeID: 1, UsuarioID: 2},
		{ID: 2, ViajeID: 1, UsuarioID: 3},
		{ID: 3, ViajeID: 2, UsuarioID: 4},
		{ID: 4, ViajeID: 3, UsuarioID: 5},
		{ID: 5, ViajeID: 4, UsuarioID: 6},
		{ID: 6, ViajeID: 5, UsuarioID: 1},
	}
	a.db.Create(&participanteViajes)

}

// Chequeo en tiempo de compilación: AlmacenSQLite debe cumplir Almacen.
var _ Almacen = (*AlmacenSQLite)(nil)
