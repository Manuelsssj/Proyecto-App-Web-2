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