package repository

import (
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetAllPaginated(c *gin.Context) (*utils.Payload, error)
	GetByEmail(c *gin.Context, email string) (*model.User, error)
}
