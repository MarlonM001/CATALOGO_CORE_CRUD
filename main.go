package main

import (
	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//middleware CORS

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})

}

func main() {
	config.ConnectDB() // Linea para conectar la DB

	r := mux.NewRouter()

	r.Use(enableCORS)

	routers.RegisterGenerosRouters(r)

	routers.RegisterExperienciaRouters(r)

	routers.RegisterModalidadesTrabajoRouters(r)

	routers.RegisterTurnoRouters(r)

	routers.RegisterDiasSemanasRouters(r)

	routers.RegisterEstadosRouters(r)

	routers.RegisterTiposRouters(r)

	routers.RegisterModalidadesSuscripcionRouters(r)

	routers.RegisterRolesUsuarioRouters(r)

	routers.RegisterTipoDocumentos(r)

	routers.RegisterUsuarios(r)

	log.Println("servidor corriendo en el puerto 8082")

	http.ListenAndServe(":8082", enableCORS(r))

}
