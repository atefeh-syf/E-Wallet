package services

import (
	"github.com/atefeh-syf/E-Wallet/data/models"
	"sync"
)

func WaitAndCloseChannel(wg *sync.WaitGroup, channel chan models.DBResponse) {
	go func(wg *sync.WaitGroup, channel chan models.DBResponse) {
		wg.Wait()
		close(channel)
	}(wg, channel)
}
