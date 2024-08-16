package routes

import (
	"cafe-backend/pkg/api/handlers"
	"cafe-backend/pkg/domains/category"
	"cafe-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, service category.Service) {
	handler := handlers.NewCategoryHandler(service)

	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.POST("/", middleware.AuthMiddleware(), handler.NewCategory)
		categoryRoutes.GET("/", handler.GetCategories)
		categoryRoutes.GET("/:id", handler.GetCategoryById)
		categoryRoutes.PUT("/:id", middleware.AuthMiddleware(), handler.UpdateCategory)
		categoryRoutes.DELETE("/:id", middleware.AuthMiddleware(), handler.DeleteCategory)
	}
}
