package wallet

import (
	"context"
	"os"
	"github.com/atefeh-syf/yumigo/internal"
	"github.com/go-kit/log"
	//"github.com/lithammer/shortuuid/v3"
)
type WalletService struct{}

func NewService() Service { 
	return &WalletService{} 
}

var logger log.Logger 

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

func (w *WalletService) Get(_ context.Context, filters ...internal.Filter) (internal.Wallet, error) {
    doc := internal.Wallet{
        Name    : "test",
		Type       : "test",
		Balance   : 1000,
		Slug        : "",
		Description : "test",
		UserId      : 1,
    }
    return doc, nil
}