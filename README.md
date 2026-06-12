# RideULEAM

RideULEAM es una API REST desarrollada en Go para organizar el transporte compartido entre estudiantes universitarios.

## Módulos actuales

El proyecto trabaja con tres módulos principales:

1. Rutas
2. Suscripciones
3. Estado

## Módulo Estado

El módulo Estado permite administrar la disponibilidad de una ruta, indicando si se encuentra activa, inactiva o en mantenimiento.

### Endpoints del módulo Estado

POST   /api/v1/estados/estado  
GET    /api/v1/estados/estados  
GET    /api/v1/estados/estado/{id}  
PUT    /api/v1/estados/estado/{id}  
DELETE /api/v1/estados/estado/{id}

## Tecnologías utilizadas

- Go
- Chi Router
- API REST
- Almacenamiento en memoria
- Postman para pruebas

## Ejecución

```bash
go run cmd/api/main.go

servidor
http://localhost:8080