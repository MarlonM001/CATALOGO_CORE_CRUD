## Estrutura del proyecto

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

##  Módulos registrados

| Módulo                     | Función registrada                 |
|--------------------------|----------------------------------|
| Géneros                  | RegisterGenerosRouters           |
| Experiencia              | RegisterExperienciaRouters       |
| Modalidades de trabajo   | RegisterModalidadesTrabajoRouters|
| Turno                    | RegisterTurnoRouters             |
| Días de la semana        | RegisterDiasSemanasRouters       |
| Estados                  | RegisterEstadosRouters           |
| Tipos                    | RegisterTiposRouters             |
| Modalidades suscripción  | RegisterModalidadesSuscripcionRouters |
| Roles de usuario         | RegisterRolesUsuarioRouters      |
| Tipo de documentos       | RegisterTipoDocumentos           |
| Usuarios                 | RegisterUsuarios                 |
| Documentos de identidad  | RegisterDocumentosIdentidadRouters |
| Credenciales             | RegisterCredencialesRouters      |
| Sesión                   | RegisterSesionRouters            |

---

##  Variables de entorno

Crea un archivo `.env` en la raíz del proyecto. **No subas este archivo al repositorio.**

```env
# Base de datos
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_contraseña
DB_NAME=catalogo_core
DB_SSLMODE=disable

# Servidor
APP_PORT=8082
APP_ENV=development

Agrega .env a tu .gitignore y sube solo el archivo .env.example como referencia.

 Ejecución 

🔧 Requisitos previos
Go 1.21+
PostgreSQL
Visual Studio Code + extensión Go
Git

 Pasos

1. Clonar el repositorio
git clone https://github.com/tu-usuario/CATALOGO_CORE_CRUD.git
cd CATALOGO_CORE_CRUD
2. Instalar dependencias
go mod tidy
3. Configurar variables de entorno
cp .env.example .env
# Edita .env con tus credenciales
4. Ejecutar en desarrollo
go run main.go
5. Compilar y ejecutar
go build -o bin/catalogo_core_crud .
./bin/catalogo_core_crud
6. Ejecutar pruebas
go test ./...

El servidor corre en:
👉 http://localhost:8082

CORS

El middleware CORS está habilitado con origen abierto (*) y soporta los métodos:

GET
POST
PUT
DELETE
OPTIONS



Licencia

JOBSY