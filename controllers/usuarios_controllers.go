package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_usuarios, nombre, apellido, correo, telefono, ciudad, fecha_nacimiento, id_genero,
       nacionalidad, foto_url, id_roles_usuarios, disponible, perfil_publico, activo
FROM core.usuarios`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Usuario

	for rows.Next() {
		var c models.Usuario

		err := rows.Scan(
			&c.ID_USUARIO,
			&c.NOMBRE,
			&c.APELLIDO,
			&c.CORREO,
			&c.TELEFONO,
			&c.CIUDAD,
			&c.FECHA_NACIMIENTO,
			&c.ID_GENERO,
			&c.NACIONALIDAD,
			&c.FOTO_URL,
			&c.ID_ROLES_USUARIOS,
			&c.DISPONIBLE,
			&c.PERFIL_PUBLICO,
			&c.ACTIVO,
		)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

func GetUsuariosByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Usuario

	err = config.DB.QueryRow(`
		SELECT id_usuarios, nombre, apellido, correo, telefono, ciudad, fecha_nacimiento, id_genero,
		       nacionalidad, foto_url, id_roles_usuarios, disponibilidad, perfil_publico, activo
		FROM core.usuarios
		WHERE id_usuarios = $1`, id).
		Scan(
			&c.ID_USUARIO,
			&c.NOMBRE,
			&c.APELLIDO,
			&c.CORREO,
			&c.TELEFONO,
			&c.CIUDAD,
			&c.FECHA_NACIMIENTO,
			&c.ID_GENERO,
			&c.NACIONALIDAD,
			&c.FOTO_URL,
			&c.ID_ROLES_USUARIOS,
			&c.DISPONIBLE,
			&c.PERFIL_PUBLICO,
			&c.ACTIVO,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		} else {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
		}
		return
	}

	respondJSON(w, 200, c)
}

func CreateUsuarios(w http.ResponseWriter, r *http.Request) {
	var c models.Usuario

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(`
		INSERT INTO core.usuarios 
		(nombre, apellido, correo, telefono, ciudad, fecha_nacimiento, id_genero,
		 nacionalidad, foto_url, id_roles_usuarios, disponibilidad, perfil_publico, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
		RETURNING id_usuarios`,
		c.NOMBRE,
		c.APELLIDO,
		c.CORREO,
		c.TELEFONO,
		c.CIUDAD,
		c.FECHA_NACIMIENTO,
		c.ID_GENERO,
		c.NACIONALIDAD,
		c.FOTO_URL,
		c.ID_ROLES_USUARIOS,
		c.DISPONIBLE,
		c.PERFIL_PUBLICO,
		c.ACTIVO,
	).Scan(&c.ID_USUARIO)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

func UpdateUsuarios(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Usuario

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(`
		UPDATE core.usuarios SET 
			nombre=$1,
			apellido=$2,
			correo=$3,
			telefono=$4,
			ciudad=$5,
			fecha_nacimiento=$6,
			id_genero=$7,
			nacionalidad=$8,
			foto_url=$9,
			id_roles_usuarios=$10,
			disponibilidad=$11,
			perfil_publico=$12,
			activo=$13
		WHERE id_usuarios = $14`,
		c.NOMBRE,
		c.APELLIDO,
		c.CORREO,
		c.TELEFONO,
		c.CIUDAD,
		c.FECHA_NACIMIENTO,
		c.ID_GENERO,
		c.NACIONALIDAD,
		c.FOTO_URL,
		c.ID_ROLES_USUARIOS,
		c.DISPONIBLE,
		c.PERFIL_PUBLICO,
		c.ACTIVO,
		id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

func DeleteUsuarios(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM core.usuarios WHERE id_usuarios = $1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
