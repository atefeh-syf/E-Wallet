package wallet

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/atefeh-syf/yumigo/internal"
	"github.com/atefeh-syf/yumigo/internal/wallet/data/models"
	"github.com/atefeh-syf/yumigo/pkg/wallet/data/repositories"
	"github.com/go-kit/log"
	//"github.com/lithammer/shortuuid/v3"
)

type WalletService struct {
	Repository *repositories.WalletRepository
	WaitGroup  *sync.WaitGroup
}

func NewWalletService() *WalletService {
	WalletRepository := repositories.NewWalletRepository()
	return &WalletService{
		Repository: WalletRepository,
		WaitGroup:  WalletRepository.WaitGroup,
	}
}

func (w *WalletService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking the Service health...")
	return http.StatusOK, nil
}
var logger log.Logger 

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

func (w *WalletService) Get(_ context.Context, userId string , filters ...internal.Filter) (models.Wallet, error) {
	channel := make(chan models.DBResponse) //channel for db response

	wallet := models.Wallet{}
	var err error

	w.WaitGroup.Add(1)
	user_id, _ := strconv.Atoi(userId)
	go w.Repository.FindWalletByUserId(user_id, channel)

	//wait for process to finish and close the channel
	WaitAndCloseChannel(w.WaitGroup, channel)

	//get data from channel
	for c := range channel {
		wallet = c.Data.(models.Wallet)
		err = c.Error
	}

	return wallet, err
}



func WaitAndCloseChannel(wg *sync.WaitGroup, channel chan models.DBResponse) {
	go func(wg *sync.WaitGroup, channel chan models.DBResponse) {
		wg.Wait()
		close(channel)
	}(wg, channel)
}