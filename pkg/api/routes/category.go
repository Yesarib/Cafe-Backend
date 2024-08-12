package routes

import (
	"cafe-backend/pkg/api/handlers"
	"cafe-backend/pkg/domains/category"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, service category.Service) {
	handler := handlers.NewCategoryHandler(service)

	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.POST("/", handler.NewCategory)
		categoryRoutes.GET("/", handler.GetCategories)
		categoryRoutes.GET("/", handler.GetCategoryById)
		categoryRoutes.PUT("/:id", handler.UpdateCategory)
		categoryRoutes.DELETE("/:id", handler.DeleteCategory)
	}
}
