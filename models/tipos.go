package models

type Tipos struct {
	ID_TIPO     int    `json:"id_tipo"`
	CONTEXTO    string `json:"contexto"`
	NOMBRE      string `json:"nombre"`
	DESCRIPCION string `json:"descripcion"`
	ACTIVO      bool   `json:"activo"`
}
