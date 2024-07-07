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

func GetCaixas(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idcaixa, data, descricao, valor, debitocredito FROM caixa ORDER BY descricao")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var caixas []models.Caixa
	for rows.Next() {
		var caixa models.Caixa
		if err := rows.Scan(&caixa.Idcaixa, &caixa.Data, &caixa.Descricao, &caixa.Valor, &caixa.Debitocredito); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		caixas = append(caixas, caixa)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(caixas)
}

func GetCaixa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid caixa ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var caixa models.Caixa
	err = db.QueryRow("SELECT idcaixa, data, descricao, valor, debitocredito FROM caixa WHERE idcaixa = ?", id).Scan(&caixa.Idcaixa, &caixa.Data, &caixa.Descricao, &caixa.Valor, &caixa.Debitocredito)
	if err == sql.ErrNoRows {
		http.Error(w, "Caixa not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(caixa)
}

func CreateCaixa(w http.ResponseWriter, r *http.Request) {
	var caixa models.Caixa
	if err := json.NewDecoder(r.Body).Decode(&caixa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO caixa (data, descricao, valor, debitocredito) VALUES (?, ?, ?, ?)", caixa.Data, caixa.Descricao, caixa.Valor, caixa.Debitocredito)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	caixa.Idcaixa = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(caixa)
}

func UpdateCaixa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid caixa ID", http.StatusBadRequest)
		return
	}

	var caixa models.Caixa
	if err := json.NewDecoder(r.Body).Decode(&caixa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE caixa SET data = ?, descricao = ?, valor = ?, debitocredito = ? WHERE idcaixa = ?", caixa.Data, caixa.Descricao, caixa.Valor, caixa.Debitocredito, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	caixa.Idcaixa = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(caixa)
}

func DeleteCaixa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid caixa ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM caixa WHERE idcaixa = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
