package wallet

import (
	"context"

	"github.com/atefeh-syf/yumigo/internal"
	"github.com/atefeh-syf/yumigo/internal/wallet/data/models"
)

type Service interface {
    Get(ctx context.Context, userId string ,filters ...internal.Filter) (models.Wallet, error)
	ServiceStatus(ctx context.Context) (int, error)
}
