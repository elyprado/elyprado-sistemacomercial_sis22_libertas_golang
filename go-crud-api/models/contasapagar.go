package models

type Conta struct {
	ID_Pagar     int    `json:"idpagar"`
	Data         string `json:"data"`
	Valor        string `json:"valor"`
	Vencimento   string `json:"vencimento"`
	Pagamento    string `json:"pagamento"`
	Valorpago    string `json:"valorpago"`
	Idfornecedor string `json:"idfornecedor"`
}
