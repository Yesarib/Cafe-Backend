package config

import (
	"cafe-backend/pkg/domains/auth"
	"cafe-backend/pkg/domains/employee"
	"cafe-backend/pkg/domains/product"
	"cafe-backend/pkg/initializers"
	"cafe-backend/pkg/models"
	"log"
)

type AppConfig struct {
	AuthService     auth.Service
	EmployeeService employee.Service
	ProductService  product.Service
	// Diğer servisler...
}

func NewConfig() *AppConfig {
	initializers.ConnectToDB()

	err := initializers.DB.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	authRepo := auth.NewRepository(initializers.DB)
	authService := auth.NewService(authRepo)

	employeeRepo := employee.NewRepository(initializers.DB)
	employeeService := employee.NewService(employeeRepo)

	productRepo := product.NewRepository(initializers.DB)
	productService := product.NewService(productRepo)

	return &AppConfig{
		AuthService:     authService,
		EmployeeService: employeeService,
		ProductService:  productService,
	}
}
