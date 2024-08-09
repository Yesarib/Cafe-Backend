package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int16
	Price     float64
}
