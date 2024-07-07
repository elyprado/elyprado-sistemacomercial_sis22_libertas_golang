package models

type Fornecedor struct {
	ID         int    `json:"idfornecedor"`
	Nome       string `json:"nome"`
	Cnpj       string `json:"cnpj"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Cep        string `json:"cep"`
	Telefone   string `json:"telefone"`
	IDCidade   int    `json:"idcidade"`
}
