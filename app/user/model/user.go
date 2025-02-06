package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"type:varchar(100);unique;not null"`
	PasswordHashed string `gorm:"type:varchar(100);not null"`
}

func (User) TableName() string {
	return "user"
}
