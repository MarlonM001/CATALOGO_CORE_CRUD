package models

import "time"

type Sesion struct {
	ID_SESIONES    int       `json:"id_sesiones"`
	ID_USUARIOS    int       `json:"id_usuarios"`
	TOKEN          string    `json:"token"`
	DISPOSITIVO    *string   `json:"dispositivo"`
	IP             *string   `json:"ip"`
	ACTIVA         bool      `json:"activa"`
	EXPIRA_EN      time.Time `json:"expira_en"`
	CREADO_EN      time.Time `json:"creado_en"`
	ACTUALIZADO_EN time.Time `json:"actualizado_en"`
}
