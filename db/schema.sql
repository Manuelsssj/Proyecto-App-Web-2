-- Esquema de la base de datos de suscripciones-api.

CREATE TABLE IF NOT EXISTS suscripciones_ruta (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    ruta_id    INTEGER NOT NULL,
    usuario_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS planes_pago (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    ruta_id       INTEGER NOT NULL,
    valor_semanal REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS historial_suscripcion (
    id             INTEGER PRIMARY KEY AUTOINCREMENT,
    suscripcion_id INTEGER NOT NULL,
    fecha_registro TEXT NOT NULL,
    estado         TEXT NOT NULL
);
