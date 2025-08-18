package route

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/controller/management"
)

func RegisterRoute(router *gin.Engine) {
	apiRoute := router.Group("/api")

	managementRoute := apiRoute.Group("/management")

	userController := management.NewUserController()
	users := managementRoute.Group("/users")
	users.GET("/", userController.Index)
	users.GET("/:id", userController.Show)
	users.POST("/", userController.Store)
	users.PATCH("/:id", userController.Update)
	users.DELETE("/:id", userController.Destroy)
}
