package models

import "database/sql/driver"

type TransactionType string

const (
	deposit  TransactionType = "deposit"
	withdraw TransactionType = "withdraw"
)

func (st *TransactionType) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		*st = TransactionType(b)
	}
	return nil
}

func (st TransactionType) Value() (driver.Value, error) {
	return string(st), nil
}

type Transaction struct {
	BaseModel
	Type            TransactionType `gorm:"type:transaction_type" json:"type" form:"type"`
	Amount          float64         `sql:"type:decimal(64,0);not null"`
	PreviousBalance float64         `sql:"type:decimal(64,0);not null"`
	Confirmed       bool            `gorm:"column:confirmed" json:"confirmed" form:"confirmed"`
	User            User            `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Wallet          Wallet          `gorm:"foreignKey:WalletId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
