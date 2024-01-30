package endpoints

import "github.com/atefeh-syf/yumigo/internal"


type GetRequest struct {
    Filters []internal.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
    Wallet internal.Wallet `json:"wallets"`
    Err       string              `json:"err,omitempty"`
}