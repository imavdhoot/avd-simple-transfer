package model

import "time"

type Transaction struct {
	ID                   uint        `gorm:"primaryKey"`
	SourceAccountID      int64       `json:"source_account_id"`
	DestinationAccountID int64       `json:"destination_account_id"`
	Amount               float64     `json:"amount" gorm:"type:numeric(20,8)"`
	CreatedAt            time.Time   `json:"created_at"`
}

func (Transaction) TableName() string { return "transactions" }
