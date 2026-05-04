package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterTipoDocumentos(r *mux.Router) {
	r.HandleFunc("/tipo_documento", controllers.GetTipoDocumentos).Methods("GET")
	r.HandleFunc("/tipo_documento/{id}", controllers.GetTiposDocumentosByID).Methods("GET")
	r.HandleFunc("/tipo_documento", controllers.CreateTiposDocumentos).Methods("POST")
	r.HandleFunc("/tipo_documento/{id}", controllers.UpdateTiposDocumentos).Methods("PUT")
	r.HandleFunc("/tipo_documento/{id}", controllers.DeleteTiposDocumentos).Methods("DELETE")

}
