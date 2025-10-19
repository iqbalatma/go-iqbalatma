package account

import (
	interfaceservice "iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/resource"
	"iqbalatma/go-iqbalatma/app/service/account"
	"iqbalatma/go-iqbalatma/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountSiteController struct {
	AccountSiteService interfaceservice.AccountSiteService
}

func (this AccountSiteController) Index(c *gin.Context) error {
	data, err := this.AccountSiteService.GetAllData(c)

	if err != nil {
		return err
	}

	var responseCollection []map[string]interface{}
	for _, item := range data {
		utils.ToDateTime(&item.CreatedAt)
		responseCollection = append(responseCollection, map[string]interface{}{
			"id":          item.ID,
			"name":        item.Name,
			"description": item.Description,
			"url":         item.Url,
			"icon":        item.Icon,
			"created_at":  utils.ToDateTime(&item.CreatedAt),
		})
	}
	c.JSON(http.StatusOK, utils.NewHttpSuccess("Get all data account sites successfully", &utils.Payload{Data: resource.ToAccountSiteResourceCollection(data)}))

	return nil
}

func (this AccountSiteController) Show(c *gin.Context) error {
	data, err := this.AccountSiteService.GetDataById(c, c.Param("id"))
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, utils.NewHttpSuccess("Get data account site by id successfully", &utils.Payload{Data: resource.ToAccountSiteResource(data)}))
	return nil
}

func NewAccountSiteController() *AccountSiteController {
	return &AccountSiteController{
		AccountSiteService: account.NewAccountSiteService(),
	}
}
