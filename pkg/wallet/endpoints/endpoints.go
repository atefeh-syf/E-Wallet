package endpoints

import (
	"context"
	"errors"

	"os"

	"github.com/atefeh-syf/yumigo/internal"
	"github.com/atefeh-syf/yumigo/internal/wallet/data/models"
	"github.com/atefeh-syf/yumigo/pkg/wallet"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Set struct {
    GetEndpoint           endpoint.Endpoint
    ServiceStatusEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc wallet.Service) Set {
    return Set{
        GetEndpoint:           MakeGetEndpoint(svc),
        ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
    }
}

func MakeGetEndpoint(svc wallet.Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetRequest)
        walllet, err := svc.Get(ctx, req.UserId, req.Filters...)
        if err != nil {
            return GetResponse{walllet, err.Error()}, nil
        }
        return GetResponse{walllet, ""}, nil
    }
}

func MakeServiceStatusEndpoint(svc wallet.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, filters ...internal.Filter) (models.Wallet, error) {
    resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
    if err != nil {
        return models.Wallet{}, err
    }
    getResp := resp.(GetResponse)
    if getResp.Err != "" {
        return models.Wallet{}, errors.New(getResp.Err)
    }
    return getResp.Wallet, nil
}

var logger log.Logger

func init() {
    logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
    logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}