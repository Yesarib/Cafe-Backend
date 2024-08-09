package models

import "gorm.io/gorm"

type TableStatus string

const (
	Available TableStatus = "Available"
	Occupied  TableStatus = "Occupied"
)

type Table struct {
	gorm.Model
	TableNo string `gorm:"unique"`
	Status  TableStatus
}
