package wallet

import (
	"github.com/go-kit/kit/log"
)
type walletService struct{}

func NewService() Service { return &walletService{} }

var logger log.Logger 

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}