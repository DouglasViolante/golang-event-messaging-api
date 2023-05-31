package model

type Payment struct {
	TransactionID string  `json:"transactionid"`
	Amount        float32 `json:"amount"`
}
