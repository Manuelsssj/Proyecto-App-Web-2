# RideULEAM - Módulo Ruta Programada

Este README documenta el módulo **Ruta Programada**, desarrollado por **Manuel Intriago** para el proyecto RideULEAM de la materia **Aplicaciones Web II**.

El módulo permite gestionar rutas programadas dentro del sistema de transporte universitario RideULEAM. También incluye la gestión de horarios asociados a una ruta y mantenimientos de vehículos, usando autenticación JWT, roles, persistencia con GORM y pruebas automatizadas.

---

## Responsable del módulo

| Integrante | Rama | Módulo |
|---|---|---|
| Manuel Intriago | `feature/ruta-programada` | Ruta Programada |

---

## Descripción del módulo

El módulo **Ruta Programada** permite registrar rutas planificadas por un conductor, indicando origen, destino y costo. Además, permite asociar horarios a cada ruta programada y registrar mantenimientos de vehículos.

Las entidades principales trabajadas en este módulo son:

- `RutaProgramada`
- `HorarioRuta`
- `MantenimientoVehiculo`

---

## Stack utilizado

El módulo fue desarrollado con:

- **Go**
- **Chi Router**
- **GORM**
- **PostgreSQL**
- **Docker**
- **Docker Compose**
- **JWT**
- **Testify**
- **GitHub Actions**

---

## Arquitectura del módulo

El módulo utiliza arquitectura por capas:

```text
Handler → Service → Repository/Storage → Base de datos
```

### Descripción de capas

- **Handler:** recibe las peticiones HTTP, decodifica el JSON y responde al cliente.
- **Service:** contiene las validaciones y reglas de negocio antes de guardar o consultar datos.
- **Repository/Storage:** se encarga del acceso a datos usando GORM.
- **Base de datos:** PostgreSQL ejecutado mediante Docker Compose.

Esta arquitectura permite separar responsabilidades, facilitar las pruebas y mantener el código más ordenado.

---

## Ejecución del proyecto

Desde la raíz del proyecto, ejecutar:

```bash
docker compose up
```

La API estará disponible en:

```text
http://localhost:8080
```

Para verificar que el servidor funciona correctamente:

```bash
curl http://localhost:8080/health
```

Respuesta esperada:

```json
{
  "status": "ok",
  "message": "Servidor funcionando correctamente"
}
```

Para detener los contenedores:

```bash
docker compose down
```

---

## Autenticación JWT

Las rutas del módulo están protegidas con JWT.

Primero se debe registrar o iniciar sesión para obtener un token.

### Registro de administrador

```http
POST /api/v1/auth/register
```

Body:

```json
{
  "email": "admin@test.com",
  "password": "123456",
  "rol": "admin"
}
```

### Registro de conductor

```http
POST /api/v1/auth/register
```

Body:

```json
{
  "email": "conductor@test.com",
  "password": "123456",
  "rol": "conductor"
}
```

### Login

```http
POST /api/v1/auth/login
```

Body:

```json
{
  "email": "admin@test.com",
  "password": "123456"
}
```

La respuesta devuelve un token JWT.  
Ese token debe enviarse en las rutas protegidas usando el header:

```http
Authorization: Bearer TOKEN
```

---

# Endpoints del módulo Ruta Programada

## Health Check

| Método | Endpoint | Descripción |
|---|---|---|
| GET | `/health` | Verifica que el servidor esté funcionando |

---

## Rutas Programadas

| Método | Endpoint | Descripción | Protección |
|---|---|---|---|
| GET | `/api/v1/rutas-programadas` | Lista todas las rutas programadas | Requiere token |
| POST | `/api/v1/rutas-programadas` | Crea una ruta programada | Requiere token |
| GET | `/api/v1/rutas-programadas/{id}` | Obtiene una ruta programada por ID | Requiere token |
| GET | `/api/v1/rutas-programadas/{id}/detalle` | Obtiene el detalle de una ruta programada | Requiere token |
| GET | `/api/v1/rutas-programadas/{id}/horarios` | Lista los horarios asociados a una ruta | Requiere token |
| PUT | `/api/v1/rutas-programadas/{id}` | Actualiza una ruta programada | Requiere token |
| DELETE | `/api/v1/rutas-programadas/{id}` | Elimina una ruta programada | Requiere token |

### Crear ruta programada

```http
POST /api/v1/rutas-programadas
```

Body:

```json
{
  "conductor_id": 1,
  "origen": "Los Esteros",
  "destino": "ULEAM",
  "costo": 0.75
}
```

Respuesta esperada:

```json
{
  "id": 1,
  "conductor_id": 1,
  "origen": "Los Esteros",
  "destino": "ULEAM",
  "costo": 0.75
}
```

---

## Horarios de Ruta

| Método | Endpoint | Descripción | Protección |
|---|---|---|---|
| GET | `/api/v1/horarios-ruta` | Lista todos los horarios de ruta | Requiere token |
| POST | `/api/v1/horarios-ruta` | Crea un horario de ruta | Requiere token |
| GET | `/api/v1/horarios-ruta/{id}` | Obtiene un horario por ID | Requiere token |
| PUT | `/api/v1/horarios-ruta/{id}` | Actualiza un horario | Requiere token |
| DELETE | `/api/v1/horarios-ruta/{id}` | Elimina un horario | Requiere token |

### Crear horario de ruta

```http
POST /api/v1/horarios-ruta
```

Body:

```json
{
  "ruta_id": 1,
  "dia": "Lunes",
  "hora": "07:00"
}
```

Nota: antes de crear un horario debe existir una ruta programada con el ID indicado en `ruta_id`.

---

## Mantenimientos de Vehículo

| Método | Endpoint | Descripción | Protección |
|---|---|---|---|
| GET | `/api/v1/mantenimientos` | Lista todos los mantenimientos | Requiere token |
| POST | `/api/v1/mantenimientos` | Crea un mantenimiento de vehículo | Solo admin |
| GET | `/api/v1/mantenimientos/{id}` | Obtiene un mantenimiento por ID | Requiere token |
| PUT | `/api/v1/mantenimientos/{id}` | Actualiza un mantenimiento | Solo admin |
| DELETE | `/api/v1/mantenimientos/{id}` | Elimina un mantenimiento | Solo admin |
| GET | `/api/v1/vehiculos/{vehiculoID}/mantenimientos` | Lista mantenimientos por vehículo | Requiere token |

### Crear mantenimiento

```http
POST /api/v1/mantenimientos
```

Body:

```json
{
  "vehiculo_id": 1,
  "motivo": "Cambio de aceite",
  "fecha_inicio": "2026-07-08",
  "fecha_fin": "2026-07-10"
}
```

Este endpoint requiere rol `admin`.

---

## Pruebas con Postman

La colección de Postman del módulo se encuentra en:

```text
postman/
```

La colección incluye pruebas para:

- Health Check
- Registro y login
- CRUD de rutas programadas
- CRUD de horarios de ruta
- CRUD de mantenimientos
- Prueba de rutas protegidas con JWT
- Prueba de permisos por rol admin/conductor

---

## Pruebas unitarias

Ejecutar todos los tests del proyecto:

```bash
go test ./...
```

Ejecutar revisión estática:

```bash
go vet ./...
```

Ejecutar cobertura del service del módulo Ruta Programada:

```bash
go test ./internal/service/rutaProgramada -cover
```

Cobertura obtenida:

```text
56.0% of statements
```

---

## Tests del módulo

Para este módulo se realizaron pruebas en las capas principales de la entidad `RutaProgramada`:

```text
Handler → Service → Storage/Repository
```

En el service se usaron mocks con `testify` para simular el repository sin depender directamente de la base de datos real.

Casos probados:

- Crear una ruta programada válida.
- Rechazar una ruta con costo negativo.
- Obtener una ruta existente.
- Manejar una ruta inexistente.
- Actualizar una ruta programada.
- Borrar una ruta programada.
- Validar horarios y mantenimientos.
- Probar acceso con token JWT.
- Probar permisos por rol.

---

## CI/CD

El proyecto utiliza GitHub Actions para validar automáticamente:

```text
build → vet → test
```

El pipeline debe estar en verde antes de integrar cambios a `main`.

---

## Docker

El proyecto cuenta con:

- Dockerfile multi-stage.
- docker-compose.yml.
- Servicio de API.
- Servicio de PostgreSQL.

Comando principal:

```bash
docker compose up
```

---

## Estado del módulo

El módulo Ruta Programada cuenta con:

- Arquitectura por capas.
- Interfaces e inyección de dependencias.
- Persistencia con GORM.
- Relaciones entre `RutaProgramada` y `HorarioRuta`.
- Migraciones automáticas.
- Endpoints protegidos con JWT.
- Control de roles para mantenimientos.
- Tests unitarios.
- Mocks con Testify.
- Cobertura del service mayor al 50%.
- Pruebas en Postman.