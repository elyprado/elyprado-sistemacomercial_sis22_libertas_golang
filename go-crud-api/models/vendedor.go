package models

type Vendedor struct {
	ID            int    `json:"idvendedor"`
	Nome          string `json:"nome"`
	CPF           string `json:"cpf"`
	Logradouro    string `json:"logradouro"`
	Numero        string `json:"numero"`
	Bairro        string `json:"bairro"`
	CEP           string `json:"cep"`
	Telefone      string `json:"telefone"`
	Perc_comissao string `json:"perc_comissao"`
	IdCidade      string `json:"idcidade"`
}
