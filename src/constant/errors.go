package constant

import "errors"

var (
	ErrAccountExists = errors.New("account already exist")
	ErrAccountNotFound  = errors.New("account not found")
	ErrInsufficientFund = errors.New("insufficient balance")
	ErrSrcAccountNotFound = errors.New("source account not found")
	ErrDstAccountNotFound = errors.New("destination account not found")
)
