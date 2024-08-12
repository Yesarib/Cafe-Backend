package main

import (
	"cafe-backend/pkg/api/routes"
	"cafe-backend/pkg/config"
	"cafe-backend/pkg/initializers"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	cfg := config.NewConfig()

	r := gin.Default()

	routes.RegisterRoutes(r,
		cfg.AuthService,
		cfg.EmployeeService,
		cfg.ProductService,
		cfg.TableService,
		cfg.CategoryService,
		cfg.OrderItemService,
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
