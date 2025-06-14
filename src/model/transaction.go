package model

type Transaction struct {
	SourceAccountID      int64   `json:"source_account_id"`
	DestinationAccountID int64   `json:"destination_account_id"`
	Amount               float64 `json:"amount"`
}
