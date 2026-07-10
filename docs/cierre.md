# Documento de cierre

## Que aprendimos

Durante el desarrollo de RideULEAM aprendimos a estructurar una API REST en Go usando capas separadas: handlers para HTTP, services para reglas de negocio y repositories para persistencia. Tambien reforzamos el uso de GORM, migraciones automaticas, Docker Compose, PostgreSQL, JWT y pruebas unitarias con mocks.

## Que hariamos distinto

Planificariamos desde el inicio los contratos entre modulos para evitar ajustes tardios de integracion. Tambien definiriamos antes los roles, permisos y relaciones entre entidades, porque son decisiones que afectan modelos, handlers, pruebas y documentacion.

## Proximos pasos

- Mejorar la cobertura global de pruebas.
- Agregar filtros por origen, destino, conductor y estado del viaje.
- Crear endpoints que devuelvan entidades relacionadas usando `Preload`.
- Fortalecer roles y permisos por modulo.
- Publicar una version desplegada para pruebas reales con estudiantes.
