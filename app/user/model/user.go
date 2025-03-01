package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHashed string `gorm:"type:varchar(100);not null"`
}

func (User) TableName() string {
	return "user"
}

func Creat(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}
