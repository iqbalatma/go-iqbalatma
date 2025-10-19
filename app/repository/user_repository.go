package repository

import (
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetAllPaginated(c *gin.Context) (*utils.Payload, error)
	GetByEmail(c *gin.Context, email string) (*model.User, error)
}

type UserRepositoryImp struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImp{}
}

func (repository *UserRepositoryImp) GetAllPaginated(c *gin.Context) (*utils.Payload, error) {
	var users []model.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	paginate, err := utils.Paginate[model.User](c, config.DB, &users)
	if err != nil {
		return nil, err
	}
	return &utils.Payload{
		Data: users,
		Meta: paginate,
	}, nil
}

func (repository *UserRepositoryImp) GetByEmail(c *gin.Context, email string) (*model.User, error) {
	var user model.User

	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
