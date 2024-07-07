package controllers

import (
	"database/sql"
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 🔍Método de consulta de vendas no geral
func GetVendas(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idvenda, numeronf, data, quantidade, valor, comissao, idcliente, idproduto, idvendedor FROM venda")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var vendas []models.Venda
	for rows.Next() {
		var venda models.Venda
		if err := rows.Scan(&venda.ID, &venda.NumeroNF, &venda.Data, &venda.Quantidade, &venda.Valor, &venda.Comissao, &venda.IdCliente, &venda.IdProduto, &venda.IdVendedor); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vendas = append(vendas, venda)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendas)
}

// 🔍Método de consulta de venda por id
func GetVenda(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid venda ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var venda models.Venda
	err = db.QueryRow("SELECT idvenda,  numeronf, data, quantidade, valor, comissao, idcliente, idproduto, idvendedor FROM venda WHERE idvenda = ?", id).Scan(&venda.ID, &venda.NumeroNF, &venda.Data, &venda.Quantidade, &venda.Valor, &venda.Comissao, &venda.IdCliente, &venda.IdProduto, &venda.IdVendedor)
	if err == sql.ErrNoRows {
		http.Error(w, "Venda not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venda)
}

// ➕Método de inserção de usuário
func CreateVenda(w http.ResponseWriter, r *http.Request) {
	var venda models.Venda

	if err := json.NewDecoder(r.Body).Decode(&venda); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	// log.Println(*db)
	log.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO venda (numeronf, data, quantidade, valor, comissao, idcliente, idproduto, idvendedor) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", venda.NumeroNF, venda.Data, venda.Quantidade, venda.Valor, venda.Comissao, venda.IdCliente, venda.IdProduto, venda.IdVendedor)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	venda.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venda)
}

// ⬆️ Método atualizar venda
func UpdateVenda(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid venda ID", http.StatusBadRequest)
		return
	}

	

	var venda models.Venda
	if err := json.NewDecoder(r.Body).Decode(&venda); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE venda SET numeronf = ?, data = ?, quantidade = ?, valor = ?, comissao = ?, idcliente = ?, idproduto = ?, idvendedor = ? WHERE idvenda = ?", venda.NumeroNF, venda.Data, venda.Quantidade, venda.Valor, venda.Comissao, venda.IdCliente, venda.IdProduto, venda.IdVendedor, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	venda.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venda)
}

// 🗑️ Método excluir venda
func DeleteVenda(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid venda ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM venda WHERE idvenda = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
