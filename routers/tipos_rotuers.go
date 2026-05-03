package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterTiposRouters(r *mux.Router) {
	r.HandleFunc("/tipos", controllers.GetTipos).Methods("GET")
	r.HandleFunc("/tipos/{id}", controllers.GetTiposByID).Methods("GET")
	r.HandleFunc("/tipos", controllers.CreateTipos).Methods("POST")
	r.HandleFunc("/tipos/{id}", controllers.UpdateTipos).Methods("PUT")
	r.HandleFunc("/tipos/{id}", controllers.DeleteTipos).Methods("DELETE")

}
