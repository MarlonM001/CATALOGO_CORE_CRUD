package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

func GetEstados(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_estado, contexto, nombre, descripcion, activo FROM estados")
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Estados

	for rows.Next() {
		var c models.Estados

		err := rows.Scan(&c.ID_ESTADO, &c.CONTEXTO, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

func GetEstadosByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Estados

	err = config.DB.QueryRow(
		"SELECT id_estado, contexto, nombre, descripcion, activo FROM estados WHERE estados = $1",
		id,
	).Scan(&c.ID_ESTADO, &c.CONTEXTO, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

func CreateEstado(w http.ResponseWriter, r *http.Request) {
	var c models.Estados

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(
		"INSERT INTO estados (contexto, nombre, descripcion, activo) VALUES ($1,$2,$3,$4) RETURNING estados",
		c.CONTEXTO, c.NOMBRE, c.DESCRIPCION, c.ACTIVO,
	).Scan(&c.ID_ESTADO)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

func UpdateEstado(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Estados

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(
		"UPDATE estados SET contexto=$1, nombre=$2, descripcion=$3, activo=$4 WHERE estados = $5",
		c.CONTEXTO, c.NOMBRE, c.DESCRIPCION, c.ACTIVO, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

func DeleteEstado(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM estados WHERE estados = $1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
