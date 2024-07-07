
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

func GetProdutos(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idproduto, descricao, precocusto, precovenda, saldoestoque, codbarras, idmarca FROM produto")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var produtos []models.Produto
	for rows.Next() {
		var produto models.Produto
		if err := rows.Scan(&produto.ID, &produto.Descricao, &produto.PrecoCusto, &produto.PrecoVenda, &produto.SaldoEstoque, &produto.CodBarras, &produto.IdMarca); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, produto)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid produto ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var produto models.Produto
	err = db.QueryRow("SELECT idproduto, descricao, precocusto,precovenda,saldoestoque,codbarras,idmarca FROM produto WHERE idproduto = ?", id).Scan(&produto.ID, &produto.Descricao, &produto.PrecoCusto, &produto.PrecoVenda, &produto.SaldoEstoque, &produto.CodBarras, &produto.IdMarca)
	if err == sql.ErrNoRows {
		http.Error(w, "produto not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	var produto models.Produto
	if err := json.NewDecoder(r.Body).Decode(&produto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO produto (descricao, precocusto, precovenda, saldoestoque, codbarras, idmarca) VALUES (?, ?, ?, ?, ?, ?)", produto.ID, produto.Descricao, produto.PrecoCusto, produto.PrecoVenda, produto.SaldoEstoque, produto.CodBarras, produto.IdMarca)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	produto.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid produto ID", http.StatusBadRequest)
		return
	}

	var produto models.Produto  
	if err := json.NewDecoder(r.Body).Decode(&produto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE produto SET descricao = ?, precocusto = ?, precovenda = ?, saldoestoque =?, codbarras = ?, idmarca = ?  WHERE idproduto = ?", produto.Descricao, produto.PrecoCusto, produto.PrecoVenda, produto.SaldoEstoque, produto.CodBarras, produto.IdMarca, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	produto.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid produto ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM produto WHERE idproduto = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
