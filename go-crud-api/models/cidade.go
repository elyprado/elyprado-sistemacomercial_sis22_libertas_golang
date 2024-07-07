package models

type Cidade struct {
	ID         int     `json:"idcidade"`
	NomeCidade string  `json:"nomecidade"`
	Uf         string  `json:"uf"`
	CodigoIbge string  `json:"codigo_ibge"`
	População  int     `json:"população"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
}
