package models

type DocumentosIdentidad struct {
	ID_DOCUMENTOS_IDENTIDAD int    `json:"id_documentos_identidad"`
	ID_USUARIOS             int    `json:"id_usuarios"`
	ID_TIPO_DOCUMENTO       int    `json:"id_tipo_documento"`
	NUMERO_DOCUMENTO        string `json:"numero_documento"`
	FECHA_EXPEDICION        string `json:"fecha_expedicion"`
	LUGAR_EXPEDICION        string `json:"lugar_expedicion"`
	ACTIVO                  bool   `json:"activo"`
}
