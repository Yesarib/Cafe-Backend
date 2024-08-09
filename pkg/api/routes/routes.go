package routes

import (
	"cafe-backend/pkg/domains/auth"
	"cafe-backend/pkg/domains/employee"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authService auth.Service, employeeService employee.Service) {
	AuthRoutes(r, authService)
	EmployeeRoutes(r, employeeService)
}
