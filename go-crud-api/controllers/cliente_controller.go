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

func GetClientes(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pesquisa := r.URL.Query().Get("pesquisa")
	if pesquisa != "" {
		pesquisa = " WHERE nome like '%" + pesquisa + "%' "
	}
	defer db.Close()
	rows, err := db.Query("SELECT idcliente, nome, cpf, logradouro,numero, bairro, cep, telefone, IFNULL(idcidade, 0) AS idcidade FROM cliente" + pesquisa + " ORDER BY idcliente")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var users []models.Cliente
	for rows.Next() {
		var user models.Cliente
		if err := rows.Scan(&user.ID, &user.Nome, &user.Cpf, &user.Logradouro, &user.Numero, &user.Bairro, &user.Cep, &user.Telefone, &user.Idcidade); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
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

	var user models.Cliente
	err = db.QueryRow("SELECT idcliente, nome, cpf, logradouro, numero, bairro, cep, telefone, IFNULL(idcidade, 0) AS idcidade FROM cliente WHERE idcliente = ?", id).Scan(&user.ID, &user.Nome, &user.Cpf, &user.Logradouro, &user.Numero, &user.Bairro, &user.Cep, &user.Telefone, &user.Idcidade)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var user models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO cliente (nome, cpf, logradouro, numero,bairro, cep, telefone, idcidade) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Nome, user.Cpf, user.Logradouro, user.Numero, user.Bairro, user.Cep, user.Telefone, user.Idcidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE cliente SET nome = ?, cpf = ?, logradouro = ?, numero = ?, bairro = ?, cep = ?, telefone = ?, idcidade = ? WHERE idcliente = ?", user.Nome, user.Cpf, user.Logradouro, user.Numero, user.Bairro, user.Cep, user.Telefone, user.Idcidade, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
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

	_, err = db.Exec("DELETE FROM cliente WHERE idcliente = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
