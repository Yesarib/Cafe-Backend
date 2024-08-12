package orderitem

import (
	"cafe-backend/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	NewOrderItem(c *gin.Context, orderItem *models.OrderItem) error
	GetOrderItems(c *gin.Context) ([]*models.OrderItem, error)
	GetOrderItemById(c *gin.Context, orderItemId uint) (*models.OrderItem, error)
	UpdateOrderItem(c *gin.Context, orderItem models.OrderItem) error
	DeleteOrderItem(c *gin.Context, orderItemId uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// DeleteOrderItem implements Repository.
func (r *repository) DeleteOrderItem(c *gin.Context, orderItemId uint) error {
	result := r.db.Delete(&models.OrderItem{}, orderItemId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetOrderItemById implements Repository.
func (r *repository) GetOrderItemById(c *gin.Context, orderItemId uint) (*models.OrderItem, error) {
	var orderItem models.OrderItem

	result := r.db.First(&orderItem, orderItemId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &orderItem, nil
}

// GetOrderItems implements Repository.
func (r *repository) GetOrderItems(c *gin.Context) ([]*models.OrderItem, error) {
	var orderItems []*models.OrderItem

	result := r.db.Find(&orderItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderItems, nil
}

// NewOrderItem implements Repository.
func (r *repository) NewOrderItem(c *gin.Context, orderItem *models.OrderItem) error {
	return r.db.Create(&orderItem).Error
}

// UpdateOrderItem implements Repository.
func (r *repository) UpdateOrderItem(c *gin.Context, orderItem models.OrderItem) error {
	return r.db.Save(&orderItem).Error
}
