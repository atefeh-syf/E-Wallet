package endpoints

import (
	"context"
	"errors"

	"github.com/atefeh-syf/yumigo/internal"
	"github.com/atefeh-syf/yumigo/pkg/wallet"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"os"
)

type Set struct {
    GetEndpoint           endpoint.Endpoint
}

func NewEndpointSet(svc wallet.Service) Set {
    return Set{
        GetEndpoint:           MakeGetEndpoint(svc),
    }
}

func MakeGetEndpoint(svc wallet.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetRequest)
        docs, err := svc.Get(ctx, req.Filters...)
        if err != nil {
            return GetResponse{docs, err.Error()}, nil
        }
        return GetResponse{docs, ""}, nil
    }
}

func (s *Set) Get(ctx context.Context, filters ...internal.Filter) (internal.Wallet, error) {
    resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
    if err != nil {
        return internal.Wallet{}, err
    }
    getResp := resp.(GetResponse)
    if getResp.Err != "" {
        return internal.Wallet{}, errors.New(getResp.Err)
    }
    return getResp.Wallet, nil
}

var logger log.Logger

func init() {
    logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
    logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}