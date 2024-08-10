package routes

import (
	"cafe-backend/pkg/api/handlers"
	"cafe-backend/pkg/domains/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, service product.Service) {
	handler := handlers.NewProductHandler(service)

	productRoutes := r.Group("/product")
	{
		productRoutes.POST("/", handler.NewProduct)
		productRoutes.GET("/", handler.GetProducts)
		productRoutes.GET("/:id", handler.GetProductByID)
		productRoutes.PUT("/:id", handler.UpdateProduct)
		productRoutes.DELETE("/:id", handler.DeleteProduct)
	}
}
