package repository

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/model"
)

type AccountSiteRepository interface {
	GetAllData(c *gin.Context) ([]model.AccountSite, error)
}
