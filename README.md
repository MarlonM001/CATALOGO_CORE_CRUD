CATALOGO_CORE_CRUD

Módulo backend desarrollado en Go que centraliza los datos de referencia y la gestión de usuarios del sistema. Administra valores como estados, tipos y configuraciones para mantener consistencia, y controla identidad, roles, credenciales y sesiones, siendo la base de todos los procesos del sistema.

Tecnologías Utilizadas

TecnologíaVersiónDescripciónGo1.21+Lenguaje principal del proyectoGo Modulesgo.modGestión de dependenciasgorilla/muxv1.8.1Router HTTPlib/pqv1.12.3Driver PostgreSQL para GoPostgreSQL—Base de datos relacionalVisual Studio Code—IDE recomendado

Estructura del Proyecto



CATALOGO_CORE_CRUD/
├── config/
│   └── db.go                  
├── routers/
│   ├── generos.go
│   ├── experiencia.go
│   ├── modalidades_trabajo.go
│   ├── turno.go
│   ├── dias_semanas.go
│   ├── estados.go
│   ├── tipos.go
│   ├── modalidades_suscripcion.go
│   ├── roles_usuario.go
│   ├── tipo_documentos.go
│   ├── usuarios.go
│   ├── documentos_identidad.go
│   ├── credenciales.go
│   └── sesion.go
├── main.go                    
├── go.mod                     
├── go.sum                     
└── README.md


Variables de Entorno

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_contraseña
DB_NAME=catalogo_core
DB_SSLMODE=disable


APP_PORT=8082
APP_ENV=development        




Ejecución del Proyecto

Prerrequisitos
Asegúrate de tener instalado:

Go 1.21+
PostgreSQL
Visual Studio Code con la extensión oficial de Go
Git

1. Clonar el repositorio
bashgit clone https://github.com/tu-usuario/CATALOGO_CORE_CRUD.git
cd CATALOGO_CORE_CRUD
2. Instalar dependencias
bashgo mod tidy
3. Configurar variables de entorno
bashcp .env.example .env
# Edita el archivo .env con tus credenciales reales
4. Ejecutar en modo desarrollo
bashgo run main.go
5. Compilar y ejecutar el binario
bash# Compilar
go build -o bin/catalogo_core_crud .

# Ejecutar
./bin/catalogo_core_crud
El servidor quedará corriendo en:
http://localhost:8082
6. Ejecutar pruebas
bashgo test ./...

CORS

El servicio tiene habilitado el middleware CORS con acceso abierto (*). Permite los métodos GET, POST, PUT, DELETE y OPTIONS, adecuado para entornos de desarrollo. Ajusta la política de origen en main.go antes de pasar a producción.

Extensiones recomendadas para VS Code
json{
  "recommendations": [
    "golang.go",
    "EditorConfig.EditorConfig",
    "eamodio.gitlens"
  ]
}



Licencia


Este proyecto está bajo la licencia JOBSY