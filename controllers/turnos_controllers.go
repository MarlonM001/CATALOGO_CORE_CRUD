package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

func GetTurno(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_turno, nombre, descripcion, activo FROM turnos")
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Turno

	for rows.Next() {
		var c models.Turno

		err := rows.Scan(&c.ID_TURNOS, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

func GetTurnoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Turno

	err = config.DB.QueryRow(
		"SELECT id_turno, nombre, descripcion, activo FROM id_turno WHERE turnos = $1",
		id,
	).Scan(&c.ID_TURNOS, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

func CreateTurno(w http.ResponseWriter, r *http.Request) {
	var c models.Turno

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(
		"INSERT INTO id_turno (nombre, descripcion, activo) VALUES ($1,$2,$3) RETURNING turnos",
		c.NOMBRE, c.DESCRIPCION, c.ACTIVO,
	).Scan(&c.ID_TURNOS)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

func UpdateTurno(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.Turno

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(
		"UPDATE id_turno SET nombre=$1, descripcion=$2, activo=$3 WHERE turnos = $4",
		c.NOMBRE, c.DESCRIPCION, c.ACTIVO, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

func DeleteTurno(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM id_turno WHERE turnos = $1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
