package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterSesionRouters(r *mux.Router) {
	r.HandleFunc("/sesion", controllers.GetSesion).Methods("GET")
	r.HandleFunc("/sesion/{id}", controllers.GetSesionByID).Methods("GET")
	r.HandleFunc("/sesion", controllers.CreateSesion).Methods("POST")
	r.HandleFunc("/sesion/{id}", controllers.UpdateSesion).Methods("PUT")
	r.HandleFunc("/sesion/{id}", controllers.DeleteSesion).Methods("DELETE")

}
