-- Consultas SQL de referencia para suscripciones-api.

-- ===== SuscripcionRuta =====
-- ListarSuscripciones
SELECT id, ruta_id, usuario_id FROM suscripciones_ruta ORDER BY id;

-- ObtenerSuscripcion
SELECT id, ruta_id, usuario_id FROM suscripciones_ruta WHERE id = ?;

-- CrearSuscripcion
INSERT INTO suscripciones_ruta (ruta_id, usuario_id) VALUES (?, ?);

-- ActualizarSuscripcion
UPDATE suscripciones_ruta SET ruta_id = ?, usuario_id = ? WHERE id = ?;

-- EliminarSuscripcion
DELETE FROM suscripciones_ruta WHERE id = ?;

-- ===== PlanPago =====
-- ListarPlanes
SELECT id, ruta_id, valor_semanal FROM planes_pago ORDER BY id;

-- ObtenerPlan
SELECT id, ruta_id, valor_semanal FROM planes_pago WHERE id = ?;

-- CrearPlan
INSERT INTO planes_pago (ruta_id, valor_semanal) VALUES (?, ?);

-- ActualizarPlan
UPDATE planes_pago SET ruta_id = ?, valor_semanal = ? WHERE id = ?;

-- EliminarPlan
DELETE FROM planes_pago WHERE id = ?;

-- ===== HistorialSuscripcion =====
-- ListarHistorial
SELECT id, suscripcion_id, fecha_registro, estado FROM historial_suscripcion ORDER BY id;

-- ObtenerHistorial
SELECT id, suscripcion_id, fecha_registro, estado FROM historial_suscripcion WHERE id = ?;

-- CrearHistorial
INSERT INTO historial_suscripcion (suscripcion_id, fecha_registro, estado) VALUES (?, ?, ?);

-- ActualizarHistorial
UPDATE historial_suscripcion SET suscripcion_id = ?, fecha_registro = ?, estado = ? WHERE id = ?;

-- EliminarHistorial
DELETE FROM historial_suscripcion WHERE id = ?;
