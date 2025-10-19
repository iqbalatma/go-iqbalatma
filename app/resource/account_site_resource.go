package resource

import (
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/utils"
)

type AccountSiteResource struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Url         string  `json:"url"`
	Icon        *string `json:"icon"`
	CreatedAt   string  `json:"created_at"`
}

func ToAccountSiteResource(accountSite *model.AccountSite) *AccountSiteResource {
	return &AccountSiteResource{
		Id:          accountSite.ID.String(),
		Name:        accountSite.Name,
		Description: accountSite.Description,
		Url:         accountSite.Url,
		Icon:        accountSite.Icon,
		CreatedAt:   utils.ToDateTime(&accountSite.CreatedAt),
	}
}

func ToAccountSiteResourceCollection(accountSites []model.AccountSite) []*AccountSiteResource {
	accountSiteResourceCollection := make([]*AccountSiteResource, len(accountSites))
	for i, accountSite := range accountSites {
		accountSiteResourceCollection[i] = ToAccountSiteResource(&accountSite)
	}

	return accountSiteResourceCollection
}
