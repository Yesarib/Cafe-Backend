package models

import (
	"cafe-backend/pkg/utils"
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Username    string
	Password    string
	Email       string
	FirstName   string
	LastName    string
	Position    string
	HourlyPrice float64
	HiredAt     time.Time
}

func (u *Employee) PassHash() error {
	pass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pass
	return nil
}
