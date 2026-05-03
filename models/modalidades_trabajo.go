package models

type ModalidadesTrabajo struct {
	ID_MODALIDAD_TRABAJO int    `json:"id_modalidad_trabajo"`
	NOMBRE               string `json:"nombre"`
	DESCRIPCION          string `json:"descripcion"`
	ACTIVO               bool   `json:"activo"`
}
