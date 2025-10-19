package service

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/model"
)

type AccountSiteService interface {
	GetAllData(c *gin.Context) ([]model.AccountSite, error)
}
