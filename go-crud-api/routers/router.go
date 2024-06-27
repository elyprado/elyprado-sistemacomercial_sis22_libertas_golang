package routers

import (
	"go-crud-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/usuario", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/usuario/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/usuario", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/usuario/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/usuario/{id}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/venda", controllers.GetVendas).Methods("GET")
	router.HandleFunc("/venda/{id}", controllers.GetVenda).Methods("GET")
	router.HandleFunc("/venda", controllers.CreateVenda).Methods("POST")
	router.HandleFunc("/venda/{id}", controllers.UpdateVenda).Methods("PUT")
	router.HandleFunc("/venda/{id}", controllers.DeleteVenda).Methods("DELETE")

	router.HandleFunc("/fornecedor", controllers.GetFornecedores).Methods("GET")
	router.HandleFunc("/fornecedor/{id}", controllers.GetFornecedor).Methods("GET")
	router.HandleFunc("/fornecedor", controllers.CreateFornecedor).Methods("POST")
	router.HandleFunc("/fornecedor/{id}", controllers.UpdateFornecedor).Methods("PUT")
	router.HandleFunc("/fornecedor/{id}", controllers.DeleteFornecedor).Methods("DELETE")

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	return router
}
