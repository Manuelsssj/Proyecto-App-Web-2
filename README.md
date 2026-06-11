# Proyecto-App-Web-2

# Descripción 

RideULEAM es un proyecto orientado a resolver los problemas de transporte que enfrentan diariamente los estudiantes universitarios de Manta. La propuesta consiste en desarrollar una API REST que permita coordinar viajes compartidos entre estudiantes que viven en sectores cercanos y tienen horarios similares de entrada y salida hacia la universidad.

El sistema permitirá conectar conductores y pasajeros mediante un mecanismo de matching inteligente basado en sector, horario y disponibilidad de asientos, facilitando una forma más organizada, económica y segura de compartir transporte.

Este proyecto está siendo desarrollado como parte de la materia Aplicaciones Web II en la Universidad Laica Eloy Alfaro de Manabí.

# Problema identificado

Muchos estudiantes universitarios de la ULEAM se trasladan diariamente desde distintos sectores de Manta utilizando taxi, buseta o transporte informal. Aunque varios estudiantes comparten rutas y horarios similares, normalmente viajan por separado debido a la falta de una herramienta organizada para coordinar transporte compartido.

Actualmente la coordinación se realiza mediante grupos de WhatsApp o conversaciones informales, lo que genera desorganización, pérdida de tiempo y mayores gastos de transporte. Además, no existe un sistema que permita conocer disponibilidad de asientos, coincidencias de rutas o niveles de confianza entre usuarios.

# Solución propuesta

RideULEAM propone una API backend que permitirá conectar estudiantes conductores con pasajeros que tengan rutas y horarios compatibles.

El sistema permitirá:

-publicar viajes
-buscar viajes compatibles
-solicitar reservas
-gestionar disponibilidad de asientos
-registrar historial y calificaciones entre usuarios

La solución busca reducir gastos de transporte, mejorar la coordinación entre estudiantes y aprovechar espacios disponibles en vehículos particulares.

# Regla de negocio no-CRUD

La principal lógica del sistema es el matching inteligente entre pasajeros y conductores.

La API no solo almacenará información de usuarios y viajes, sino que analizará:

-sector de origen
-horario deseado
-compatibilidad de rutas
-disponibilidad de asientos
-estado del viaje

Con esta información, el sistema podrá recomendar automáticamente viajes compatibles para cada estudiante.

Por esta razón, el proyecto no se limita a un CRUD tradicional, ya que incorpora lógica de recomendación y compatibilidad entre usuarios.

# Tecnologías

Go
API REST
GitHub

# Estructura del proyecto

rideuleam/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── services/
│   ├── repositories/
│   └── storage/
│
├── go.mod
├── README.md
└── .gitignore

# Descripción de carpetas

|       Carpeta                            Función                        |
|                       |                                                 |
| cmd/api               | Punto de entrada principal de la API            |
| internal/handlers     | Manejo de endpoints y requests HTTP             |
| internal/models       | Structs y modelos del sistema                   |
| internal/services     | Lógica de negocio                               |
| internal/repositories | Acceso y manipulación de datos                  |
| internal/storage      | Configuración de almacenamiento y base de datos |

#Módulos del sistema

| Integrante   |           Módulo          |                Responsabilidad                    |
|              |                           |                                                   |
| Integrante A | Viajes                    | Gestión de viajes publicados por conductores      |
| Integrante B | Reservas y Matching       | Búsqueda de viajes compatibles y solicitudes      |
| Integrante C | Usuarios y Calificaciones | Gestión de usuarios, roles y sistema de confianza |

# Structs principales

El sistema estará compuesto por los siguientes structs principales:
-User
-Viaje
-Reserva
-Calificacion

Estos structs permitirán modelar usuarios, viajes compartidos, reservas y el sistema de reputación entre conductores y pasajeros.

# Endpoints principales

| Método |        Ruta            |      Descripción          |
|        |                        |                           |
| POST   | /api/v1/viajes         | Crear un viaje            |
| GET    | /api/v1/viajes         | Obtener viajes            |
| POST   | /api/v1/reservas       | Solicitar reserva         |
| GET    | /api/v1/matching       | Buscar viajes compatibles |
| POST   | /api/v1/calificaciones | Registrar calificación    |

# Estado actual

Proyecto en fase de descubrimiento y diseño correspondiente al Hito 1 de Aplicaciones Web II.

Actualmente se encuentra en desarrollo:
-levantamiento de información
-entrevistas
-diseño de structs
-diseño de endpoints
-organización inicial del repositorio

