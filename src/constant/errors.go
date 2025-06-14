package constant

import "errors"

var (
	ErrAccountNotFound  = errors.New("account not found")
	ErrInsufficientFund = errors.New("insufficient balance")
)
