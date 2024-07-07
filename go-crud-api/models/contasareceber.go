package models

type ContaReceber struct {
	ID         int    `json:"idreceber"`
	Data       string `json:"data"`
	Valor      string `json:"valor"`
	Vencimento string `json:"vencimento"`
	Pagamento  string `json:"pagamento"`
	Valorpago  string `json:"valorpago"`
	Idcliente  int    `json:"idcliente"`
}
