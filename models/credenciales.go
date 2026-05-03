package models

import "time"

type Credenciales struct {
	ID_CREDENCIALES   int        `json:"id_credenciales"`
	ID_USUARIOS       int        `json:"id_usuarios"`
	PASSWORD_HASH     string     `json:"password_hash"`
	SALT              string     `json:"salt"`
	DOBLE_FACTOR      bool       `json:"doble_factor"`
	TOKEN_RESET       *string    `json:"token_reset"`
	TOKEN_EXPIRA_EN   *time.Time `json:"token_expira_en"`
	ULTIMO_LOGIN      *time.Time `json:"ultimo_login"`
	INTENTOS_FALLIDOS int        `json:"intentos_fallidos"`
	BLOQUEADO         bool       `json:"bloqueado"`
	ACTIVO            bool       `json:"activo"`
}
