package models

type Compra struct {
	ID           int    `json:"idcompra"`
	NumeroNF     string `json:"numeronf"`
	Data         string `json:"data"`
	Quantidade   string `json:"quantidade"`
	Valor        string `json:"valor"`
	IdProduto    string `json:"idproduto"`
	IdFornecedor string `json:"idfornecedor"`
}
