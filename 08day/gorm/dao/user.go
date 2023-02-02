package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Password string `gorm:"column:password;type:varchar(20);" json:"password"`
}

func (user *User) TableName() string {
	return "user"
}
