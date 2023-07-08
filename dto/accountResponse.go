package dto

type AccountResponse struct {
	AccountId   string  `json:"account_id"`
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	OpeningDate string  `json:"opening_date"`
	Status      string  `json:"status"`
}
