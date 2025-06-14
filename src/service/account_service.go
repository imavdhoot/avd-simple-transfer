package service

import (
	"context"

	"github.com/imavdhoot/avd-simple-transfer/src/model"
	"github.com/imavdhoot/avd-simple-transfer/src/repository"
)

type AccountService struct {
	repo *repository.Repository
}

func NewAccountService(r *repository.Repository) *AccountService {
	return &AccountService{repo: r}
}

func (s *AccountService) Create(ctx context.Context, acc model.Account) error {
	return s.repo.CreateAccount(ctx, acc)
}

func (s *AccountService) Get(ctx context.Context, id int64) (model.Account, error) {
	return s.repo.GetAccount(ctx, id)
}
