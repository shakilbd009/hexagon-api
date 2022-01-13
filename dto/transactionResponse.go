package dto

type NewTransactionResponse struct {
	TransactionID string  `json:"transaction_id,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
}
