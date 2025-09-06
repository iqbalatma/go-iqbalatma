package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	Index(c *gin.Context) error
	Show(c *gin.Context) error
	Store(c *gin.Context) error
	Update(c *gin.Context) error
	Destroy(c *gin.Context) error
}
