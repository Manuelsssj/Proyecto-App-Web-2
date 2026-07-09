# RideULEAM

RideULEAM es una API REST desarrollada en Go para gestionar viajes compartidos entre estudiantes universitarios. Permite registrar usuarios, iniciar sesion con JWT, administrar vehiculos y coordinar viajes inmediatos con solicitudes y participantes.

## Integrantes y responsables

- Modulo Viajes Inmediatos: gestiona viajes, solicitudes y participantes.
- Modulo Usuarios y Vehiculos: gestiona registro, login, roles y vehiculos.
- Infraestructura grupal: Docker, PostgreSQL, CI/CD, middleware JWT y documentacion.

## Stack

- Go
- Chi Router
- GORM
- PostgreSQL en Docker
- SQLite para desarrollo local opcional
- JWT con roles basicos
- Testify para pruebas unitarias y mocks
- GitHub Actions para CI

## Arquitectura

El proyecto usa arquitectura por capas:

```text
HTTP handler -> service -> repository/storage -> database
```

La inyeccion de dependencias se realiza desde `cmd/rideUleam/main.go`. La inicializacion de base de datos, migraciones y seeders esta en `internal/storage/factory.go`.

## Ejecucion con Docker

```bash
docker compose up --build
```

La API queda disponible en:

```text
http://localhost:8080
```

El `docker-compose.yml` levanta:

- API RideULEAM
- PostgreSQL
- Seeders internos mediante `SembrarSiVacio()`

## Ejecucion local

```bash
go run ./cmd/rideUleam
```

## Autenticacion

Primero registra un usuario y luego inicia sesion. El login devuelve un token JWT.

Roles disponibles:

- `pasajero`
- `conductor`
- `admin`

Ejemplo de registro:

```json
{
  "email": "admin@uleam.edu.ec",
  "password": "secreta123",
  "rol": "admin"
}
```

Para rutas protegidas se usa:

```text
Authorization: Bearer <token>
```

## Endpoints

### Auth

| Metodo | Ruta | Descripcion | Responsable |
| --- | --- | --- | --- |
| POST | `/api/v1/auth/register` | Registrar usuario con rol | Usuarios y Vehiculos |
| POST | `/api/v1/auth/login` | Iniciar sesion y obtener JWT | Usuarios y Vehiculos |

### Viajes inmediatos

| Metodo | Ruta | Descripcion | Responsable |
| --- | --- | --- | --- |
| GET | `/api/v1/viajes-inmediatos` | Listar viajes | Viajes Inmediatos |
| POST | `/api/v1/viajes-inmediatos` | Crear viaje, requiere rol `admin` o `conductor` | Viajes Inmediatos |
| GET | `/api/v1/viajes-inmediatos/{id}` | Obtener viaje por ID | Viajes Inmediatos |
| PUT | `/api/v1/viajes-inmediatos/{id}` | Actualizar viaje | Viajes Inmediatos |
| DELETE | `/api/v1/viajes-inmediatos/{id}` | Eliminar viaje | Viajes Inmediatos |

### Solicitudes de viaje

| Metodo | Ruta | Descripcion | Responsable |
| --- | --- | --- | --- |
| GET | `/api/v1/solicitudes-viajes` | Listar solicitudes | Viajes Inmediatos |
| POST | `/api/v1/solicitudes-viajes` | Crear solicitud | Viajes Inmediatos |
| GET | `/api/v1/solicitudes-viajes/{id}` | Obtener solicitud por ID | Viajes Inmediatos |
| PUT | `/api/v1/solicitudes-viajes/{id}` | Actualizar solicitud | Viajes Inmediatos |
| DELETE | `/api/v1/solicitudes-viajes/{id}` | Eliminar solicitud | Viajes Inmediatos |

### Participantes de viaje

| Metodo | Ruta | Descripcion | Responsable |
| --- | --- | --- | --- |
| GET | `/api/v1/participantes-viajes` | Listar participantes | Viajes Inmediatos |
| POST | `/api/v1/participantes-viajes` | Crear participante | Viajes Inmediatos |
| GET | `/api/v1/participantes-viajes/{id}` | Obtener participante por ID | Viajes Inmediatos |
| PUT | `/api/v1/participantes-viajes/{id}` | Actualizar participante | Viajes Inmediatos |
| DELETE | `/api/v1/participantes-viajes/{id}` | Eliminar participante | Viajes Inmediatos |

### Vehiculos

| Metodo | Ruta | Descripcion | Responsable |
| --- | --- | --- | --- |
| GET | `/api/v1/vehiculos` | Listar vehiculos | Usuarios y Vehiculos |
| POST | `/api/v1/vehiculos` | Crear vehiculo | Usuarios y Vehiculos |
| GET | `/api/v1/vehiculos/{id}` | Obtener vehiculo por ID | Usuarios y Vehiculos |
| PUT | `/api/v1/vehiculos/{id}` | Actualizar vehiculo | Usuarios y Vehiculos |
| DELETE | `/api/v1/vehiculos/{id}` | Eliminar vehiculo, requiere rol `admin` | Usuarios y Vehiculos |

## Pruebas

```bash
go test ./... -cover
```

## CI/CD

El pipeline esta en `.github/workflows/ci.yml` y ejecuta:

- `go mod download`
- `go build ./...`
- `go vet ./...`
- `go test ./... -cover`
