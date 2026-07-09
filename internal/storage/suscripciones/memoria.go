package storage

import (
	"sync"

	"suscripciones-api/internal/models/suscripciones"
)

// AlmacenMemoria es una implementación de Almacen que guarda los datos en memoria.
// Es segura para uso concurrente gracias al mutex.
type AlmacenMemoria struct {
	mu sync.Mutex

	suscripciones map[int]models.SuscripcionRuta
	planes        map[int]models.PlanPago
	historial     map[int]models.HistorialSuscripcion

	seqSuscripcion int
	seqPlan        int
	seqHistorial   int
}

// NewAlmacenMemoria crea un nuevo almacén en memoria vacío.
func NewAlmacenMemoria() *AlmacenMemoria {
	return &AlmacenMemoria{
		suscripciones: make(map[int]models.SuscripcionRuta),
		planes:        make(map[int]models.PlanPago),
		historial:     make(map[int]models.HistorialSuscripcion),
	}
}

// ---------- SuscripcionRuta ----------

func (a *AlmacenMemoria) ListarSuscripciones() ([]models.SuscripcionRuta, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	lista := make([]models.SuscripcionRuta, 0, len(a.suscripciones))
	for _, s := range a.suscripciones {
		lista = append(lista, s)
	}
	return lista, nil
}

func (a *AlmacenMemoria) ObtenerSuscripcion(id int) (models.SuscripcionRuta, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	s, ok := a.suscripciones[id]
	if !ok {
		return models.SuscripcionRuta{}, ErrNoEncontrado
	}
	return s, nil
}

func (a *AlmacenMemoria) CrearSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.seqSuscripcion++
	s.ID = uint(a.seqSuscripcion)
	a.suscripciones[int(s.ID)] = s
	return s, nil
}

func (a *AlmacenMemoria) ActualizarSuscripcion(s models.SuscripcionRuta) (models.SuscripcionRuta, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.suscripciones[int(s.ID)]; !ok {
		return models.SuscripcionRuta{}, ErrNoEncontrado
	}
	a.suscripciones[int(s.ID)] = s
	return s, nil
}

func (a *AlmacenMemoria) EliminarSuscripcion(id int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.suscripciones[id]; !ok {
		return ErrNoEncontrado
	}
	delete(a.suscripciones, id)
	return nil
}

// ---------- PlanPago ----------

func (a *AlmacenMemoria) ListarPlanes() ([]models.PlanPago, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	lista := make([]models.PlanPago, 0, len(a.planes))
	for _, p := range a.planes {
		lista = append(lista, p)
	}
	return lista, nil
}

func (a *AlmacenMemoria) ObtenerPlan(id int) (models.PlanPago, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	p, ok := a.planes[id]
	if !ok {
		return models.PlanPago{}, ErrNoEncontrado
	}
	return p, nil
}

func (a *AlmacenMemoria) CrearPlan(p models.PlanPago) (models.PlanPago, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.seqPlan++
	p.ID = uint(a.seqPlan)
	a.planes[int(p.ID)] = p
	return p, nil
}

func (a *AlmacenMemoria) ActualizarPlan(p models.PlanPago) (models.PlanPago, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.planes[int(p.ID)]; !ok {
		return models.PlanPago{}, ErrNoEncontrado
	}
	a.planes[int(p.ID)] = p
	return p, nil
}

func (a *AlmacenMemoria) EliminarPlan(id int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.planes[id]; !ok {
		return ErrNoEncontrado
	}
	delete(a.planes, id)
	return nil
}

// ---------- HistorialSuscripcion ----------

func (a *AlmacenMemoria) ListarHistorial() ([]models.HistorialSuscripcion, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	lista := make([]models.HistorialSuscripcion, 0, len(a.historial))
	for _, h := range a.historial {
		lista = append(lista, h)
	}
	return lista, nil
}

func (a *AlmacenMemoria) ObtenerHistorial(id int) (models.HistorialSuscripcion, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	h, ok := a.historial[id]
	if !ok {
		return models.HistorialSuscripcion{}, ErrNoEncontrado
	}
	return h, nil
}

func (a *AlmacenMemoria) CrearHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.seqHistorial++
	h.ID = uint(a.seqHistorial)
	a.historial[int(h.ID)] = h
	return h, nil
}

func (a *AlmacenMemoria) ActualizarHistorial(h models.HistorialSuscripcion) (models.HistorialSuscripcion, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.historial[int(h.ID)]; !ok {
		return models.HistorialSuscripcion{}, ErrNoEncontrado
	}
	a.historial[int(h.ID)] = h
	return h, nil
}

func (a *AlmacenMemoria) EliminarHistorial(id int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.historial[id]; !ok {
		return ErrNoEncontrado
	}
	delete(a.historial, id)
	return nil
}

// Verificación en tiempo de compilación de que AlmacenMemoria implementa Almacen.
var _ Almacen = (*AlmacenMemoria)(nil)
