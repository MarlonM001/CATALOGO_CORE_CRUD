package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

func GetTipoDocumentos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_tipo_documento, nombre, descripcion, activo FROM core.tipo_documento")
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.TiposDocumentos

	for rows.Next() {
		var c models.TiposDocumentos

		err := rows.Scan(&c.ID_TIPO_DOCUMENTO, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

func GetTiposDocumentosByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.TiposDocumentos

	err = config.DB.QueryRow(
		"SELECT id_tipo_documento, nombre, descripcion, activo FROM core.tipo_documento WHERE id_tipo_documento = $1",
		id,
	).Scan(&c.ID_TIPO_DOCUMENTO, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

func CreateTiposDocumentos(w http.ResponseWriter, r *http.Request) {
	var c models.TiposDocumentos

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(
		"INSERT INTO core.tipo_documento (nombre, descripcion, activo) VALUES ($1,$2,$3) RETURNING id_tipo_documento",
		c.NOMBRE, c.DESCRIPCION, c.ACTIVO,
	).Scan(&c.ID_TIPO_DOCUMENTO)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

func UpdateTiposDocumentos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.TiposDocumentos

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(
		"UPDATE core.tipo_documento SET nombre=$1, descripcion=$2, activo=$3 WHERE id_tipo_documento = $4",
		c.NOMBRE, c.DESCRIPCION, c.ACTIVO, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

func DeleteTiposDocumentos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM core.tipo_documento WHERE id_tipo_documento = $1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
