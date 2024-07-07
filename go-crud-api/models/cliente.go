package models

type Cliente struct {
	ID         int    `json:"idcliente"`
	Nome       string `json:"nome"`
	Cpf        string `json:"cpf"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Cep        string `json:"cep"`
	Telefone   string `json:"telefone"`
	Idcidade   int    `json:"idcidade"`
}
