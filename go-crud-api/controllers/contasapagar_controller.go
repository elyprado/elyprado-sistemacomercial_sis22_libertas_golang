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

func GetContas(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//
	defer db.Close()
	rows, err := db.Query("SELECT idpagar, data, valor,vencimento,pagamento,valorpago,idfornecedor FROM conta_pagar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var contas []models.Conta
	for rows.Next() {
		// A variável conta_pagar armazena os dados da tabela do arquivo contasapagar, do pacote models
		var conta_pagar models.Conta
		if err := rows.Scan(&conta_pagar.ID_Pagar, &conta_pagar.Data, &conta_pagar.Valor, &conta_pagar.Vencimento, &conta_pagar.Pagamento, &conta_pagar.Valorpago, &conta_pagar.Idfornecedor); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		contas = append(contas, conta_pagar)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contas)
}

func GetConta(w http.ResponseWriter, r *http.Request) {
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

	var conta models.Conta
	err = db.QueryRow("SELECT * FROM conta_pagar WHERE idpagar = ?", id).Scan(&conta.ID_Pagar, &conta.Data, &conta.Valor, &conta.Vencimento, &conta.Pagamento, &conta.Valorpago, &conta.Idfornecedor)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conta)
}

func CreateConta(w http.ResponseWriter, r *http.Request) {
	var conta models.Conta
	if err := json.NewDecoder(r.Body).Decode(&conta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO conta_pagar (idpagar, data, valor, vencimento, pagamento, valorpago, idfornecedor) VALUES (?, ?, ?, ?, ?, ?, ?)", conta.ID_Pagar, conta.Data, conta.Valor, conta.Vencimento, conta.Pagamento, conta.Valorpago, conta.Idfornecedor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	conta.ID_Pagar = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conta)
}

func UpdateConta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var conta models.Conta
	if err := json.NewDecoder(r.Body).Decode(&conta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE conta_pagar SET data = ?, valor = ?, vencimento = ?, pagamento = ?, valorpago = ?, idfornecedor = ? WHERE idpagar = ?", conta.Data, conta.Valor, conta.Vencimento, conta.Pagamento, conta.Valorpago, conta.Idfornecedor, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	conta.ID_Pagar = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conta)
}

func DeleteConta(w http.ResponseWriter, r *http.Request) {
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

	_, err = db.Exec("DELETE FROM conta_pagar WHERE idpagar = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
