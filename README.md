# 🚍 RideUleam API

API REST desarrollada en **Go** para la gestión de suscripciones de rutas, planes de pago e historial de suscripciones del sistema **RideUleam**.

El proyecto fue desarrollado siguiendo una **arquitectura por capas**, utilizando **GORM**, **PostgreSQL**, **Docker**, autenticación mediante **JWT** y un flujo de integración continua con **GitHub Actions**.

---

# 👥 Integrantes

- Manuel Intriago 
- Andrés Romero
- Joseph Paredes

---

# 🛠 Tecnologías utilizadas

- Go 1.25
- PostgreSQL 16
- GORM
- Chi Router
- JWT (golang-jwt)
- Docker
- Docker Compose
- GitHub Actions

---

# Características

- API REST desarrollada en Go.
- Arquitectura por capas.
- Persistencia con PostgreSQL mediante GORM.
- Autenticación con JWT.
- Middleware CORS.
- Docker y Docker Compose.
- GitHub Actions para Integración Continua.
- Seeder automático.
- CRUD completo para:
  - Suscripciones
  - Planes
  - Historial

---

# 🏗 Arquitectura del proyecto

El proyecto sigue una arquitectura por capas para separar responsabilidades.

```
Cliente
    │
HTTP Request
    │
Handlers
    │
Services
    │
Repositories (Storage)
    │
PostgreSQL
```

Cada capa tiene una responsabilidad específica:

- **Handlers:** reciben las peticiones HTTP y generan las respuestas.
- **Services:** contienen la lógica de negocio.
- **Storage:** administra el acceso a la base de datos mediante GORM.
- **Models:** representan las entidades del sistema.
- **Middleware:** autenticación JWT y configuración CORS.

---

# 📁 Estructura del proyecto

```
Proyecto-App-Web-2
│
├── cmd/
│   └── rideUleam/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── service/
│   └── storage/
│
├── db/
│   ├── schema.sql
│   └── queries.sql
│
├── Dockerfile
├── docker-compose.yml
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

---

# ⚙ Variables de entorno

El proyecto utiliza las siguientes variables de entorno:

| Variable | Descripción |
|----------|-------------|
| DATABASE_URL | Cadena de conexión a PostgreSQL |
| PORT | Puerto donde se ejecuta la API |
| JWT_SECRETO | Clave utilizada para firmar los tokens JWT |

Ejemplo:

```env
PORT=8080
DATABASE_URL=host=db user=postgres password=postgres dbname=rideuleam port=5432 sslmode=disable
JWT_SECRETO=rideuleam-secreto-dev
```

---

# 🚀 Ejecución local

Instalar dependencias

```bash
go mod tidy
```

Ejecutar la aplicación

```bash
go run ./cmd/rideUleam
```

La API quedará disponible en

```
http://localhost:8080
```

---

# 🐳 Ejecución con Docker

Construir y levantar los contenedores

```bash
docker compose up --build
```

Detener los contenedores

```bash
docker compose down
```

La aplicación quedará disponible en

```
http://localhost:8080
```

---

# 🗄 Base de datos

El proyecto utiliza **PostgreSQL** como sistema gestor de base de datos.

Al iniciar la aplicación:

- Se establece la conexión mediante GORM.
- Se ejecutan automáticamente las migraciones (`AutoMigrate`).
- Se cargan datos iniciales mediante un Seeder.

---

# 🔐 Autenticación

La autenticación se realiza utilizando **JSON Web Tokens (JWT)**.

Endpoints públicos:

```
POST /api/v1/auth/register

POST /api/v1/auth/login
```

Los demás endpoints requieren un token JWT válido.

---

# 📌 Endpoints principales

## Autenticación

| Método | Endpoint |
|---------|----------|
| POST | /api/v1/auth/register |
| POST | /api/v1/auth/login |

## Suscripciones

| Método | Endpoint |
|---------|----------|
| GET | /api/v1/suscripciones |
| POST | /api/v1/suscripciones |
| GET | /api/v1/suscripciones/{id} |
| PUT | /api/v1/suscripciones/{id} |
| DELETE | /api/v1/suscripciones/{id} |

## Planes

| Método | Endpoint |
|---------|----------|
| GET | /api/v1/planes |
| POST | /api/v1/planes |
| GET | /api/v1/planes/{id} |
| PUT | /api/v1/planes/{id} |
| DELETE | /api/v1/planes/{id} |

## Historial

| Método | Endpoint |
|---------|----------|
| GET | /api/v1/historial |
| POST | /api/v1/historial |
| GET | /api/v1/historial/{id} |
| PUT | /api/v1/historial/{id} |
| DELETE | /api/v1/historial/{id} |

## Estado del servidor

```
GET /health
```

---

# 🌱 Seeder

Durante el arranque de la aplicación se insertan datos iniciales en la base de datos cuando las tablas se encuentran vacías.

Esto facilita las pruebas del sistema.

---

# 🧪 Pruebas

Ejecutar todas las pruebas

```bash
go test ./...
```

Ejecutar pruebas con cobertura

```bash
go test ./... --cover
```

---

# ⚙ Integración Continua

El proyecto utiliza **GitHub Actions** para ejecutar automáticamente:

- Descarga de dependencias
- Verificación del código (`go vet`)
- Ejecución de pruebas (`go test`)
- Cobertura de pruebas

Cada cambio enviado al repositorio ejecuta automáticamente el pipeline de integración continua.

---

# 📄 Licencia

Proyecto académico desarrollado para la asignatura **Aplicaciones Web II**.