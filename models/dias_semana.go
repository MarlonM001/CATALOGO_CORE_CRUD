package models

type DiaSemana struct {
	ID_DIAS_SEMANA  int    `json:"id_dias_semana"`
	NOMBRE          string `json:"nombre"`
	NOMBRE_COMPLETO string `json:"nombre_completo"`
	DESCRIPCION     string `json:"descripcion"`
	ORDEN           int    `json:"orden"`
	ACTIVO          bool   `json:"activo"`
}
