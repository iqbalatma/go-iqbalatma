package route

import (
	"iqbalatma/go-iqbalatma/app/controller/auth"
	"iqbalatma/go-iqbalatma/app/controller/management"
	"iqbalatma/go-iqbalatma/middleware"

	"github.com/gin-gonic/gin"
)

func ErrorHandleWrapper(h func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			c.Error(err)
			c.Abort()
		}
	}
}

func RegisterRoute(router *gin.Engine) {
	apiRoute := router.Group("/api")

	{
		authRoute := apiRoute.Group("/auth")
		{
			authController := auth.NewAuthController()
			authRoute.POST("/authenticate", ErrorHandleWrapper(authController.Authenticate))
		}
	}

	authenticatedRoute := apiRoute.Group("")
	authenticatedRoute.Use(middleware.AuthMiddleware())

	{
		managementRoute := authenticatedRoute.Group("/management")
		{
			userController := management.NewUserController()
			users := managementRoute.Group("/users")
			users.GET("/", ErrorHandleWrapper(userController.Index))
			users.GET("/:id", ErrorHandleWrapper(userController.Show))
			users.POST("/", ErrorHandleWrapper(userController.Store))
			users.PATCH("/:id", ErrorHandleWrapper(userController.Update))
			users.DELETE("/:id", ErrorHandleWrapper(userController.Destroy))
		}
	}

}
