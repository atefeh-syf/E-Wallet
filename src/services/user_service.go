package services

import (
	"github.com/atefeh-syf/E-Wallet/api/dto"
	"github.com/atefeh-syf/E-Wallet/config"
	"github.com/atefeh-syf/E-Wallet/data/db"
	"github.com/atefeh-syf/E-Wallet/data/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	cfg          *config.Config
	database     *gorm.DB
	tokenService *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	return &UserService{
		cfg:          cfg,
		database:     database,
		tokenService: NewTokenService(cfg),
	}
}

func (s *UserService) Login(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.
		Model(&user).
		Where("username = ?", req.Username).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	tdto := tokenDto{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, MobileNumber: user.MobileNumber}
	token, err := s.tokenService.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil
}
