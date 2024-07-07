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

	router.HandleFunc("/produto", controllers.GetProduto).Methods("GET")
	router.HandleFunc("/produto/{id}", controllers.GetProduto).Methods("GET")
	router.HandleFunc("/produto", controllers.CreateProduto).Methods("POST")
	router.HandleFunc("/produto/{id}", controllers.UpdateProduto).Methods("PUT")
	router.HandleFunc("/produto/{id}", controllers.DeleteProduto).Methods("DELETE")

	router.HandleFunc("/vendedor", controllers.GetVendedor).Methods("GET")
	router.HandleFunc("/vendedor/{id}", controllers.GetVendedor).Methods("GET")
	router.HandleFunc("/vendedor", controllers.CreateVendedor).Methods("POST")
	router.HandleFunc("/vendedor/{id}", controllers.UpdateVendedor).Methods("PUT")
	router.HandleFunc("/vendedor/{id}", controllers.DeleteVendedor).Methods("DELETE")

	router.HandleFunc("/caixa", controllers.GetCaixas).Methods("GET")
	router.HandleFunc("/caixa/{id}", controllers.GetCaixa).Methods("GET")
	router.HandleFunc("/caixa", controllers.CreateCaixa).Methods("POST")
	router.HandleFunc("/caixa/{id}", controllers.UpdateCaixa).Methods("PUT")
	router.HandleFunc("/caixa/{id}", controllers.DeleteCaixa).Methods("DELETE")

	router.HandleFunc("/conta", controllers.GetConta).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	return router
}
