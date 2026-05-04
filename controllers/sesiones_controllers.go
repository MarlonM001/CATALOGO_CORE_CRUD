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
func GetSesion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_sesiones, id_usuarios, token, dispositivo, ip,
		       activa, expira_en, creado_en, actualizado_en
		FROM core.sesiones`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Sesion

	for rows.Next() {
		var s models.Sesion

		err := rows.Scan(
			&s.ID_SESIONES,
			&s.ID_USUARIOS,
			&s.TOKEN,
			&s.DISPOSITIVO,
			&s.IP,
			&s.ACTIVA,
			&s.EXPIRA_EN,
			&s.CREADO_EN,
			&s.ACTUALIZADO_EN,
		)

		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, s)
	}

	respondJSON(w, 200, list)
}

// 🔹 GET BY ID
func GetSesionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var s models.Sesion

	err = config.DB.QueryRow(`
		SELECT id_sesiones, id_usuarios, token, dispositivo, ip,
		       activa, expira_en, creado_en, actualizado_en
		FROM core.sesiones
		WHERE id_sesiones = $1`,
		id,
	).Scan(
		&s.ID_SESIONES,
		&s.ID_USUARIOS,
		&s.TOKEN,
		&s.DISPOSITIVO,
		&s.IP,
		&s.ACTIVA,
		&s.EXPIRA_EN,
		&s.CREADO_EN,
		&s.ACTUALIZADO_EN,
	)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		return
	}

	respondJSON(w, 200, s)
}

// 🔹 CREATE
func CreateSesion(w http.ResponseWriter, r *http.Request) {
	var s models.Sesion

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(`
		INSERT INTO core.sesiones
		(id_usuarios, token, dispositivo, ip, activa, expira_en)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id_sesiones, creado_en, actualizado_en`,
		s.ID_USUARIOS,
		s.TOKEN,
		s.DISPOSITIVO,
		s.IP,
		s.ACTIVA,
		s.EXPIRA_EN,
	).Scan(
		&s.ID_SESIONES,
		&s.CREADO_EN,
		&s.ACTUALIZADO_EN,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, s)
}

// 🔹 UPDATE
func UpdateSesion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var s models.Sesion

	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(`
		UPDATE core.sesiones SET
			id_usuarios=$1,
			token=$2,
			dispositivo=$3,
			ip=$4,
			activa=$5,
			expira_en=$6,
			actualizado_en=NOW()
		WHERE id_sesiones=$7`,
		s.ID_USUARIOS,
		s.TOKEN,
		s.DISPOSITIVO,
		s.IP,
		s.ACTIVA,
		s.EXPIRA_EN,
		id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

// 🔹 DELETE
func DeleteSesion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec(`
		DELETE FROM core.sesiones 
		WHERE id_sesiones = $1`, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
