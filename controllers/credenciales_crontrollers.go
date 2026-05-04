package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

// 🔹 GET ALL
func GetCredenciales(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_credenciales, id_usuarios, password_hash, salt, doble_factor,
		       token_reset, token_expira_en, ultimo_login,
		       intentos_fallidos, bloqueado, activo
		FROM credenciales`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Credenciales

	for rows.Next() {
		var c models.Credenciales

		err := rows.Scan(
			&c.ID_CREDENCIALES,
			&c.ID_USUARIOS,
			&c.PASSWORD_HASH,
			&c.SALT,
			&c.DOBLE_FACTOR,
			&c.TOKEN_RESET,
			&c.TOKEN_EXPIRA_EN,
			&c.ULTIMO_LOGIN,
			&c.INTENTOS_FALLIDOS,
			&c.BLOQUEADO,
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

// 🔹 GET BY ID
func GetCredencialesByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Credenciales

	err = config.DB.QueryRow(`
		SELECT id_credenciales, id_usuarios, password_hash, salt, doble_factor,
		       token_reset, token_expira_en, ultimo_login,
		       intentos_fallidos, bloqueado, activo
		FROM credenciales
		WHERE id_credenciales = $1`,
		id,
	).Scan(
		&c.ID_CREDENCIALES,
		&c.ID_USUARIOS,
		&c.PASSWORD_HASH,
		&c.SALT,
		&c.DOBLE_FACTOR,
		&c.TOKEN_RESET,
		&c.TOKEN_EXPIRA_EN,
		&c.ULTIMO_LOGIN,
		&c.INTENTOS_FALLIDOS,
		&c.BLOQUEADO,
		&c.ACTIVO,
	)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

// 🔹 CREATE
func CreateCredenciales(w http.ResponseWriter, r *http.Request) {
	var c models.Credenciales

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(`
		INSERT INTO credenciales 
		(id_usuarios, password_hash, salt, doble_factor, token_reset,
		 token_expira_en, ultimo_login, intentos_fallidos, bloqueado, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING id_credenciales`,
		c.ID_USUARIOS,
		c.PASSWORD_HASH,
		c.SALT,
		c.DOBLE_FACTOR,
		c.TOKEN_RESET,
		c.TOKEN_EXPIRA_EN,
		c.ULTIMO_LOGIN,
		c.INTENTOS_FALLIDOS,
		c.BLOQUEADO,
		c.ACTIVO,
	).Scan(&c.ID_CREDENCIALES)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

// 🔹 UPDATE
func UpdateCredenciales(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Credenciales

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(`
		UPDATE credenciales SET
			id_usuarios=$1,
			password_hash=$2,
			salt=$3,
			doble_factor=$4,
			token_reset=$5,
			token_expira_en=$6,
			ultimo_login=$7,
			intentos_fallidos=$8,
			bloqueado=$9,
			activo=$10
		WHERE id_credenciales=$11`,
		c.ID_USUARIOS,
		c.PASSWORD_HASH,
		c.SALT,
		c.DOBLE_FACTOR,
		c.TOKEN_RESET,
		c.TOKEN_EXPIRA_EN,
		c.ULTIMO_LOGIN,
		c.INTENTOS_FALLIDOS,
		c.BLOQUEADO,
		c.ACTIVO,
		id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

// 🔹 DELETE
func DeleteCredenciales(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec(`DELETE FROM credenciales WHERE id_credenciales = $1`, id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}