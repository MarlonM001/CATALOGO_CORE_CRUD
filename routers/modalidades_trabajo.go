package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterModalidadesTrabajoRouters(r *mux.Router) {

	r.HandleFunc("/ModalidadesTrabajo", controllers.GetModalidadesTrabajo).Methods("GET")
	r.HandleFunc("/ModalidadesTrabajo/{id}", controllers.GetModalidadesTrabajoByID).Methods("GET")
	r.HandleFunc("/ModalidadesTrabajo", controllers.CreateModalidadesTrabajo).Methods("POST")
	r.HandleFunc("/ModalidadesTrabajo/{id}", controllers.UpdateModalidadesTrabajo).Methods("PUT")
	r.HandleFunc("/ModalidadesTrabajo/{id}", controllers.DeleteModalidadesTrabajo).Methods("DELETE")

}
