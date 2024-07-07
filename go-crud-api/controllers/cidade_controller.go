package controllers

import (
	"database/sql"
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCidades(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM cidade")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cidades []models.Cidade
	for rows.Next() {
		var cidade models.Cidade
		if err := rows.Scan(&cidade.ID, &cidade.NomeCidade, &cidade.Uf, &cidade.CodigoIbge, &cidade.População, &cidade.Latitude, &cidade.Longitude); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cidades = append(cidades, cidade)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cidades)
}

func GetCidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid cidade ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var cidade models.Cidade
	err = db.QueryRow("SELECT * FROM cidade WHERE idcidade = ?", id).Scan(&cidade.ID, &cidade.NomeCidade, &cidade.Uf, &cidade.CodigoIbge, &cidade.População, &cidade.Latitude, &cidade.Longitude)
	if err == sql.ErrNoRows {
		http.Error(w, "Cidade not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cidade)
}

func CreateCidade(w http.ResponseWriter, r *http.Request) {
	var cidade models.Cidade
	if err := json.NewDecoder(r.Body).Decode(&cidade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO cidade (nomecidade, uf, codigo_ibge, população, latitude, longitude) VALUES (?, ?, ?, ?, ?, ?)",
		cidade.NomeCidade, cidade.Uf, cidade.CodigoIbge, cidade.População, cidade.Latitude, cidade.Longitude)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cidade.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cidade)
}

func UpdateCidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid cidade ID", http.StatusBadRequest)
		return
	}

	var cidade models.Cidade
	if err := json.NewDecoder(r.Body).Decode(&cidade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE cidade SET nomecidade = ?, uf = ?, codigo_ibge = ?, população = ?, latitude = ?, longitude = ? WHERE idcidade = ?",
		cidade.NomeCidade, cidade.Uf, cidade.CodigoIbge, cidade.População, cidade.Latitude, cidade.Longitude, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cidade.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cidade)
}

func DeleteCidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid cidade ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM cidade WHERE idcidade = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
