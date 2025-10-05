package service

import (
	"iqbalatma/go-iqbalatma/app/model"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authenticate(c *gin.Context) (*model.User, error)
}
