package dto

type CreateAccountRequest struct {
	AccountID      int64   `json:"account_id" binding:"required,gt=0"`
	InitialBalance float64 `json:"initial_balance" binding:"required"`
}

type AccountURI struct {
	AccountID int64 `uri:"account_id" binding:"required,gt=0"`
}