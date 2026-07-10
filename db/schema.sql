
-- ==========================================
-- Módulo 1: Viajes Inmediatos
-- ==========================================

CREATE TABLE viaje_inmediatos (
    id            INTEGER PRIMARY KEY,
    conductor_id  INTEGER NOT NULL,
    origen        TEXT NOT NULL,
    destino       TEXT NOT NULL,
    hora_salida   TEXT NOT NULL,
    cupos         INTEGER NOT NULL,
    estado        TEXT NOT NULL
);

CREATE TABLE solicitud_viajes (
    id            INTEGER PRIMARY KEY,
    viaje_id      INTEGER NOT NULL,
    pasajero_id   INTEGER NOT NULL,
    estado        TEXT NOT NULL
);

CREATE TABLE participante_viajes (
    id            INTEGER PRIMARY KEY,
    viaje_id      INTEGER NOT NULL,
    usuario_id    INTEGER NOT NULL
);

-- ==========================================
-- Módulo 4: Usuarios y Vehículos
-- ==========================================



CREATE TABLE vehiculos (
    id            INTEGER PRIMARY KEY,
    conductor_id  INTEGER NOT NULL,
    placa         TEXT NOT NULL,
    marca         TEXT NOT NULL,
    modelo        TEXT NOT NULL,
    capacidad     INTEGER NOT NULL
);