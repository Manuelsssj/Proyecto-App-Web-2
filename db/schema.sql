CREATE TABLE IF NOT EXISTS rutas_programadas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    conductor_id INTEGER NOT NULL,
    origen TEXT NOT NULL,
    destino TEXT NOT NULL,
    costo REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS horarios_ruta (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    ruta_id INTEGER NOT NULL,
    dia TEXT NOT NULL,
    hora TEXT NOT NULL,
    FOREIGN KEY (ruta_id) REFERENCES rutas_programadas(id)
);

CREATE TABLE IF NOT EXISTS mantenimientos_vehiculo (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    vehiculo_id INTEGER NOT NULL,
    fecha_inicio TEXT NOT NULL,
    fecha_fin TEXT NOT NULL,
    motivo TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    correo TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);