package models

type TiposDocumentos struct {
	ID_TIPO_DOCUMENTO int    `json:"id_tipo_documento"`
	NOMBRE            string `json:"nombre"`
	DESCRIPCION       string `json:"descripcion"`
	ACTIVO            bool   `json:"activo"`
}
