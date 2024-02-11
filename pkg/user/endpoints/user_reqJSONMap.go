package endpoints

import (
	"github.com/atefeh-syf/yumigo/internal"
	"github.com/atefeh-syf/yumigo/pkg/wallet/data/models"
)

type GetRequest struct {
	UserId  string            `json:"user_id,omitempty"`
	Filters []internal.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
	Wallet models.Wallet `json:"wallets"`
	Err    string        `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
