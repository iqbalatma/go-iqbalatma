package repository

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/interface/repository"
	"iqbalatma/go-iqbalatma/app/model"
)

type AccountSiteRepository struct {
}

func (a AccountSiteRepository) GetAllData(c *gin.Context) ([]model.AccountSite, error) {
	panic("implement me")
}

func NewAccountSiteRepository() repository.AccountSiteRepository {
	return &AccountSiteRepository{}
}
