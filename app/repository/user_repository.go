package repository

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repository *UserRepository) GetAllData() (*[]model.User, error) {
	var users []model.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (repository *UserRepository) GetDataById(id string) (*model.User, error) {
	var user model.User
	result := config.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepository) AddNewData(c *gin.Context) {

}

func (repository *UserRepository) UpdateDataById(c *gin.Context) {

}

func (repository *UserRepository) DeleteDataById(c *gin.Context) {

}
