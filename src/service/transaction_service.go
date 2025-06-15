package service

import (
	"context"

	"github.com/imavdhoot/avd-simple-transfer/src/model"
	"github.com/imavdhoot/avd-simple-transfer/src/repository"
)

type TransactionService struct {
	repo *repository.Repository
}

func NewTransactionService(r *repository.Repository) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) Transfer(ctx context.Context, srcID, dstID int64,
	amt float64) (model.Transaction, error) {
	return s.repo.TransferTx(ctx, srcID, dstID, amt)
}
