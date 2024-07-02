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

func GetFornecedores(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idfornecedor, nome, cnpj, logradouro, numero, bairro, cep, telefone, idcidade FROM fornecedor")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var fornecedores []models.Fornecedor
	for rows.Next() {
		var fornecedor models.Fornecedor
		if err := rows.Scan(&fornecedor.ID, &fornecedor.Nome, &fornecedor.Cnpj, &fornecedor.Logradouro, &fornecedor.Numero, &fornecedor.Bairro, &fornecedor.Cep, &fornecedor.Telefone, &fornecedor.IDCidade); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fornecedores = append(fornecedores, fornecedor)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fornecedores)
}

func GetFornecedor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID do fornecedor inválido", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var fornecedor models.Fornecedor
	err = db.QueryRow("SELECT idfornecedor, nome, cnpj, logradouro, numero, bairro, cep, telefone, idcidade FROM fornecedor WHERE idfornecedor = ?", id).Scan(&fornecedor.ID, &fornecedor.Nome, &fornecedor.Cnpj, &fornecedor.Logradouro, &fornecedor.Numero, &fornecedor.Bairro, &fornecedor.Cep, &fornecedor.Telefone, &fornecedor.IDCidade)
	if err == sql.ErrNoRows {
		http.Error(w, "Fornecedor não encontrado", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fornecedor)
}

func CreateFornecedor(w http.ResponseWriter, r *http.Request) {
	var fornecedor models.Fornecedor
	if err := json.NewDecoder(r.Body).Decode(&fornecedor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO fornecedor (nome, cnpj, logradouro, numero, bairro, cep, telefone, idcidade) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", fornecedor.Nome, fornecedor.Cnpj, fornecedor.Logradouro, fornecedor.Numero, fornecedor.Bairro, fornecedor.Cep, fornecedor.Telefone, fornecedor.IDCidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fornecedor.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fornecedor)
}

func UpdateFornecedor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID fornecedor inválido", http.StatusBadRequest)
		return
	}

	var fornecedor models.Fornecedor
	if err := json.NewDecoder(r.Body).Decode(&fornecedor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE fornecedor SET nome = ?, cnpj = ?, logradouro = ?, numero = ?, bairro = ?, cep = ?, telefone = ?, idcidade = ? WHERE idfornecedor = ?", fornecedor.Nome, fornecedor.Cnpj, fornecedor.Logradouro, fornecedor.Numero, fornecedor.Bairro, fornecedor.Cep, fornecedor.Telefone, fornecedor.IDCidade, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fornecedor.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fornecedor)
}

func DeleteFornecedor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID fornecedor inválido", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM fornecedor WHERE idfornecedor = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
