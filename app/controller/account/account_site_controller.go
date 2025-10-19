package account

import (
	"github.com/gin-gonic/gin"
	interfaceservice "iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/service/account"
	"iqbalatma/go-iqbalatma/utils"
	"net/http"
)

type AccountSiteController struct {
	AccountSiteService interfaceservice.AccountSiteService
}

func (this AccountSiteController) Index(c *gin.Context) error {
	data, err := this.AccountSiteService.GetAllData(c)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, utils.NewHttpSuccess("Get all data user successfully", &utils.Payload{Data: data}))

	return nil
}

func NewAccountSiteController() *AccountSiteController {
	return &AccountSiteController{
		AccountSiteService: account.NewAccountSiteService(),
	}
}
