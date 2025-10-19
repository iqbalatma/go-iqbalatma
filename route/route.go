package route

import (
	"iqbalatma/go-iqbalatma/app/controller/account"
	"iqbalatma/go-iqbalatma/app/controller/auth"
	"iqbalatma/go-iqbalatma/app/controller/management"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/middleware"
	"iqbalatma/go-iqbalatma/utils"
	"net/http"

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
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.NewHttpError("Route does not exist", enum.ERR_ROUTE_NOT_FOUND))
	})

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

	{
		accountRoute := authenticatedRoute.Group("/account")
		{
			accountSiteController := account.NewAccountSiteController()
			accountSites := accountRoute.Group("/account-sites")
			accountSites.GET("/", ErrorHandleWrapper(accountSiteController.Index))
			accountSites.GET("/:id", ErrorHandleWrapper(accountSiteController.Show))
			accountSites.POST("/", ErrorHandleWrapper(accountSiteController.Store))
		}
	}

}
