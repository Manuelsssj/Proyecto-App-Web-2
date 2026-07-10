
-- ==========================================
-- MÓDULO 1: VIAJES INMEDIATOS
-- ==========================================

-- VIAJES INMEDIATOS

-- name: ListarViajesInmediatos :many
SELECT id, conductor_id, origen, destino, hora_salida, cupos, estado FROM viaje_inmediatos;

-- name: BuscarViajeInmediatoPorID :one
SELECT id,conductor_id,origen,destino, hora_salida, cupos, estado FROM viaje_inmediatos
WHERE id = ?;

-- name: CrearViajeInmediato :one
INSERT INTO viaje_inmediatos (conductor_id, origen, destino, hora_salida, cupos, estado)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id, conductor_id, origen, destino, hora_salida, cupos, estado;

-- name: ActualizarViajeInmediato :one
UPDATE viaje_inmediatos
SET conductor_id=?, origen=?, destino=?, hora_salida=?, cupos=?, estado=?
WHERE id=?
RETURNING id, conductor_id, origen, destino, hora_salida, cupos, estado;

-- name: BorrarViajeInmediato :execrows
DELETE FROM viaje_inmediatos WHERE id=?;

-- SOLICITUDES VIAJES

-- name: ListarSolicitudesViajes :many
SELECT id, viaje_id, pasajero_id, estado FROM solicitud_viajes;

-- name: BuscarSolicitudViajePorID :one
SELECT id, viaje_id, pasajero_id, estado FROM solicitud_viajes
WHERE id=?;

-- name: CrearSolicitudViaje :one
INSERT INTO solicitud_viajes (viaje_id, pasajero_id, estado)
VALUES (?, ?, ?)
RETURNING id, viaje_id, pasajero_id, estado;

-- name: ActualizarSolicitudViaje :one
UPDATE solicitud_viajes
SET viaje_id=?, pasajero_id=?, estado=?
WHERE id=?
RETURNING id, viaje_id, pasajero_id, estado;

-- name: BorrarSolicitudViaje :execrows
DELETE FROM solicitud_viajes WHERE id=?;

-- PARTICIPANTES VIAJES

-- name: ListarParticipantesViajes :many
SELECT id, viaje_id, usuario_id FROM participante_viajes;

-- name: BuscarParticipanteViajePorID :one
SELECT id, viaje_id, usuario_id FROM participante_viajes
WHERE id=?;

-- name: CrearParticipanteViaje :one
INSERT INTO participante_viajes (viaje_id, usuario_id)
VALUES (?, ?)
RETURNING id, viaje_id, usuario_id;

-- name: ActualizarParticipanteViaje :one
UPDATE participante_viajes
SET viaje_id=?, usuario_id=?
WHERE id=?
RETURNING id, viaje_id, usuario_id;

-- name: BorrarParticipanteViaje :execrows
DELETE FROM participante_viajes WHERE id=?;

-- ==========================================
-- MÓDULO 4: USUARIOS Y VEHICULOS
-- ==========================================



-- VEHICULOS

-- name: ListarVehiculos :many
SELECT id, conductor_id, placa, marca, modelo, capacidad FROM vehiculos;

-- name: BuscarVehiculoPorID :one
SELECT id, conductor_id, placa, marca, modelo, capacidad FROM vehiculos
WHERE id=?;

-- name: CrearVehiculo :one
INSERT INTO vehiculos (conductor_id, placa, marca, modelo, capacidad)
VALUES (?, ?, ?, ?, ?)
RETURNING id, conductor_id, placa, marca, modelo, capacidad;

-- name: ActualizarVehiculo :one
UPDATE vehiculos
SET conductor_id=?, placa=?, marca=?, modelo=?, capacidad=?
WHERE id=?
RETURNING id, conductor_id, placa, marca, modelo, capacidad;

-- name: BorrarVehiculo :execrows
DELETE FROM vehiculos WHERE id=?;
