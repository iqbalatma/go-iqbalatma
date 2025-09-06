package management

import (
	repository2 "iqbalatma/go-iqbalatma/app/interface/repository"
	"iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/repository"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserRepository repository2.UserRepository
}

func (service *UserService) GetAllPaginated(c *gin.Context) (*utils.Payload, error) {
	data, err := service.UserRepository.GetAllPaginated(c)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewUserService() service.UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(),
	}
}

//func (service *UserService) GetAllDataPaginated(c *gin.Context) (*[]model.User, error) {
//	data, err := service.UserRepository.GetAllPaginated(c)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return data, nil
//}
//
//func (service *UserService) GetAllData() (*[]model.User, error) {
//	data, err := service.UserRepository.GetAllData()
//
//	if err != nil {
//		return nil, err
//	}
//
//	return data, nil
//}
//
//func (service *UserService) GetDataById(id string) (*model.User, error) {
//	data, err := service.UserRepository.GetDataById(id)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return data, nil
//}
//
//func (service *UserService) UpdateDataById(c *gin.Context, id string) {
//
//}
//
//func (service *UserService) DeleteDataById(c *gin.Context, id string) {
//
//}
//
//func (service *UserService) AddNewData(c *gin.Context, id string) {
//
//}
