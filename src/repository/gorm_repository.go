package repository

import (
	"log"
	"context"
	"errors"
	"strings"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/jackc/pgconn"
	"github.com/imavdhoot/avd-simple-transfer/src/model"
	"github.com/imavdhoot/avd-simple-transfer/src/constant"
)

type Repository struct{ db *gorm.DB }

func New(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) CreateAccount(ctx context.Context, acc model.Account) error {
	rid := ctx.Value("request_id")
	log.Printf("[RID=%s][RepoCreateAccount] creating a new account for AccountID:: %d", rid, acc.AccountID)	
	
	err := r.db.WithContext(ctx).Create(&acc).Error
	if err != nil {
		if duplicatedKey(err) {
			return constant.ErrAccountExists
		}
	}
	return err
}

func (r *Repository) GetAccount(ctx context.Context, id int64) (model.Account, error) {
	rid := ctx.Value("request_id")
	log.Printf("[RID=%s][RepoGetAccount] fetching AccountID:: %d", rid, id)
	var acc model.Account

	err := r.db.WithContext(ctx).First(&acc, "account_id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return acc, constant.ErrAccountNotFound
	}
	return acc, err
}

func (r *Repository) TransferTx(ctx context.Context, srcID, dstID int64, amt float64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		rid := ctx.Value("request_id")
		log.Printf("[RID=%s][RepoTransferTx] executing the transfer from:: %d to:: %d of Amount:: %f",
			rid, srcID, dstID, amt)

		var src model.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&src, "account_id = ?", srcID).Error; err != nil {
			return constant.ErrSrcAccountNotFound
		}
		if src.Balance < amt {
			return constant.ErrInsufficientFund
		}

		var dst model.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&dst, "account_id = ?", dstID).Error; err != nil {
			return constant.ErrDstAccountNotFound
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



// duplicatedKey returns true if err is, or wraps, a duplicate-key error.
func duplicatedKey(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}

	return strings.Contains(err.Error(), "duplicate key value")
}