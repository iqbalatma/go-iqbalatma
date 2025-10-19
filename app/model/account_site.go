package model

type AccountSite struct {
	BaseModel   `gorm:"embedded"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	Url         string  `json:"url" gorm:"column:url"`
	Icon        *string `json:"icon" gorm:"column:icon"`
}
