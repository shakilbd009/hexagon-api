package dto

type GetAccountResponse struct {
	AccountID   string  `json:"account_id,omitempty"`
	CustomerID  string  `json:"customer_id,omitempty"`
	OpeningDate string  `json:"opening_date,omitempty"`
	AccountType string  `json:"account_type,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Status      string  `json:"status,omitempty"`
}
