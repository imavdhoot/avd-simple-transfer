package model

type Account struct {
	AccountID int64   `json:"account_id" gorm:"primaryKey;column:account_id"`
	Balance   float64 `json:"balance"    gorm:"type:numeric(20,8)"`
}

func (Account) TableName() string { return "accounts" }
