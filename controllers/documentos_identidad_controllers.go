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
func GetDocumentosIdentidad(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_documentos_identidad, id_usuarios, id_tipo_documento,
		       numero_documento, fecha_expedicion, lugar_expedicion, activo
		FROM core.documentos_identidad`)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.DocumentosIdentidad

	for rows.Next() {
		var c models.DocumentosIdentidad

		err := rows.Scan(
			&c.ID_DOCUMENTOS_IDENTIDAD,
			&c.ID_USUARIOS,
			&c.ID_TIPO_DOCUMENTO,
			&c.NUMERO_DOCUMENTO,
			&c.FECHA_EXPEDICION,
			&c.LUGAR_EXPEDICION,
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
func GetDocumentosIdentidadByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.DocumentosIdentidad

	err = config.DB.QueryRow(`
		SELECT id_documentos_identidad, id_usuarios, id_tipo_documento,
		       numero_documento, fecha_expedicion, lugar_expedicion, activo
		FROM core.documentos_identidad
		WHERE id_documentos_identidad = $1`,
		id,
	).Scan(
		&c.ID_DOCUMENTOS_IDENTIDAD,
		&c.ID_USUARIOS,
		&c.ID_TIPO_DOCUMENTO,
		&c.NUMERO_DOCUMENTO,
		&c.FECHA_EXPEDICION,
		&c.LUGAR_EXPEDICION,
		&c.ACTIVO,
	)

	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "No encontrado"})
		return
	}

	respondJSON(w, 200, c)
}

// 🔹 CREATE
func CreateDocumentosIdentidad(w http.ResponseWriter, r *http.Request) {
	var c models.DocumentosIdentidad

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err = config.DB.QueryRow(`
		INSERT INTO core.documentos_identidad
		(id_usuarios, id_tipo_documento, numero_documento,
		 fecha_expedicion, lugar_expedicion, activo)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id_documentos_identidad`,
		c.ID_USUARIOS,
		c.ID_TIPO_DOCUMENTO,
		c.NUMERO_DOCUMENTO,
		c.FECHA_EXPEDICION,
		c.LUGAR_EXPEDICION,
		c.ACTIVO,
	).Scan(&c.ID_DOCUMENTOS_IDENTIDAD)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, c)
}

// 🔹 UPDATE
func UpdateDocumentosIdentidad(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	var c models.DocumentosIdentidad

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err = config.DB.Exec(`
		UPDATE core.documentos_identidad SET
			id_usuarios=$1,
			id_tipo_documento=$2,
			numero_documento=$3,
			fecha_expedicion=$4,
			lugar_expedicion=$5,
			activo=$6
		WHERE id_documentos_identidad = $7`,
		c.ID_USUARIOS,
		c.ID_TIPO_DOCUMENTO,
		c.NUMERO_DOCUMENTO,
		c.FECHA_EXPEDICION,
		c.LUGAR_EXPEDICION,
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
func DeleteDocumentosIdentidad(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondJSON(w, 400, map[string]string{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Exec(`
		DELETE FROM core.documentos_identidad
		WHERE id_documentos_identidad = $1`, id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Eliminado"})
}
