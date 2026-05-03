package models

type Turno struct {
	ID_TURNOS   int    `json:"id_turno"`
	NOMBRE      string `json:"nombre"`
	DESCRIPCION string `json:"descripcion"`
	ACTIVO      bool   `json:"activo"`
}
