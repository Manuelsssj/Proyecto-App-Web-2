# suscripciones-api

API REST en Go para gestionar suscripciones a rutas, planes de pago e historial de suscripciones.
Arquitectura por capas (models / storage / service / handlers / middleware) con dos backends de
almacenamiento intercambiables: **memoria** y **SQLite**.

## Estructura del proyecto

```
suscripciones-api/
├── cmd/
│   └── suscripciones-api/
│       └── main.go              # Punto de entrada: configura storage, servicios, handlers y servidor
├── internal/
│   ├── models/                  # Structs de dominio
│   │   ├── suscripcionruta.go
│   │   ├── planpago.go
│   │   └── historialsuscripcion.go
│   ├── storage/                 # Persistencia (interfaz + implementaciones)
│   │   ├── almacen.go           # Interfaz Almacen
│   │   ├── memoria.go           # Implementación en memoria
│   │   └── sqlite.go            # Implementación SQLite
│   ├── service/                 # Lógica de negocio y validaciones
│   │   ├── errores.go
│   │   ├── suscripcion.go
│   │   ├── plan.go
│   │   └── historial.go
│   ├── handlers/                # Capa HTTP (controllers)
│   │   ├── respond.go           # Helpers de respuesta JSON
│   │   ├── suscripcion.go
│   │   ├── plan.go
│   │   └── historial.go
│   └── middleware/
│       └── cors.go
├── db/
│   ├── schema.sql               # Definición de tablas
│   └── queries.sql              # Consultas SQL de referencia
├── go.mod
└── README.md
```

## Requisitos

- Go 1.22 o superior.

## Instalación y ejecución

```bash
# 1. Descargar dependencias (genera go.sum automáticamente)
go mod tidy

# 2a. Ejecutar con almacenamiento en MEMORIA (por defecto)
go run ./cmd/suscripciones-api

# 2b. Ejecutar con almacenamiento en SQLITE
go run ./cmd/suscripciones-api -storage=sqlite -db=suscripciones.db

# Opcional: cambiar el puerto
go run ./cmd/suscripciones-api -puerto=:9090
```

El servidor queda escuchando en `http://localhost:8080`.

## Endpoints

Cada entidad expone un CRUD completo.

### Suscripciones (`/suscripciones`)

| Método | Ruta                 | Descripción              |
|--------|----------------------|--------------------------|
| GET    | /suscripciones       | Listar todas             |
| POST   | /suscripciones       | Crear una nueva          |
| GET    | /suscripciones/{id}  | Obtener por ID           |
| PUT    | /suscripciones/{id}  | Actualizar por ID        |
| DELETE | /suscripciones/{id}  | Eliminar por ID          |

### Planes de pago (`/planes`)

| Método | Ruta           | Descripción       |
|--------|----------------|-------------------|
| GET    | /planes        | Listar todos      |
| POST   | /planes        | Crear uno nuevo   |
| GET    | /planes/{id}   | Obtener por ID    |
| PUT    | /planes/{id}   | Actualizar por ID |
| DELETE | /planes/{id}   | Eliminar por ID   |

### Historial (`/historial`)

| Método | Ruta              | Descripción       |
|--------|-------------------|-------------------|
| GET    | /historial        | Listar todo       |
| POST   | /historial        | Crear uno nuevo   |
| GET    | /historial/{id}   | Obtener por ID    |
| PUT    | /historial/{id}   | Actualizar por ID |
| DELETE | /historial/{id}   | Eliminar por ID   |

También hay un endpoint de salud: `GET /health`.

## Ejemplos con curl

```bash
# Crear una suscripción
curl -X POST http://localhost:8080/suscripciones \
  -H "Content-Type: application/json" \
  -d '{"ruta_id": 1, "usuario_id": 5}'

# Listar suscripciones
curl http://localhost:8080/suscripciones

# Crear un plan de pago
curl -X POST http://localhost:8080/planes \
  -H "Content-Type: application/json" \
  -d '{"ruta_id": 1, "valor_semanal": 12.50}'

# Crear un historial (si no envías fecha_registro, usa la fecha actual)
curl -X POST http://localhost:8080/historial \
  -H "Content-Type: application/json" \
  -d '{"suscripcion_id": 1, "estado": "activa"}'

# Actualizar una suscripción
curl -X PUT http://localhost:8080/suscripciones/1 \
  -H "Content-Type: application/json" \
  -d '{"ruta_id": 2, "usuario_id": 5}'

# Eliminar un plan
curl -X DELETE http://localhost:8080/planes/1
```

## Notas de diseño

- **Almacen** es una interfaz, por lo que memoria y SQLite son intercambiables sin tocar los servicios.
- Los **services** contienen las validaciones de negocio y devuelven errores tipados (`ErrNoEncontrado`,
  `ErrDatosInvalidos`) que los handlers traducen a códigos HTTP.
- SQLite usa `modernc.org/sqlite`, un driver en Go puro (no requiere CGO ni compilador de C).
