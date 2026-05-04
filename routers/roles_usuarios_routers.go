package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterRolesUsuarioRouters(r *mux.Router) {
	r.HandleFunc("/roles_usuarios", controllers.GetRolesUsuario).Methods("GET")
	r.HandleFunc("/roles_usuarios/{id}", controllers.GetRolesUsuarioByID).Methods("GET")
	r.HandleFunc("/roles_usuarios", controllers.CreateRolesUsuario).Methods("POST")
	r.HandleFunc("/roles_usuarios/{id}", controllers.UpdateRolesUsuario).Methods("PUT")
	r.HandleFunc("/roles_usuarios/{id}", controllers.DeleteRolesUsuario).Methods("DELETE")

}
