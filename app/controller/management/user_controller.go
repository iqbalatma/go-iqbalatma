package management

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/service/management"
	exception "iqbalatma/go-iqbalatma/error"
	"net/http"
)

type UserController struct {
}

func (ctrl *UserController) Index(c *gin.Context) {
	somethingWentWrong := true

	if somethingWentWrong {
		c.Error(exception.InvalidAction())
		return
	}

	service := management.NewUserService()
	data, err := service.GetAllData(c)

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (ctrl *UserController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (ctrl *UserController) Store(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (ctrl *UserController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (ctrl *UserController) Destroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func NewUserController() *UserController {
	return &UserController{}
}
