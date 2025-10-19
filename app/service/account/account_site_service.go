package account

import (
	"iqbalatma/go-iqbalatma/app/interface/repository"
	"iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/model"
	repository2 "iqbalatma/go-iqbalatma/app/repository"

	"github.com/gin-gonic/gin"
)

type AccountSiteService struct {
	AccountSiteRepository repository.AccountSiteRepository
}

func (this AccountSiteService) GetAllData(c *gin.Context) ([]model.AccountSite, error) {
	data, err := this.AccountSiteRepository.GetAllData(c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (this AccountSiteService) GetDataById(c *gin.Context, id string) (*model.AccountSite, error) {
	data, err := this.AccountSiteRepository.GetDataById(c, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewAccountSiteService() service.AccountSiteService {
	return &AccountSiteService{
		AccountSiteRepository: repository2.NewAccountSiteRepository(),
	}
}
