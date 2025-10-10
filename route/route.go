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
	authController := auth.NewAuthController()

	{
		authRoute := apiRoute.Group("/auth")
		{
			authRoute.POST("/authenticate", ErrorHandleWrapper(authController.Authenticate))
			authRoute.POST("/refresh", middleware.RefreshMiddleware(), ErrorHandleWrapper(authController.Refresh))
		}
	}

	authenticatedRoute := apiRoute.Group("")
	authenticatedRoute.Use(middleware.AuthMiddleware())

	authenticatedRoute.POST("/auth/logout", ErrorHandleWrapper(authController.Logout))

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
