package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

func GetModalidadesSuscripcion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_modalidad_suscripcion, nombre, descripcion, activo FROM catalogo.modalidades_suscripcion")
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.ModalidadesSuscripcion

	for rows.Next() {
		var c models.ModalidadesSuscripcion

		err := rows.Scan(&c.ID_MODALIDAD_SUSCRIPCION, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

func GetModalidadesSuscripcionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.ModalidadesSuscripcion

	err = config.DB.QueryRow(
		
		"SELECT id_modalidad_suscripcion, nombre, descripcion, activo FROM catalogo.modalidades_suscripcion WHERE id_modalidad_suscripcion = $1",
		id,
	).Scan(&c.ID_MODALIDAD_SUSCRIPCION, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

func CreateModalidadesSuscripcion(w http.ResponseWriter, r *http.Request) {
	var c models.ModalidadesSuscripcion

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(
		"INSERT INTO catalogo.modalidades_suscripcion (nombre, descripcion, activo) VALUES ($1,$2,$3) RETURNING id_modalidad_suscripcion",
		c.NOMBRE, c.DESCRIPCION, c.ACTIVO,
	).Scan(&c.ID_MODALIDAD_SUSCRIPCION)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

func UpdateModalidadesSuscripcion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.ModalidadesSuscripcion

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(
		"UPDATE catalogo.modalidades_suscripcion SET nombre=$1, descripcion=$2, activo=$3 WHERE id_modalidad_suscripcion = $4",
		c.NOMBRE, c.DESCRIPCION, c.ACTIVO, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

func DeleteModalidadesSuscripcion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM catalogo.modalidades_suscripcion WHERE id_modalidad_suscripcion = $1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
