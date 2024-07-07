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

	router.HandleFunc("/cliente", controllers.GetClientes).Methods("GET")
	router.HandleFunc("/cliente/pesquisa", controllers.GetClientes).Methods("GET")
	router.HandleFunc("/cliente/{id}", controllers.GetCliente).Methods("GET")
	router.HandleFunc("/cliente", controllers.CreateCliente).Methods("POST")
	router.HandleFunc("/cliente/{id}", controllers.UpdateCliente).Methods("PUT")
	router.HandleFunc("/cliente/{id}", controllers.DeleteCliente).Methods("DELETE")

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

  router.HandleFunc("/produto", controllers.GetProdutos).Methods("GET")
	router.HandleFunc("/produto/pesquisa", controllers.GetProduto).Methods("GET")
	router.HandleFunc("/produto/{id}", controllers.GetProduto).Methods("GET")
	router.HandleFunc("/produto", controllers.CreateProduto).Methods("POST")
	router.HandleFunc("/produto/{id}", controllers.UpdateProduto).Methods("PUT")
	router.HandleFunc("/produto/{id}", controllers.DeleteProduto).Methods("DELETE")

	router.HandleFunc("/vendedor", controllers.GetVendedor).Methods("GET")
	router.HandleFunc("/vendedor/{id}", controllers.GetVendedorByID).Methods("GET")
	router.HandleFunc("/vendedor", controllers.CreateVendedor).Methods("POST")
	router.HandleFunc("/vendedor/{id}", controllers.UpdateVendedor).Methods("PUT")
	router.HandleFunc("/vendedor/{id}", controllers.DeleteVendedor).Methods("DELETE")

	router.HandleFunc("/caixa", controllers.GetCaixas).Methods("GET")
	router.HandleFunc("/caixa/{id}", controllers.GetCaixa).Methods("GET")
	router.HandleFunc("/caixa", controllers.CreateCaixa).Methods("POST")
	router.HandleFunc("/caixa/{id}", controllers.UpdateCaixa).Methods("PUT")
	router.HandleFunc("/caixa/{id}", controllers.DeleteCaixa).Methods("DELETE")

	router.HandleFunc("/cidade", controllers.GetCidades).Methods("GET")
	router.HandleFunc("/cidade/{id}", controllers.GetCidade).Methods("GET")
	router.HandleFunc("/cidade", controllers.CreateCidade).Methods("POST")
	router.HandleFunc("/cidade/{id}", controllers.UpdateCidade).Methods("PUT")
	router.HandleFunc("/cidade/{id}", controllers.DeleteCidade).Methods("DELETE")

	router.HandleFunc("/compra", controllers.GetCompra).Methods("GET")
	router.HandleFunc("/compra/{id}", controllers.GetCompraByID).Methods("GET")
	router.HandleFunc("/compra", controllers.CreateCompra).Methods("POST")
	router.HandleFunc("/compra/{id}", controllers.UpdateCompra).Methods("PUT")
	router.HandleFunc("/compra/{id}", controllers.DeleteCompra).Methods("DELETE")

	router.HandleFunc("/contasareceber", controllers.GetContasReceber).Methods("GET")
	router.HandleFunc("/contasareceber/{id}", controllers.GetContaReceber).Methods("GET")
	router.HandleFunc("/contasareceber", controllers.CreateContasReceber).Methods("POST")
	router.HandleFunc("/contasareceber/{id}", controllers.UpdateContasReceber).Methods("PUT")
	router.HandleFunc("/contasareceber/{id}", controllers.DeleteContasReceber).Methods("DELETE")

	router.HandleFunc("/conta", controllers.GetContas).Methods("GET")
	router.HandleFunc("/conta/{id}", controllers.GetConta).Methods("GET")
	router.HandleFunc("/conta", controllers.CreateConta).Methods("POST")
	router.HandleFunc("/conta/{id}", controllers.UpdateConta).Methods("PUT")
	router.HandleFunc("/conta/{id}", controllers.DeleteConta).Methods("DELETE")


	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	return router
}
