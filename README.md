# CATALOGO_CORE_CRUD

Módulo backend en Go que centraliza los datos de referencia y la gestión de usuarios del sistema. Administra estados, tipos, configuraciones, identidad, roles, credenciales y sesiones, siendo la base de todos los procesos del sistema.

---

## Nombre del Schema

**Schema:** `catalogo`

Este schema organiza todas las tablas relacionadas con datos de referencia y gestión de usuarios dentro de la base de datos.

---

## Tecnologías utilizadas

- **Go (v1.21+)** — Lenguaje principal  
- **Gorilla Mux (v1.8.1)** — Enrutador HTTP  
- **PostgreSQL** — Base de datos relacional  
- **lib/pq (v1.12.3)** — Driver de PostgreSQL para Go  
- **Go Modules** — Gestión de dependencias  
- **Visual Studio Code** — Entorno de desarrollo  

---

## Estructura del proyecto


CATALOGO_CORE_CRUD/
├── config/
│ └── db.go
├── routers/
│ ├── generos.go
│ ├── experiencia.go
│ ├── modalidades_trabajo.go
│ ├── turno.go
│ ├── dias_semanas.go
│ ├── estados.go
│ ├── tipos.go
│ ├── modalidades_suscripcion.go
│ ├── roles_usuario.go
│ ├── tipo_documentos.go
│ ├── usuarios.go
│ ├── documentos_identidad.go
│ ├── credenciales.go
│ └── sesion.go
├── main.go
├── go.mod
├── go.sum
└── README.md


---

##  Variables de entorno

Crea un archivo `.env` en la raíz del proyecto (**no subir al repositorio**):

```env
# Base de datos
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_contraseña
DB_NAME=catalogo_core
DB_SCHEMA=catalogo
DB_SSLMODE=disable

# Servidor
APP_PORT=8082
APP_ENV=development



▶️ Ejecución del proyecto
🔧 Requisitos previos
Go 1.21+
PostgreSQL
Git
Visual Studio Code
📌 Pasos
1. Clonar repositorio
git clone https://github.com/tu-usuario/CATALOGO_CORE_CRUD.git
cd CATALOGO_CORE_CRUD
2. Instalar dependencias
go mod tidy
3. Configurar entorno
cp .env.example .env
4. Ejecutar proyecto
go run main.go
5. Compilar
go build -o bin/catalogo_core_crud .
./bin/catalogo_core_crud
6. Tests
go test ./...

Servidor disponible en:
http://localhost:8082

CORS

Configurado con origen abierto (*) y métodos:

GET
POST
PUT
DELETE
OPTIONS



📄 Licencia

JOBSY