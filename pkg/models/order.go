package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	CreditCard PaymentMethod = "CreditCard"
	Cash       PaymentMethod = "Cash"
)

// var validPaymentMethods = map[PaymentMethod]bool{
// 	CreditCard: true,
// 	Cash:       true,
// }

type Order struct {
	gorm.Model
	EmployeeID    uint
	TotalAmount   float64
	Status        string
	OrderDate     time.Time
	TableID       uint
	PaymentMethod PaymentMethod
}

// func (o *Order) BeforeSave(tx *gorm.DB) (err error) {
// 	if !validPaymentMethods[o.PaymentMethod] {
// 		return errors.New("invalid payment method")
// 	}
// 	return nil
// }
