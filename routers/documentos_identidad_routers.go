package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterDocumentosIdentidadRouters(r *mux.Router) {
	r.HandleFunc("/documentos_identidad", controllers.GetDocumentosIdentidad).Methods("GET")
	r.HandleFunc("/documentos_identidad/{id}", controllers.GetDocumentosIdentidadByID).Methods("GET")
	r.HandleFunc("/documentos_identidad", controllers.CreateDocumentosIdentidad).Methods("POST")
	r.HandleFunc("/documentos_identidad/{id}", controllers.UpdateDocumentosIdentidad).Methods("PUT")
	r.HandleFunc("/documentos_identidad/{id}", controllers.DeleteDocumentosIdentidad).Methods("DELETE")

}
