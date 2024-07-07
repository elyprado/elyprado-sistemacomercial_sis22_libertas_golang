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

func GetContasReceber(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idreceber, data, valor, vencimento, pagamento, valorpago, idcliente FROM conta_receber ORDER BY data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var contasareceber []models.ContaReceber
	for rows.Next() {
		var contareceber models.ContaReceber
		if err := rows.Scan(&contareceber.ID, &contareceber.Data, &contareceber.Valor, &contareceber.Vencimento, &contareceber.Pagamento, &contareceber.Valorpago, &contareceber.Idcliente); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		contasareceber = append(contasareceber, contareceber)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contasareceber)
}

func GetContaReceber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var contareceber models.ContaReceber
	err = db.QueryRow("SELECT idreceber, data, valor, vencimento, pagamento, valorpago, idcliente FROM conta_receber WHERE idreceber = ?", id).Scan(&contareceber.ID, &contareceber.Data, &contareceber.Valor, &contareceber.Vencimento, &contareceber.Pagamento, &contareceber.Valorpago, &contareceber.Idcliente)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contareceber)
}

func CreateContasReceber(w http.ResponseWriter, r *http.Request) {
	var contareceber models.ContaReceber
	if err := json.NewDecoder(r.Body).Decode(&contareceber); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO conta_receber (data, valor, vencimento, pagamento, valorpago, idcliente) VALUES (?, ?, ?, ?, ?, ?)", contareceber.Data, contareceber.Valor, contareceber.Vencimento, contareceber.Pagamento, contareceber.Valorpago, contareceber.Idcliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	contareceber.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contareceber)
}

func UpdateContasReceber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var contareceber models.ContaReceber
	if err := json.NewDecoder(r.Body).Decode(&contareceber); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE conta_receber SET data = ?, valor = ?, vencimento = ?, pagamento = ?, valorpago = ?, idcliente = ? WHERE idreceber = ?",
		&contareceber.Data, &contareceber.Valor, &contareceber.Vencimento, &contareceber.Pagamento, &contareceber.Valorpago, &contareceber.Idcliente, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contareceber.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contareceber)
}

func DeleteContasReceber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM conta_receber WHERE idreceber = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
