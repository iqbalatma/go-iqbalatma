package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	BaseUUID
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.BaseUUID.GenerateUUID()
	return nil
}
