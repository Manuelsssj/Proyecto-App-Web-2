# Diagrama de arquitectura

RideULEAM usa arquitectura por capas. Cada request entra por el router de Chi, pasa por middleware y llega a un handler. El handler valida HTTP/JSON, delega reglas de negocio al service y el service usa interfaces de repository/storage. La implementacion concreta de persistencia se inyecta desde `cmd/rideUleam/main.go`.

```text
Cliente HTTP / Postman
        |
        v
Chi Router + Middleware
        |
        v
Handler
        |
        v
Service
        |
        v
Repository / Storage interface
        |
        v
GORM + PostgreSQL / SQLite
```

## Capas principales

- `cmd/rideUleam/main.go`: compone dependencias, configura router y registra rutas.
- `internal/handlers`: recibe requests, decodifica JSON y responde HTTP.
- `internal/service`: valida reglas de negocio y coordina operaciones.
- `internal/storage`: define interfaces e implementaciones de persistencia.
- `internal/models`: entidades con tags JSON y GORM.
- `internal/middleware`: JWT, roles y CORS.

## Flujo de autenticacion

```text
POST /auth/login -> Auth handler -> AuthService -> UsuarioRepository -> JWT firmado
```

El JWT incluye `uid` y `rol`. Las rutas protegidas primero validan el token y luego pueden exigir roles con `RequiereRol`.

## Relaciones GORM

- `Usuario` tiene relacion logica con `Vehiculo` mediante `ConductorID`.
- `ViajeInmediato` pertenece a un conductor y tiene muchas solicitudes y participantes.
- `SolicitudViaje` pertenece a un viaje y a un pasajero.
- `ParticipanteViaje` pertenece a un viaje y a un usuario.
