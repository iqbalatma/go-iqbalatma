package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	BaseUUID  `gorm:"embedded"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;autoUpdateTime"`
}

func (baseModel *BaseModel) BeforeCreate(tx *gorm.DB) error {
	baseModel.GenerateUUID()
	return nil
}
