package routes

import (
	"cafe-backend/pkg/domains/auth"
	"cafe-backend/pkg/domains/employee"
	"cafe-backend/pkg/domains/product"
	"cafe-backend/pkg/domains/table"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	authService auth.Service,
	employeeService employee.Service,
	productService product.Service,
	tableService table.Service,
) {
	AuthRoutes(r, authService)
	EmployeeRoutes(r, employeeService)
	ProductRoutes(r, productService)
	TableRoutes(r, tableService)
}
