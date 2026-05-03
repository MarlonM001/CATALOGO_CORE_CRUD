package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterExperienciaRouters(r *mux.Router) {
	r.HandleFunc("/niveles_experiencia", controllers.GetExperiencia).Methods("GET")
	r.HandleFunc("/niveles_experiencia/{id}", controllers.GetExperienciaByID).Methods("GET")
	r.HandleFunc("/niveles_experiencia", controllers.CreateExperiencia).Methods("POST")
	r.HandleFunc("/niveles_experiencia/{id}", controllers.UpdateExperiencia).Methods("PUT")
	r.HandleFunc("/niveles_experiencia/{id}", controllers.DeleteExperiencia).Methods("DELETE")

}
