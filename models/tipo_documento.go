package models

type TiposDocumentos struct {
	ID_TIPO_DOCUMENTOS int    `json:"id_tipo_documentos"`
	NOMBRE             string `json:"nombre"`
	DESCRIPCION        string `json:"descripcion"`
	ACTIVO             bool   `json:"activo"`
}
