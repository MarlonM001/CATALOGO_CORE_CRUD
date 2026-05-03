package models

type NivelesExperiencia struct {
	ID_EXPERIENCIA int    `json:"id_nivel_experiencia"`
	NOMBRE         string `json:"nombre"`
	DESCRIPCION    string `json:"descripcion"`
	ACTIVO         bool   `json:"activo"`
}
