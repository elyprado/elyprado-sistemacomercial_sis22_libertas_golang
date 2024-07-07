package models

type Caixa struct {
	Idcaixa       int    `json:"idcaixa"`
	Data          string `json:"data"`
	Descricao     string `json:"descricao"`
	Valor         string `json:"valor"`
	Debitocredito string `json:"debitocredito"`
}
