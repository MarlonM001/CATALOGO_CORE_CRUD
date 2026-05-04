package models

type Usuario struct {
	ID_USUARIO        int    `json:"id_usuario"`
	NOMBRE            string `json:"nombre"`
	APELLIDO          string `json:"apellido"`
	CORREO            string `json:"correo"`
	TELEFONO          string `json:"telefono"`
	CIUDAD            string `json:"ciudad"`
	FECHA_NACIMIENTO  string `json:"fecha_nacimiento"`
	ID_GENERO         int    `json:"id_genero"`
	NACIONALIDAD      string `json:"nacionalidad"`
	FOTO_URL          string `json:"foto_url"`
	ID_ROLES_USUARIOS int    `json:"id_roles_usarios"`
	DISPONIBLE        bool   `json:"disponible"`
	PERFIL_PUBLICO    bool   `json:"perfil_publico"`
	ACTIVO            bool   `json:"activo"`
}
