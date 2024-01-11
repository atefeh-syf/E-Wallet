package models

import "database/sql/driver"

type transactionType string

const (
	deposit  transactionType = "deposit"
	withdraw transactionType = "withdraw"
)

func (st *transactionType) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		*st = transactionType(b)
	}
	return nil
}

func (st transactionType) Value() (driver.Value, error) {
	return string(st), nil
}

type Transaction struct {
	BaseModel
	Type            transactionType `sql:"type:ENUM('deposit', 'withdraw')" json:"type" form:"type"`
	Amount          float64         `sql:"type:decimal(64,0);not null"`
	PreviousBalance float64         `sql:"type:decimal(64,0);not null"`
	Confirmed       bool            `gorm:"column:confirmed" json:"confirmed" form:"confirmed"`
	UserId          uint
	WalletId        uint
	User            User   `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Wallet          Wallet `gorm:"foreignKey:WalletId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
