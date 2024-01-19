package dto

import "github.com/atefeh-syf/E-Wallet/data/models"

type TransactionAction struct {
	Amount float64 `json:"amount" binding:"required,min=1"`
	UserId int     `json:"user_id" binding:"required,min=1"`
}

type WalletBalanceUpdate struct {
	Amount          float64
	PreviousBalance float64
	Type            models.TransactionType
	Wallet          models.Wallet
}

type TransactionResult struct {
	Wallet models.Wallet `json:"wallet"`
	Result bool
}
