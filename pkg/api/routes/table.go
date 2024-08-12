package routes

import (
	"cafe-backend/pkg/api/handlers"
	"cafe-backend/pkg/domains/table"

	"github.com/gin-gonic/gin"
)

func TableRoutes(r *gin.Engine, service table.Service) {
	handler := handlers.NewTableHandler(service)

	tableRoutes := r.Group("/table")
	{
		tableRoutes.POST("/", handler.NewTable)
		tableRoutes.GET("/", handler.GetTables)
		tableRoutes.GET("/:id", handler.GetTableByID)
		tableRoutes.PUT("/:id", handler.UpdateTable)
		tableRoutes.DELETE("/:id", handler.DeleteTable)
	}
}
