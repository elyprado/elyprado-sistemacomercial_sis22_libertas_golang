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

// GET
func GetCompra(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idcompra, numeronf, data, quantidade, valor, idproduto, idfornecedor FROM compra")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var compras []models.Compra
	for rows.Next() {
		var compra models.Compra
		if err := rows.Scan(&compra.ID, &compra.NumeroNF, &compra.Data, &compra.Quantidade, &compra.Valor, &compra.IdProduto, &compra.IdFornecedor); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		compras = append(compras, compra)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compras)
}

// GET ID
func GetCompraByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid compra ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var compra models.Compra
	err = db.QueryRow("SELECT idcompra, numeronf, data, quantidade, valor, idproduto, idfornecedor FROM compra WHERE idcompra = ?", id).Scan(&compra.ID, &compra.NumeroNF, &compra.Data, &compra.Quantidade, &compra.Valor, &compra.IdProduto, &compra.IdFornecedor)
	if err == sql.ErrNoRows {
		http.Error(w, "compra not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compra)
}

// POST
func CreateCompra(w http.ResponseWriter, r *http.Request) {
	var compra models.Compra
	if err := json.NewDecoder(r.Body).Decode(&compra); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO compra (numeronf, data, quantidade, valor, idproduto, idfornecedor) VALUES (?, ?, ?, ?, ?, ?)", compra.NumeroNF, compra.Data, compra.Quantidade, compra.Valor, compra.IdProduto, compra.IdFornecedor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	compra.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compra)
}

// PUT
func UpdateCompra(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid compra ID", http.StatusBadRequest)
		return
	}

	var compra models.Compra
	if err := json.NewDecoder(r.Body).Decode(&compra); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE compra SET numeronf = ?, data = ?, quantidade = ?, valor = ?, idproduto = ?, idfornecedor = ? WHERE idcompra = ?", compra.NumeroNF, compra.Data, compra.Quantidade, compra.Valor, compra.IdProduto, compra.IdFornecedor, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	compra.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compra)
}

// DELETE
func DeleteCompra(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid Compra ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM compra WHERE idcompra = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
