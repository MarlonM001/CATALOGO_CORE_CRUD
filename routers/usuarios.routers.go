package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterUsuarios(r *mux.Router) {
	// ✅ GET ALL apunta a GetUsuario, no a GetUsuariosByID
	r.HandleFunc("/usuarios", controllers.GetUsuario).Methods("GET")
	r.HandleFunc("/usuarios/{id}", controllers.GetUsuariosByID).Methods("GET")
	r.HandleFunc("/usuarios", controllers.CreateUsuarios).Methods("POST")
	r.HandleFunc("/usuarios/{id}", controllers.UpdateUsuarios).Methods("PUT")
	r.HandleFunc("/usuarios/{id}", controllers.DeleteUsuarios).Methods("DELETE")
}
