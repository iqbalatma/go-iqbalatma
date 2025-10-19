package dto

type CreateAccountSiteRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Url         string `json:"url" binding:"required,min=2"`
}
