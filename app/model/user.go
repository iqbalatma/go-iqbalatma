package model

type User struct {
	BaseModel `gorm:"embedded"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Email     string `json:"email" gorm:"column:email;unique"`
	Password  string `json:"password" gorm:"column:password;not null"`
}

func (user *User) GetSubjectKey() string {
	return user.ID.String()
}
