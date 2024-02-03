package wallet

import (
	"context"

	"github.com/atefeh-syf/yumigo/internal"
)

type Service interface {
    Get(ctx context.Context, userId string ,filters ...internal.Filter) (internal.Wallet, error)
	ServiceStatus(ctx context.Context) (int, error)
}
