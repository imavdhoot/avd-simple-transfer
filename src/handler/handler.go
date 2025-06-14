package handler

import (
	"database/sql"

	"github.com/imavdhoot/avd-simple-transfer/src/repository"
	"github.com/imavdhoot/avd-simple-transfer/src/service"
)

type Handler struct {
	*AccountHandler
	*TransactionHandler
}

func New(db *sql.DB) *Handler {
	repo := repository.New(db)
	accSvc := service.NewAccountService(repo)
	txSvc := service.NewTransactionService(repo)

	return &Handler{
		AccountHandler: &AccountHandler{svc: accSvc},
		TransactionHandler: &TransactionHandler{svc: txSvc},
	}
}
