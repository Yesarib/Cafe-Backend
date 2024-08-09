package main

import (
	"cafe-backend/pkg/initializers"
	"cafe-backend/pkg/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Category{})
	initializers.DB.AutoMigrate(&models.Employee{})
	initializers.DB.AutoMigrate(&models.Order{})
	initializers.DB.AutoMigrate(&models.OrderItem{})
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.Table{})
}
