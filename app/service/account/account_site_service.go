package account

import (
	"github.com/gin-gonic/gin"
	"iqbalatma/go-iqbalatma/app/interface/repository"
	"iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/model"
	repository2 "iqbalatma/go-iqbalatma/app/repository"
)

type AccountSiteService struct {
	AccountSiteRepository repository.AccountSiteRepository
}

func (a AccountSiteService) GetAllData(c *gin.Context) ([]model.AccountSite, error) {
	panic("implement me")
}

func NewAccountSiteService() service.AccountSiteService {
	return &AccountSiteService{
		AccountSiteRepository: repository2.NewAccountSiteRepository(),
	}
}
