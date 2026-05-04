package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CATALOGO_CORE_CRUD/config"
	"CATALOGO_CORE_CRUD/models"

	"github.com/gorilla/mux"
)

func GetModalidadesTrabajo(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_modalidad_trabajo, nombre, descripcion, activo FROM catalogo.modalidades_trabajo")
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.ModalidadesTrabajo

	for rows.Next() {
		var c models.ModalidadesTrabajo

		err := rows.Scan(&c.ID_MODALIDAD_TRABAJO, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)
		if err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}

		list = append(list, c)
	}

	respondJSON(w, 200, list)
}

func GetModalidadesTrabajoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.ModalidadesTrabajo

	err = config.DB.QueryRow(
		"SELECT id_modalidad_trabajo, nombre, descripcion, activo FROM catalogo.modalidades_trabajo WHERE id_modalidad_trabajo = $1", id,
	).Scan(&c.ID_MODALIDAD_TRABAJO, &c.NOMBRE, &c.DESCRIPCION, &c.ACTIVO)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "ID no encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

func CreateModalidadesTrabajo(w http.ResponseWriter, r *http.Request) {
	var c models.ModalidadesTrabajo

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(
		"INSERT INTO catalogo.modalidades_trabajo (nombre, descripcion, activo) VALUES ($1,$2,$3) RETURNING id_modalidad_trabajo",
		c.NOMBRE, c.DESCRIPCION, c.ACTIVO,
	).Scan(&c.ID_MODALIDAD_TRABAJO)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

func UpdateModalidadesTrabajo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.ModalidadesTrabajo

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(
		"UPDATE catalogo.modalidades_trabajo SET nombre=$1, descripcion=$2, activo=$3 WHERE id_modalidad_trabajo = $4",
		c.NOMBRE, c.DESCRIPCION, c.ACTIVO, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Actualizado"})
}

func DeleteModalidadesTrabajo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM catalogo.modalidades_trabajo WHERE id_modalidad_trabajo = $1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
