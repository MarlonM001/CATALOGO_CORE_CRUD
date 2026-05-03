package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

func GetDiasSemanas(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_dia_semana, nombre, nombre_completo, orden FROM dias_semana")
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.DiaSemana

	for rows.Next() {
		var c models.DiaSemana

		err := rows.Scan(&c.ID_DIA_SEMANA, &c.NOMBRE, &c.NOMBRE_COMPLETO, &c.ORDEN)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

func GetDiasSemanaByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.DiaSemana

	err = config.DB.QueryRow(
		"SELECT id_dia_semana, nombre, nombre_completo, orden FROM dias_semana WHERE dias_semana = $1",
		id,
	).Scan(&c.ID_DIA_SEMANA, &c.NOMBRE, &c.NOMBRE_COMPLETO, &c.ORDEN)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

func CreateDiasSemana(w http.ResponseWriter, r *http.Request) {
	var c models.DiaSemana

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(
		"INSERT INTO id_dia_semana (nombre, nombre_completo, orden) VALUES ($1,$2,$3) RETURNING dias_semana",
		c.NOMBRE, c.NOMBRE_COMPLETO, c.ORDEN,
	).Scan(&c.ID_DIA_SEMANA)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

func UpdateDiasSemana(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.DiaSemana

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(
		"UPDATE id_dia_semana SET nombre=$1, nombre_completo=$2, orden=$3 WHERE dias_semana = $4",
		c.NOMBRE, c.NOMBRE_COMPLETO, c.ORDEN, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

func DeleteDiasSemana(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM id_dia_semana WHERE dias_semana = $1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
