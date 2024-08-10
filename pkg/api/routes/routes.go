package routes

import (
	"cafe-backend/pkg/domains/auth"
	"cafe-backend/pkg/domains/employee"
	"cafe-backend/pkg/domains/product"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	authService auth.Service,
	employeeService employee.Service,
	productService product.Service,
) {
	AuthRoutes(r, authService)
	EmployeeRoutes(r, employeeService)
	ProductRoutes(r, productService)
}
