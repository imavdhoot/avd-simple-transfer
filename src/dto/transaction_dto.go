package dto

type SubmitTransactionRequest struct {
	SourceAccountID      int64   `json:"source_account_id"      binding:"required,gt=0"`
	DestinationAccountID int64   `json:"destination_account_id" binding:"required,gt=0,nefield=SourceAccountID"`
	Amount               float64 `json:"amount"                 binding:"required,gt=0"`
}

type SubmitTransactionResponse struct {
	TransactionID uint   `json:"transaction_id"`    // DB-generated PK
	Message       string `json:"message"`           // "success" (or any human string)
	Status        int    `json:"status"`            // 200, 201, etc.
	CreatedAt     string `json:"created_at"`        // RFC3339 string
	RequestID     string `json:"request_id"`        // associated internal request id
}