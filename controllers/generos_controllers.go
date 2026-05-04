package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func GetGeneros(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_genero, nombre, activo FROM catalogo.generos")
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Genero

	for rows.Next() {
		var c models.Genero

		err := rows.Scan(&c.ID_GENERO, &c.NOMBRE, &c.ACTIVO)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

func GetGenerosByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Genero

	err = config.DB.QueryRow(
		"SELECT id_genero, nombre, activo FROM catalogo.generos WHERE id_genero = $1", id,
	).Scan(&c.ID_GENERO, &c.NOMBRE, &c.ACTIVO)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "ID no encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

func CreateGeneros(w http.ResponseWriter, r *http.Request) {
	var c models.Genero

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(
		"INSERT INTO catalogo.generos (nombre, activo) VALUES ($1,$2) RETURNING id_genero",
		c.NOMBRE, c.ACTIVO,
	).Scan(&c.ID_GENERO)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

func UpdateGeneros(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Genero

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(
		"UPDATE catalogo.generos SET nombre=$1, activo=$2 WHERE id_genero=$3",
		c.NOMBRE, c.ACTIVO, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato actualizado"})
}

func DeleteGeneros(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM catalogo.generos WHERE id_genero=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Dato eliminado"})
}
