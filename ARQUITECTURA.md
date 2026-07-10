# Diagrama de arquitectura — RideUleam API

```text
Cliente / Postman
       |
       v
API REST - Chi Router
       |
       v
Middleware
- CORS
- JWT Auth
       |
       v
Handlers
- Auth
- Suscripciones
- Planes
- Historial
       |
       v
Services
- Validaciones
- Reglas de negocio
       |
       v
Repositories (Storage)
- Interfaces
- Implementación GORM
       |
       v
GORM
       |
       v
PostgreSQL
```

## Flujo de una petición

1. El cliente (Postman o cualquier consumidor de la API) envía una solicitud HTTP.
2. El router Chi identifica el endpoint correspondiente.
3. El middleware procesa la solicitud, aplicando CORS y validando el token JWT cuando la ruta está protegida.
4. El Handler recibe la petición, procesa el JSON y delega la lógica al Service.
5. El Service ejecuta las validaciones y reglas de negocio.
6. El Repository accede a PostgreSQL utilizando GORM.
7. PostgreSQL devuelve la información solicitada.
8. El Handler construye la respuesta en formato JSON y la envía al cliente.