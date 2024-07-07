package models

type Venda struct {
	ID       int    `json:"idvenda"`
	NumeroNF     string `json:"numeronf"`
	Data    string `json:"data"`
	Quantidade    string `json:"quantidade"`
	Valor float64 `json:"valor"`
	Comissao float64 `json:"comissao"`
	IdCliente int `json:"idcliente"`
	IdProduto int `json:"idproduto"`
	IdVendedor int `json:"idvendedor"`
}