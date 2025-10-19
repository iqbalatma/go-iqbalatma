package repository

import (
	"iqbalatma/go-iqbalatma/app/interface/repository"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"

	"github.com/gin-gonic/gin"
)

type AccountSiteRepository struct {
}

func (this AccountSiteRepository) GetAllData(c *gin.Context) ([]model.AccountSite, error) {
	var accountSites []model.AccountSite
	err := config.DB.Find(&accountSites).Error
	if err != nil {
		return nil, err
	}

	return accountSites, nil
}

func (this AccountSiteRepository) GetDataById(c *gin.Context, id string) (*model.AccountSite, error) {
	var accountSite model.AccountSite

	err := config.DB.Where("id = ?", id).First(&accountSite).Error
	if err != nil {
		return nil, err
	}

	return &accountSite, nil
}

func NewAccountSiteRepository() repository.AccountSiteRepository {
	return &AccountSiteRepository{}
}
