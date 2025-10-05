package auth

import (
	repository2 "iqbalatma/go-iqbalatma/app/interface/repository"
	"iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/app/repository"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	UserRepository repository2.UserRepository
}
type AuthenticateRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthService) Authenticate(c *gin.Context) (*model.User, error) {
	var request AuthenticateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	user, err := a.UserRepository.GetByEmail(c, request.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewAuthService() service.AuthService {
	return &AuthService{
		UserRepository: repository.NewUserRepository(),
	}
}
