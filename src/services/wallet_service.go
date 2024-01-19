package services

import (
	"github.com/atefeh-syf/E-Wallet/api/dto"
	"github.com/atefeh-syf/E-Wallet/data/models"
	"github.com/atefeh-syf/E-Wallet/data/repositories"
	"sync"
)

type WalletService struct {
	Repository *repositories.WalletRepository
	WaitGroup  *sync.WaitGroup
}

var walletService *WalletService

// Init new service
func NewWalletService(walletRepo *repositories.WalletRepository) {
	walletService = &WalletService{Repository: walletRepo, WaitGroup: walletRepo.WaitGroup}
}

func GetWalletService() *WalletService {
	return walletService
}

func (s *WalletService) Deposit(req *dto.TransactionAction) (*dto.TransactionResult, error) {
	result := true
	wallet, err := s.FindWalletByUserId(req.UserId)
	if err != nil {
		return nil, err
	}

	channel := make(chan models.DBResponse) //channel for db response

	s.WaitGroup.Add(1)
	previousBalance := wallet.Balance
	wallet.Balance = wallet.Balance + req.Amount

	walletBalanceUpdate := &dto.WalletBalanceUpdate{
		Amount:          req.Amount,
		PreviousBalance: previousBalance,
		Type:            models.Deposit,
		Wallet:          wallet,
	}
	go s.Repository.UpdateWalletBalance(walletBalanceUpdate, channel)

	WaitAndCloseChannel(s.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		wallet = c.Data.(models.Wallet)
		err = c.Error
	}

	if err != nil {
		result = false
	}
	return &dto.TransactionResult{Wallet: wallet, Result: result}, err
}

func (s *WalletService) FindWalletByUserId(userId int) (models.Wallet, error) {
	channel := make(chan models.DBResponse) //channel for db response

	wallet := models.Wallet{}
	var err error

	s.WaitGroup.Add(1)
	go s.Repository.FindWalletByUserId(userId, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(s.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		wallet = c.Data.(models.Wallet)
		err = c.Error
	}

	return wallet, err
}

func (s *WalletService) Withdraw(req *dto.TransactionAction) (*dto.TransactionResult, error) {
	return &dto.TransactionResult{
		Wallet: models.Wallet{
			BaseModel:   models.BaseModel{},
			Name:        "",
			Type:        "",
			Balance:     0,
			Slug:        "",
			Description: "",
			UserId:      0,
			User:        models.User{},
		},
		Result: false,
	}, nil
}

func (s *WalletService) ForceWithdraw(req *dto.TransactionAction) (*dto.TransactionResult, error) {
	return &dto.TransactionResult{
		Wallet: models.Wallet{
			BaseModel:   models.BaseModel{},
			Name:        "",
			Type:        "",
			Balance:     0,
			Slug:        "",
			Description: "",
			UserId:      0,
			User:        models.User{},
		},
		Result: false,
	}, nil
}
