package repositories

import (
	"context"

	"sync"

	"github.com/atefeh-syf/yumigo/pkg/wallet/data/models"
	"github.com/atefeh-syf/yumigo/pkg/wallet/data/db"
	"github.com/atefeh-syf/yumigo/pkg/wallet/dto"
	"gorm.io/gorm"
)

type WalletRepositoryInterface interface {
	//FindAll() ([]models.Wallet, error)
	//FindById(id string) (models.Wallet, error)
	//Create(user models.Wallet) (models.Wallet, error)
	//Update(id string, user models.Wallet) (models.Wallet, error)
	//Exists(request requests.WalletRequest)
	//Count() int64
}

type WalletRepository struct {
	DB        *gorm.DB
	WaitGroup *sync.WaitGroup
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{
		DB:        db.GetDb(),
		WaitGroup: &sync.WaitGroup{},
	}
}

func (repo *WalletRepository) FindWalletByUserId(userId int, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	wallet := models.Wallet{}

	err := repo.DB.Where("user_id = ?", userId).First(&wallet).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		channel <- models.DBResponse{
			Data:  wallet,
			Error: err,
		}
	} else {
		channel <- models.DBResponse{
			Data:  wallet,
			Error: nil,
		}
	}
}

// Updates wallet
func (repo *WalletRepository) UpdateWalletBalance(walletBalanceUpdate *dto.WalletBalanceUpdate, channel chan models.DBResponse) {
	defer repo.WaitGroup.Done()
	wallet := walletBalanceUpdate.Wallet
	ctx := context.Background()
	tx := repo.DB.WithContext(ctx).Begin()
	err := tx.Model(&models.Wallet{}).Where("id = ?", wallet.ID).Updates(wallet).Error

	t := models.Transaction{
		Type:            walletBalanceUpdate.Type,
		Amount:          walletBalanceUpdate.Amount,
		PreviousBalance: walletBalanceUpdate.PreviousBalance,
		Confirmed:       true,
		UserId:          wallet.UserId,
		WalletId:        wallet.ID,
	}
	err = repo.DB.Create(&t).Error
	tx.Commit()
	channel <- models.DBResponse{
		Data:  wallet,
		Error: err,
	}
}
