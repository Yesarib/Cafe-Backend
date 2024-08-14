package handlers

import (
	orderitem "cafe-backend/pkg/domains/orderItem"
	"cafe-backend/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderItemHandler struct {
	Service orderitem.Service
}

func NewOrderItemHandler(service orderitem.Service) *OrderItemHandler {
	return &OrderItemHandler{
		Service: service,
	}
}

func (h *OrderItemHandler) NewOrderItem(c *gin.Context) {
	var orderItem models.OrderItem

	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.NewOrderItem(c, &orderItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order Item created"})
}

func (h *OrderItemHandler) GetOrderItems(c *gin.Context) {
	orderItems, err := h.Service.GetOrderItems(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderItems)
}

func (h *OrderItemHandler) GetOrderItemById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	orderItem, err := h.Service.GetOrderItemById(c, uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func (h *OrderItemHandler) UpdateOrderItem(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var orderItem models.OrderItem

	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Service.UpdateOrderItem(c, uint(id), orderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order Item successfully updated"})

}

func (h *OrderItemHandler) DeleteOrderItem(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.Service.DeleteOrderItem(c, uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete table"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order Item deleted successfully"})
}
