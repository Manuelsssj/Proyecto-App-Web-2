
-- RUTAS PROGRAMADAS

-- name: ListarRutasProgramadas :many
SELECT id, conductor_id, origen, destino, costo
FROM rutas_programadas;

-- name: ObtenerRutaProgramada :one
SELECT id, conductor_id, origen, destino, costo
FROM rutas_programadas
WHERE id = ?;

-- name: CrearRutaProgramada :one
INSERT INTO rutas_programadas (conductor_id, origen, destino, costo)
VALUES (?, ?, ?, ?)
RETURNING id, conductor_id, origen, destino, costo;

-- name: ActualizarRutaProgramada :one
UPDATE rutas_programadas
SET conductor_id = ?, origen = ?, destino = ?, costo = ?
WHERE id = ?
RETURNING id, conductor_id, origen, destino, costo;

-- name: BorrarRutaProgramada :exec
DELETE FROM rutas_programadas
WHERE id = ?;


-- HORARIOS DE RUTA

-- name: ListarHorariosRuta :many
SELECT id, ruta_id, dia, hora
FROM horarios_ruta;

-- name: ObtenerHorarioRuta :one
SELECT id, ruta_id, dia, hora
FROM horarios_ruta
WHERE id = ?;

-- name: ListarHorariosPorRutaID :many
SELECT id, ruta_id, dia, hora
FROM horarios_ruta
WHERE ruta_id = ?;

-- name: CrearHorarioRuta :one
INSERT INTO horarios_ruta (ruta_id, dia, hora)
VALUES (?, ?, ?)
RETURNING id, ruta_id, dia, hora;

-- name: ActualizarHorarioRuta :one
UPDATE horarios_ruta
SET ruta_id = ?, dia = ?, hora = ?
WHERE id = ?
RETURNING id, ruta_id, dia, hora;

-- name: BorrarHorarioRuta :exec
DELETE FROM horarios_ruta
WHERE id = ?;


-- MANTENIMIENTOS VEHICULO

-- name: ListarMantenimientosVehiculo :many
SELECT id, vehiculo_id, fecha_inicio, fecha_fin, motivo
FROM mantenimientos_vehiculo;

-- name: ObtenerMantenimientoVehiculo :one
SELECT id, vehiculo_id, fecha_inicio, fecha_fin, motivo
FROM mantenimientos_vehiculo
WHERE id = ?;

-- name: ListarMantenimientosPorVehiculoID :many
SELECT id, vehiculo_id, fecha_inicio, fecha_fin, motivo
FROM mantenimientos_vehiculo
WHERE vehiculo_id = ?;

-- name: CrearMantenimientoVehiculo :one
INSERT INTO mantenimientos_vehiculo (vehiculo_id, fecha_inicio, fecha_fin, motivo)
VALUES (?, ?, ?, ?)
RETURNING id, vehiculo_id, fecha_inicio, fecha_fin, motivo;

-- name: ActualizarMantenimientoVehiculo :one
UPDATE mantenimientos_vehiculo
SET vehiculo_id = ?, fecha_inicio = ?, fecha_fin = ?, motivo = ?
WHERE id = ?
RETURNING id, vehiculo_id, fecha_inicio, fecha_fin, motivo;

-- name: BorrarMantenimientoVehiculo :exec
DELETE FROM mantenimientos_vehiculo
WHERE id = ?;

-- USUARIOS / AUTH

-- name: CrearUsuario :one
INSERT INTO usuarios (nombre, correo, password)
VALUES (?, ?, ?)
RETURNING id, nombre, correo, password;

-- name: ObtenerUsuarioPorCorreo :one
SELECT id, nombre, correo, password
FROM usuarios
WHERE correo = ?;

-- name: ObtenerUsuarioPorID :one
SELECT id, nombre, correo, password
FROM usuarios
WHERE id = ?;
=======

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

