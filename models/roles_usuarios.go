package models

type RolesUsuario struct {
	ID_ROLES_USUARIO int    `json:"id_roles_usuario"`
	NOMBRE           string `json:"nombre"`
	DESCRIPCION      string `json:"descripcion"`
	ACTIVO           bool   `json:"activo"`
}
