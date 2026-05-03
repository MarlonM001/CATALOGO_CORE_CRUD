package models

type ModalidadesSuscripcion struct {
	ID_MODALIDAD_SUSCRIPCION int    `json:"id_modalidad_suscripcion"`
	NOMBRE                   string `json:"nombre"`
	DESCRIPCION              string `json:"descripcion"`
	ACTIVO                   bool   `json:"activo"`
}
