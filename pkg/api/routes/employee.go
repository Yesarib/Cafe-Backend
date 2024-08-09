package routes

import (
	"cafe-backend/pkg/api/handlers"
	"cafe-backend/pkg/domains/employee"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(r *gin.Engine, service employee.Service) {
	handler := handlers.NewEmployeeHandler(service)

	employeeRoutes := r.Group("/employee")
	{
		employeeRoutes.POST("/", handler.NewEmployee)
		employeeRoutes.GET("/", handler.GetEmployee)
		employeeRoutes.GET("/:id", handler.GetEmployeeByID)
		employeeRoutes.PUT("/:id", handler.UpdateEmployee)
		employeeRoutes.DELETE("/:id", handler.DeleteEmployee)
	}

}
