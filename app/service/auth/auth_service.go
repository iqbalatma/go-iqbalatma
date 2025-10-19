package auth

import (
	"iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/app/repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authenticate(c *gin.Context) (*model.User, error)
}

type AuthServiceImp struct {
	UserRepository repository.UserRepository
}
type AuthenticateRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthServiceImp) Authenticate(c *gin.Context) (*model.User, error) {
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
	return &AuthServiceImp{
		UserRepository: repository.NewUserRepository(),
	}
}
