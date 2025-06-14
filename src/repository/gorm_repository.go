// src/repository/gorm_repository.go
package repository

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/imavdhoot/avd-simple-transfer/src/model"
	"github.com/imavdhoot/avd-simple-transfer/src/constant"
)

type Repository struct{ db *gorm.DB }

func New(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) CreateAccount(ctx context.Context, acc model.Account) error {
	return r.db.WithContext(ctx).Create(&acc).Error
}

func (r *Repository) GetAccount(ctx context.Context, id int64) (model.Account, error) {
	var acc model.Account
	err := r.db.WithContext(ctx).First(&acc, "account_id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return acc, constant.ErrAccountNotFound
	}
	return acc, err
}

func (r *Repository) TransferTx(ctx context.Context, srcID, dstID int64, amt float64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var src model.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&src, "account_id = ?", srcID).Error; err != nil {
			return constant.ErrAccountNotFound
		}
		if src.Balance < amt {
			return constant.ErrInsufficientFund
		}

		if err := tx.Model(&model.Account{}).
			Where("account_id = ?", srcID).
			Update("balance", gorm.Expr("balance - ?", amt)).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.Account{}).
			Where("account_id = ?", dstID).
			Update("balance", gorm.Expr("balance + ?", amt)).Error; err != nil {
			return err
		}

		return tx.Create(&model.Transaction{
			SourceAccountID:      srcID,
			DestinationAccountID: dstID,
			Amount:               amt,
		}).Error
	})
}
