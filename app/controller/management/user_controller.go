package management

import (
	"iqbalatma/go-iqbalatma/app/enum"
	interfaceservice "iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/service/management"
	"iqbalatma/go-iqbalatma/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService interfaceservice.UserService
}

func (ctrl *UserController) Index(c *gin.Context) error {
	payload, err := ctrl.UserService.GetAllPaginated(c)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, &utils.HTTPResponse{
		Message:   "Get all data user successfully",
		Timestamp: time.Now(),
		Code:      enum.SUCCESS,
		Payload:   payload,
	})
	return nil
}

func (ctrl *UserController) Show(c *gin.Context) error {
	//service := management.NewUserService()
	//data, err := service.GetById()
	//
	//if err != nil {
	//	return err
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "success",
	//	"data":    data,
	//})

	return nil
}

func (ctrl *UserController) Store(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return nil
}

func (ctrl *UserController) Update(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})

	return nil
}

func (ctrl *UserController) Destroy(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return nil
}

func NewUserController() *UserController {
	return &UserController{
		UserService: management.NewUserService(),
	}
}
