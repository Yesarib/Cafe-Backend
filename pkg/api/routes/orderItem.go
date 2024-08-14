package routes

import (
	"cafe-backend/pkg/api/handlers"
	orderitem "cafe-backend/pkg/domains/orderItem"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(r *gin.Engine, service orderitem.Service) {
	handler := handlers.NewOrderItemHandler(service)

	orderItemRoutes := r.Group("/orderItem")
	{
		orderItemRoutes.POST("/", handler.NewOrderItem)
		orderItemRoutes.GET("/", handler.GetOrderItems)
		orderItemRoutes.GET("/:id", handler.GetOrderItemById)
		orderItemRoutes.PUT("/:id", handler.UpdateOrderItem)
		orderItemRoutes.DELETE("/:id", handler.DeleteOrderItem)
	}
}
