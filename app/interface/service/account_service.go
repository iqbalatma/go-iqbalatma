package service

import (
	"iqbalatma/go-iqbalatma/app/model"

	"github.com/gin-gonic/gin"
)

type AccountSiteService interface {
	GetAllData(c *gin.Context) ([]model.AccountSite, error)
	GetDataById(c *gin.Context, id string) (*model.AccountSite, error)
}
