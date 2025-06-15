package dto

type SubmitTransactionRequest struct {
	SourceAccountID      int64   `json:"source_account_id"      binding:"required,gt=0"`
	DestinationAccountID int64   `json:"destination_account_id" binding:"required,gt=0,nefield=SourceAccountID"`
	Amount               float64 `json:"amount"                 binding:"required,gt=0"`
}