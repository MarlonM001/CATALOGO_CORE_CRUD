package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterRolesUsuarioRouters(r *mux.Router) {
	r.HandleFunc("/roles_usuario", controllers.GetExperiencia).Methods("GET")
	r.HandleFunc("/roles_usuario/{id}", controllers.GetExperienciaByID).Methods("GET")
	r.HandleFunc("/roles_usuario", controllers.CreateExperiencia).Methods("POST")
	r.HandleFunc("/roles_usuario/{id}", controllers.UpdateExperiencia).Methods("PUT")
	r.HandleFunc("/roles_usuario/{id}", controllers.DeleteExperiencia).Methods("DELETE")

}
