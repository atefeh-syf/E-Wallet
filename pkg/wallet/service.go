package wallet

import (
	"context"

	"github.com/atefeh-syf/yumigo/internal"
)

type Service interface {
    Get(ctx context.Context, filters ...internal.Filter) (internal.Wallet, error)
}
