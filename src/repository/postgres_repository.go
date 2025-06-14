package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/imavdhoot/avd-simple-transfer/src/model"
)

var (
	ErrAccountNotFound  = errors.New("account not found")
	ErrInsufficientFund = errors.New("insufficient balance")
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateAccount(ctx context.Context, acc model.Account) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO accounts (account_id, balance) VALUES ($1, $2)",
		acc.AccountID, acc.Balance)
	return err
}

func (r *Repository) GetAccount(ctx context.Context, id int64) (model.Account, error) {
	var acc model.Account
	err := r.db.QueryRowContext(ctx,
		"SELECT account_id, balance FROM accounts WHERE account_id=$1", id).
		Scan(&acc.AccountID, &acc.Balance)
	if errors.Is(err, sql.ErrNoRows) {
		return acc, ErrAccountNotFound
	}
	return acc, err
}

// TransferTx performs an atomic transfer and logs it.
func (r *Repository) TransferTx(ctx context.Context, srcID, dstID int64, amt float64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // safe no-op on commit

	var srcBalance float64
	err = tx.QueryRowContext(ctx,
		"SELECT balance FROM accounts WHERE account_id=$1 FOR UPDATE", srcID).
		Scan(&srcBalance)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrAccountNotFound
	}
	if srcBalance < amt {
		return ErrInsufficientFund
	}

	// update balances
	if _, err = tx.ExecContext(ctx,
		"UPDATE accounts SET balance = balance - $1 WHERE account_id=$2", amt, srcID); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx,
		"UPDATE accounts SET balance = balance + $1 WHERE account_id=$2", amt, dstID); err != nil {
		return err
	}

	// insert transaction log
	if _, err = tx.ExecContext(ctx,
		`INSERT INTO transactions (source_account_id,destination_account_id,amount)
		 VALUES ($1,$2,$3)`,
		srcID, dstID, amt); err != nil {
		return err
	}

	return tx.Commit()
}
