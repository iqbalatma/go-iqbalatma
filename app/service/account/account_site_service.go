package account

import (
	"iqbalatma/go-iqbalatma/app/dto"
	"iqbalatma/go-iqbalatma/app/model"
	repository2 "iqbalatma/go-iqbalatma/app/repository"

	"github.com/gin-gonic/gin"
)

type AccountSiteService interface {
	GetAllData(c *gin.Context) ([]model.AccountSite, error)
	GetDataById(c *gin.Context, id string) (*model.AccountSite, error)
	AddNewData(c *gin.Context, request dto.CreateAccountSiteRequest) (*model.AccountSite, error)
}

type AccountSiteServiceImp struct {
	AccountSiteRepository repository2.AccountSiteRepository
}

func (this AccountSiteServiceImp) GetAllData(c *gin.Context) ([]model.AccountSite, error) {
	data, err := this.AccountSiteRepository.GetAllData(c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (this AccountSiteServiceImp) GetDataById(c *gin.Context, id string) (*model.AccountSite, error) {
	data, err := this.AccountSiteRepository.GetDataById(c, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (this AccountSiteServiceImp) AddNewData(c *gin.Context, request dto.CreateAccountSiteRequest) (*model.AccountSite, error) {
	accountSite := model.AccountSite{
		Name:        request.Name,
		Description: &request.Description,
		Url:         request.Url,
	}
	data, err := this.AccountSiteRepository.AddNewData(c, &accountSite)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewAccountSiteService() AccountSiteService {
	return &AccountSiteServiceImp{
		AccountSiteRepository: repository2.NewAccountSiteRepository(),
	}
}
