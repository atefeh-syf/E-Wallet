package services

import (
	"context"

	"github.com/atefeh-syf/yumigo/pkg/user/api/dto"
	"github.com/atefeh-syf/yumigo/pkg/user/config"
	"github.com/atefeh-syf/yumigo/pkg/user/data/db"
	"github.com/atefeh-syf/yumigo/pkg/user/data/models"
	"github.com/atefeh-syf/yumigo/pkg/user/pkg/logging"
	"gorm.io/gorm"
)

type UserAddressService struct {
	base     *BaseService[models.UserAddress, dto.CreateUserAddressModelRequest, dto.UpdateUserAddressModelRequest, dto.UserAddressResponse]
	logger   logging.Logger
	cfg      *config.Config
	database *gorm.DB
}

func NewUserAddressService(cfg *config.Config) *UserAddressService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserAddressService{
		base:     &BaseService[models.UserAddress, dto.CreateUserAddressModelRequest, dto.UpdateUserAddressModelRequest, dto.UserAddressResponse]{
			Database: database,
			Logger:   logger,
		},
		cfg:      cfg,
		database: database,
		logger:   logger,
	}
}

func (s *UserAddressService) Create(ctx context.Context, req *dto.CreateUserAddressModelRequest) (*dto.UserAddressResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *UserAddressService) Update(ctx context.Context, id int, req *dto.UpdateUserAddressModelRequest) (*dto.UserAddressResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *UserAddressService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *UserAddressService) GetById(ctx context.Context, id int) (*dto.UserAddressResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *UserAddressService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.UserAddressResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
