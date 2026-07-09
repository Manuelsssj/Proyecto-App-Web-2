# RideULEAM

RideULEAM es una API REST para coordinar movilidad compartida dentro de la comunidad universitaria. Permite registrar usuarios, autenticar solicitudes mediante JWT, administrar vehículos, publicar viajes inmediatos y gestionar rutas programadas con horarios y mantenimientos.

## Integrantes y responsabilidades

| Integrante | Módulo principal | Responsabilidades |
| --- | --- | --- |
| Manuel Intriago | Rutas programadas | Rutas programadas, horarios, mantenimientos y persistencia GORM del módulo |
| George Paredes | Viajes inmediatos | Viajes inmediatos, solicitudes, participantes, services y pruebas |
| José Romero | Usuarios y vehículos | Registro, login, roles, JWT, vehículos e integración de autenticación |

La infraestructura compartida incluye Docker, PostgreSQL, GitHub Actions, middleware, documentación e integración de los módulos.

## Stack tecnológico

- Go 1.26.2
- Chi Router
- GORM
- PostgreSQL mediante Docker
- SQLite para desarrollo local
- SQLC como alternativa de persistencia local
- JWT para autenticación y autorización por roles
- Testify para pruebas unitarias y mocks
- Docker y Docker Compose
- GitHub Actions para integración continua

## Arquitectura

El proyecto utiliza arquitectura por capas:

```text
Cliente HTTP / Postman
          |
          v
Chi Router + Middleware JWT/CORS
          |
          v
Handler HTTP
          |
          v
Service / reglas de negocio
          |
          v
Repository / Storage interface
          |
          v
GORM o SQLC
          |
          v
PostgreSQL / SQLite
```

Los handlers reciben y validan solicitudes HTTP. Los services concentran las reglas de negocio. Los repositories definen contratos de persistencia y sus implementaciones trabajan con GORM o SQLC. Las dependencias se construyen e inyectan desde `cmd/rideUleam/main.go`.

El diagrama ampliado se encuentra en [`docs/arquitectura.md`](docs/arquitectura.md).

## Requisitos

- Docker Desktop con Docker Compose, o
- Go 1.26.2 para ejecución local

## Ejecución con Docker

Desde la raíz del proyecto:

```bash
docker compose up --build
```

Este comando construye la API, inicia PostgreSQL, espera a que la base esté disponible, ejecuta las migraciones automáticas y carga los datos iniciales.

La API queda disponible en:

```text
http://localhost:8080
```

Para detener los contenedores:

```bash
docker compose down
```

Para detenerlos y eliminar el volumen de la base de datos:

```bash
docker compose down -v
```

> `docker compose down -v` elimina todos los datos almacenados en PostgreSQL.

## Ejecución local

1. Copiar la configuración de ejemplo:

```bash
cp .env.example .env
```

2. Iniciar la aplicación:

```bash
go run ./cmd/rideUleam
```

Por defecto, la ejecución local utiliza SQLite y crea el archivo configurado mediante `RUTA_DB`.

## Variables de entorno

| Variable | Descripción | Valor de desarrollo |
| --- | --- | --- |
| `PUERTO` | Puerto HTTP de la API | `:8080` |
| `DB_DRIVER` | Motor de base de datos: `sqlite` o `postgres` | `sqlite` |
| `DB_DSN` | Cadena de conexión de PostgreSQL | Vacío |
| `RUTA_DB` | Archivo de SQLite | `rideUleam.db` |
| `STORAGE` | Backend: `gorm` o `sqlc` | `gorm` |
| `JWT_SECRETO` | Secreto utilizado para firmar JWT | Solo desarrollo |
| `JWT_DURACION` | Tiempo de validez del JWT | `24h` |
| `HTTP_READ_TIMEOUT` | Tiempo máximo de lectura HTTP | `10s` |
| `HTTP_WRITE_TIMEOUT` | Tiempo máximo de escritura HTTP | `10s` |

En producción deben utilizarse credenciales y secretos diferentes a los valores de ejemplo.

## Autenticación y roles

El registro y el login son públicos. Los demás endpoints requieren un JWT enviado mediante:

```text
Authorization: Bearer <token>
```

Roles disponibles:

- `pasajero`
- `conductor`
- `admin`

Ejemplo de registro:

```json
{
  "email": "usuario@uleam.edu.ec",
  "password": "secreta123",
  "rol": "pasajero"
}
```

Ejemplo de login:

```json
{
  "email": "usuario@uleam.edu.ec",
  "password": "secreta123"
}
```

El login devuelve:

```json
{
  "token": "<jwt>"
}
```

## Endpoints

### Autenticación — responsable: José Romero

| Método | Ruta | Descripción | Acceso |
| --- | --- | --- | --- |
| POST | `/api/v1/auth/register` | Registrar un usuario | Público |
| POST | `/api/v1/auth/login` | Iniciar sesión y obtener JWT | Público |

### Vehículos — responsable: José Romero

| Método | Ruta | Descripción | Acceso |
| --- | --- | --- | --- |
| GET | `/api/v1/vehiculos` | Listar vehículos | JWT |
| POST | `/api/v1/vehiculos` | Crear vehículo | JWT |
| GET | `/api/v1/vehiculos/{id}` | Obtener vehículo por ID | JWT |
| PUT | `/api/v1/vehiculos/{id}` | Actualizar vehículo | JWT |
| DELETE | `/api/v1/vehiculos/{id}` | Eliminar vehículo | Rol `admin` |

### Viajes inmediatos — responsable: George Paredes

| Método | Ruta | Descripción | Acceso |
| --- | --- | --- | --- |
| GET | `/api/v1/viajes-inmediatos` | Listar viajes inmediatos | JWT |
| POST | `/api/v1/viajes-inmediatos` | Crear viaje inmediato | Rol `admin` o `conductor` |
| GET | `/api/v1/viajes-inmediatos/{id}` | Obtener viaje por ID | JWT |
| PUT | `/api/v1/viajes-inmediatos/{id}` | Actualizar viaje | JWT |
| DELETE | `/api/v1/viajes-inmediatos/{id}` | Eliminar viaje | JWT |

### Solicitudes de viaje — responsable: George Paredes

| Método | Ruta | Descripción | Acceso |
| --- | --- | --- | --- |
| GET | `/api/v1/solicitudes-viajes` | Listar solicitudes | JWT |
| POST | `/api/v1/solicitudes-viajes` | Crear solicitud | JWT |
| GET | `/api/v1/solicitudes-viajes/{id}` | Obtener solicitud por ID | JWT |
| PUT | `/api/v1/solicitudes-viajes/{id}` | Actualizar solicitud | JWT |
| DELETE | `/api/v1/solicitudes-viajes/{id}` | Eliminar solicitud | JWT |

### Participantes de viaje — responsable: George Paredes

| Método | Ruta | Descripción | Acceso |
| --- | --- | --- | --- |
| GET | `/api/v1/participantes-viajes` | Listar participantes | JWT |
| POST | `/api/v1/participantes-viajes` | Crear participante | JWT |
| GET | `/api/v1/participantes-viajes/{id}` | Obtener participante por ID | JWT |
| PUT | `/api/v1/participantes-viajes/{id}` | Actualizar participante | JWT |
| DELETE | `/api/v1/participantes-viajes/{id}` | Eliminar participante | JWT |

### Rutas programadas — responsable: Manuel Intriago

| Método | Ruta | Descripción | Acceso |
| --- | --- | --- | --- |
| GET | `/api/v1/rutas-programadas` | Listar rutas programadas | JWT |
| POST | `/api/v1/rutas-programadas` | Crear ruta programada | JWT |
| GET | `/api/v1/rutas-programadas/{id}` | Obtener ruta por ID | JWT |
| GET | `/api/v1/rutas-programadas/{id}/horarios` | Listar horarios de una ruta | JWT |
| GET | `/api/v1/rutas-programadas/{id}/detalle` | Obtener ruta con detalle | JWT |
| PUT | `/api/v1/rutas-programadas/{id}` | Actualizar ruta | JWT |
| DELETE | `/api/v1/rutas-programadas/{id}` | Eliminar ruta | JWT |

### Horarios de ruta — responsable: Manuel Intriago

| Método | Ruta | Descripción | Acceso |
| --- | --- | --- | --- |
| GET | `/api/v1/horarios-ruta` | Listar horarios | JWT |
| POST | `/api/v1/horarios-ruta` | Crear horario | JWT |
| GET | `/api/v1/horarios-ruta/{id}` | Obtener horario por ID | JWT |
| PUT | `/api/v1/horarios-ruta/{id}` | Actualizar horario | JWT |
| DELETE | `/api/v1/horarios-ruta/{id}` | Eliminar horario | JWT |

### Mantenimientos — responsable: Manuel Intriago

| Método | Ruta | Descripción | Acceso |
| --- | --- | --- | --- |
| GET | `/api/v1/mantenimientos` | Listar mantenimientos | JWT |
| GET | `/api/v1/mantenimientos/{id}` | Obtener mantenimiento por ID | JWT |
| POST | `/api/v1/mantenimientos` | Crear mantenimiento | Rol `admin` |
| PUT | `/api/v1/mantenimientos/{id}` | Actualizar mantenimiento | Rol `admin` |
| DELETE | `/api/v1/mantenimientos/{id}` | Eliminar mantenimiento | Rol `admin` |
| GET | `/api/v1/vehiculos/{vehiculoID}/mantenimientos` | Listar mantenimientos de un vehículo | JWT |

## Pruebas

Ejecutar todas las pruebas:

```bash
go test ./...
```

Ejecutar las pruebas mostrando cobertura:

```bash
go test ./... -cover
```

Ejecutar el análisis estático:

```bash
go vet ./...
```

Las pruebas incluyen services con repositories simulados mediante Testify, validaciones, recursos inexistentes e inputs inválidos.

## CI/CD

El workflow `.github/workflows/ci.yml` se ejecuta en pushes a `main`, ramas `feature/**` y pull requests hacia `main`.

El pipeline realiza:

1. Descarga de dependencias.
2. Compilación con `go build ./...`.
3. Análisis estático con `go vet ./...`.
4. Pruebas y cobertura con `go test ./... -cover`.

## Colecciones Postman

Las colecciones exportadas se encuentran en `postman/`:

- `RideULEAM.postman_collection.json`
- `RideULEAM - Ruta Programada.postman_collection.json`

## Documentación adicional

- [`docs/arquitectura.md`](docs/arquitectura.md): diagrama y explicación de arquitectura.
- [`docs/cierre.md`](docs/cierre.md): aprendizajes, retrospectiva y próximos pasos.

## Estructura principal

```text
cmd/rideUleam/       Punto de entrada e inyección de dependencias
internal/handlers/   Capa HTTP
internal/service/    Reglas de negocio
internal/storage/    Interfaces y persistencia
internal/models/     Entidades y relaciones GORM
internal/middleware/ JWT, roles y CORS
db/                  Esquema y consultas SQLC
docs/                Arquitectura y documento de cierre
postman/             Colecciones para probar la API
```
