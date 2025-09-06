package service

import (
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetAllPaginated(c *gin.Context) (*utils.Payload, error)
	//GetAll(c *gin.Context) (*[]model.User, error)
	//GetById(c *gin.Context) (*model.User, error)
	//UpdateById(c *gin.Context) (*model.User, error)
	//AddNew(c *gin.Context) (*model.User, error)
	//DeleteById(c *gin.Context) error
}
