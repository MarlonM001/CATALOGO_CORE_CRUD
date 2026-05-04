package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterRolesUsuarioRouters(r *mux.Router) {
	r.HandleFunc("/roles_usuario", controllers.GetRolesUsuario).Methods("GET")
	r.HandleFunc("/roles_usuario/{id}", controllers.GetRolesUsuarioByID).Methods("GET")
	r.HandleFunc("/roles_usuario", controllers.CreateRolesUsuario).Methods("POST")
	r.HandleFunc("/roles_usuario/{id}", controllers.UpdateRolesUsuario).Methods("PUT")
	r.HandleFunc("/roles_usuario/{id}", controllers.DeleteRolesUsuario).Methods("DELETE")

}
