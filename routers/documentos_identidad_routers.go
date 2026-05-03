package routers

import (
	"CATALOGO_CORE_CRUD/controllers"

	"github.com/gorilla/mux"
)

func RegisterDocumentosIdentidadRouters(r *mux.Router) {
	r.HandleFunc("/documentos_identidad", controllers.GetDocumentoIdentidadByID).Methods("GET")
	r.HandleFunc("/documentos_identidad/{id}", controllers.GetDocumentoIdentidadByID).Methods("GET")
	r.HandleFunc("/documentos_identidad", controllers.CreateDocumentoIdentidad).Methods("POST")
	r.HandleFunc("/documentos_identidad/{id}", controllers.UpdateDocumentoIdentidad).Methods("PUT")
	r.HandleFunc("/documentos_identidad/{id}", controllers.DeleteDocumentoIdentidad).Methods("DELETE")

}
