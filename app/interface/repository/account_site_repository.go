package repository

import (
	"iqbalatma/go-iqbalatma/app/model"

	"github.com/gin-gonic/gin"
)

type AccountSiteRepository interface {
	GetAllData(c *gin.Context) ([]model.AccountSite, error)
	GetDataById(c *gin.Context, id string) (*model.AccountSite, error)
}
