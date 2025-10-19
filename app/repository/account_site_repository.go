package repository

import (
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"

	"github.com/gin-gonic/gin"
)

type AccountSiteRepository interface {
	GetAllData(c *gin.Context) ([]model.AccountSite, error)
	GetDataById(c *gin.Context, id string) (*model.AccountSite, error)
	AddNewData(c *gin.Context, accountSite *model.AccountSite) (*model.AccountSite, error)
}

type AccountSiteRepositoryImp struct {
}

func (this AccountSiteRepositoryImp) GetAllData(c *gin.Context) ([]model.AccountSite, error) {
	var accountSites []model.AccountSite
	err := config.DB.Find(&accountSites).Error
	if err != nil {
		return nil, err
	}

	return accountSites, nil
}

func (this AccountSiteRepositoryImp) GetDataById(c *gin.Context, id string) (*model.AccountSite, error) {
	var accountSite model.AccountSite

	err := config.DB.Where("id = ?", id).First(&accountSite).Error
	if err != nil {
		return nil, err
	}

	return &accountSite, nil
}

func (this AccountSiteRepositoryImp) AddNewData(c *gin.Context, accountSite *model.AccountSite) (*model.AccountSite, error) {
	err := config.DB.Create(&accountSite).Error
	if err != nil {
		return nil, err
	}

	return accountSite, nil
}

func NewAccountSiteRepository() AccountSiteRepository {
	return &AccountSiteRepositoryImp{}
}
