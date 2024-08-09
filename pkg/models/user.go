package models

import (
	"cafe-backend/pkg/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Role     string
}

func (u *User) PassHash() error {
	pass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pass
	return nil
}
