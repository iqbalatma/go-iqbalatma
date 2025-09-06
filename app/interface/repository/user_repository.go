package repository

import (
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetAllPaginated(c *gin.Context) (*utils.Payload, error)
}
