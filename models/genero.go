package models

type Genero struct {
	ID_GENERO int    `json:"id_genero"`
	NOMBRE    string `json:"nombre"`
	ACTIVO    bool   `json:"activo"`
}
