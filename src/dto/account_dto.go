package dto

type CreateAccountRequest struct {
	AccountID      int64   `json:"account_id" binding:"required,gt=0"`
	InitialBalance float64 `json:"initial_balance" binding:"required"`
}
