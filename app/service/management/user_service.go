package management

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/app/repository"
	exception "iqbalatma/go-iqbalatma/error"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(),
	}
}

func (service *UserService) GetAllData() (*[]model.User, error) {
	data, err := service.UserRepository.GetAllData()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *UserService) GetDataById(id string) (*model.User, error) {
	data, err := service.UserRepository.GetDataById(id)

	something := false
	if something {
		return nil, exception.InvalidAction()
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *UserService) UpdateDataById(c *gin.Context, id string) {

}

func (service *UserService) DeleteDataById(c *gin.Context, id string) {

}

func (service *UserService) AddNewData(c *gin.Context, id string) {

}
