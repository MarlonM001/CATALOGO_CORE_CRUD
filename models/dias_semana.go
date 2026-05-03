package models

type DiaSemana struct {
	ID_DIA_SEMANA   int    `json:"id_dia_semana"`
	NOMBRE          string `json:"nombre"`
	NOMBRE_COMPLETO string `json:"nombre_completo"`
	DESCRIPCION     string `json:"descripcion"`
	ORDEN           int    `json:"orden"`
	ACTIVO          bool   `json:"activo"`
}
