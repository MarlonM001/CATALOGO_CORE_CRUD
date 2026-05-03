package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterTipoDocumentos(r *mux.Router) {
	r.HandleFunc("/tipo_documentos", controllers.GetExperiencia).Methods("GET")
	r.HandleFunc("/tipo_documentos/{id}", controllers.GetExperienciaByID).Methods("GET")
	r.HandleFunc("/tipo_documentos", controllers.CreateExperiencia).Methods("POST")
	r.HandleFunc("/tipo_documentos/{id}", controllers.UpdateExperiencia).Methods("PUT")
	r.HandleFunc("/tipo_documentos/{id}", controllers.DeleteExperiencia).Methods("DELETE")

}
