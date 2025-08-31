package management

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/service/management"
	"iqbalatma/go-iqbalatma/config"
	"net/http"
)

type UserController struct {
	UserService *management.UserService
}

func (ctrl *UserController) Index(c *gin.Context) {
	config.AppLogger.WithField("Nama", "IQBAL")
	config.AppLogger.Info("UserController.Index")
	service := management.NewUserService()
	data, err := service.GetAllData()

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (ctrl *UserController) Show(c *gin.Context) {
	service := management.NewUserService()
	data, err := service.GetDataById(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
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
	return &UserController{
		UserService: management.NewUserService(),
	}
}
