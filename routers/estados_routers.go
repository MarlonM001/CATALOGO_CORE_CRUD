package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterEstadosRouters(r *mux.Router) {
	r.HandleFunc("/estados", controllers.GetEstados).Methods("GET")
	r.HandleFunc("/estados/{id}", controllers.GetEstadosByID).Methods("GET")
	r.HandleFunc("/estados", controllers.CreateEstado).Methods("POST")
	r.HandleFunc("/estados/{id}", controllers.UpdateEstado).Methods("PUT")
	r.HandleFunc("/estados/{id}", controllers.DeleteEstado).Methods("DELETE")

}
