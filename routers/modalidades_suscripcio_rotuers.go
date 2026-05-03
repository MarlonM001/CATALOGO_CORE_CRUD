package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterModalidadesSuscripcionRouters(r *mux.Router) {
	r.HandleFunc("/modalidades-suscripcion", controllers.GetModalidadesSuscripcion).Methods("GET")
	r.HandleFunc("/modalidades-suscripcion/{id}", controllers.GetModalidadesSuscripcionByID).Methods("GET")
	r.HandleFunc("/modalidades-suscripcion", controllers.CreateModalidadesSuscripcion).Methods("POST")
	r.HandleFunc("/modalidades-suscripcion/{id}", controllers.UpdateModalidadesSuscripcion).Methods("PUT")
	r.HandleFunc("/modalidades-suscripcion/{id}", controllers.DeleteModalidadesSuscripcion).Methods("DELETE")

}
