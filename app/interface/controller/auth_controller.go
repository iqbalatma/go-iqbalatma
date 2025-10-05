package controller

import "github.com/gin-gonic/gin"

type AuthControllerInterface interface {
	Authenticate(c *gin.Context) error
}
