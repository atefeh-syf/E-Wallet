package repositories

import (
	//"context"

	"sync"

	//"github.com/atefeh-syf/yumigo/pkg/user/data/models"
	"github.com/atefeh-syf/yumigo/pkg/user/data/db"
	//"github.com/atefeh-syf/yumigo/pkg/user/dto"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	//FindAll() ([]models.User, error)
	//FindById(id string) (models.User, error)
	//Create(user models.User) (models.User, error)
	//Update(id string, user models.User) (models.User, error)
	//Exists(request requests.UserRequest)
	//Count() int64
}

type UserRepository struct {
	DB        *gorm.DB
	WaitGroup *sync.WaitGroup
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB:        db.GetDb(),
		WaitGroup: &sync.WaitGroup{},
	}
}

