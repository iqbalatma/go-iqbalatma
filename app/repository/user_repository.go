package repository

import (
	"iqbalatma/go-iqbalatma/app/interface/repository"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (repository *UserRepository) GetAllPaginated(c *gin.Context) (*utils.Payload, error) {
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
