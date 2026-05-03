package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterGenerosRouters(r *mux.Router) {
	r.HandleFunc("/generos", controllers.GetGeneros).Methods("GET")
	r.HandleFunc("/generos/{id}", controllers.GetGenerosByID).Methods("GET")
	r.HandleFunc("/generos", controllers.CreateGeneros).Methods("POST")
	r.HandleFunc("/generos/{id}", controllers.UpdateGeneros).Methods("PUT")
	r.HandleFunc("/generos/{id}", controllers.DeleteGeneros).Methods("DELETE")

}
