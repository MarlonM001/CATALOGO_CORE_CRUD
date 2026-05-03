package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterTurnoRouters(r *mux.Router) {
	r.HandleFunc("/turnos", controllers.GetTurno).Methods("GET")
	r.HandleFunc("/turnos/{id}", controllers.GetTurnoByID).Methods("GET")
	r.HandleFunc("/turnos", controllers.CreateTurno).Methods("POST")
	r.HandleFunc("/turnos/{id}", controllers.UpdateTurno).Methods("PUT")
	r.HandleFunc("/turnos/{id}", controllers.DeleteTurno).Methods("DELETE")

}
