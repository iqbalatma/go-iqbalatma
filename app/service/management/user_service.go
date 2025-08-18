package management

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/app/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(),
	}
}

func (service *UserService) GetAllData(c *gin.Context) ([]model.User, error) {
	fmt.Println("GET ALL DATA")
	data, err := service.UserRepository.GetAllData(c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *UserService) GetDataById(c *gin.Context, id string) {

}

func (service *UserService) UpdateDataById(c *gin.Context, id string) {

}

func (service *UserService) DeleteDataById(c *gin.Context, id string) {

}

func (service *UserService) AddNewData(c *gin.Context, id string) {

}
