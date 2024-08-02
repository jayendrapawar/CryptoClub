package models

type CryptoData struct {
	Currency string  `json:"code"`
	Price    float64 `json:"rate"`
}
