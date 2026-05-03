package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterDiasSemanasRouters(r *mux.Router) {
	r.HandleFunc("/dias_semana", controllers.GetDiasSemanas).Methods("GET")
	r.HandleFunc("/dias_semana/{id}", controllers.GetDiasSemanaByID).Methods("GET")
	r.HandleFunc("/dias_semana", controllers.CreateDiasSemana).Methods("POST")
	r.HandleFunc("/dias_semana/{id}", controllers.UpdateDiasSemana).Methods("PUT")
	r.HandleFunc("/dias_semana/{id}", controllers.DeleteDiasSemana).Methods("DELETE")

}
