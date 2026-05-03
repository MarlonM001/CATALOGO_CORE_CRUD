package models

type Estados struct {
	ID_ESTADO   int    `json:"id_estado"`
	CONTEXTO    string `json:"contexto"`
	NOMBRE      string `json:"nombre"`
	DESCRIPCION string `json:"descripcion"`
	ACTIVO      bool   `json:"activo"`
}
