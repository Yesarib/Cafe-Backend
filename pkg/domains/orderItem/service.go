package orderitem

import (
	"cafe-backend/pkg/models"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	NewOrderItem(c *gin.Context, orderItem *models.OrderItem) error
	GetOrderItems(c *gin.Context) ([]*models.OrderItem, error)
	GetOrderItemById(c *gin.Context, orderItemId uint) (*models.OrderItem, error)
	UpdateOrderItem(c *gin.Context, oderItemId uint, orderItem models.OrderItem) error
	DeleteOrderItem(c *gin.Context, orderItemId uint) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// DeleteOrderItem implements Service.
func (s *service) DeleteOrderItem(c *gin.Context, orderItemId uint) error {
	err := s.repository.DeleteOrderItem(c, orderItemId)

	if err != nil {
		return err
	}

	return nil
}

// GetOrderItemById implements Service.
func (s *service) GetOrderItemById(c *gin.Context, orderItemId uint) (*models.OrderItem, error) {
	orderItem, err := s.repository.GetOrderItemById(c, orderItemId)

	if err != nil {
		return nil, err
	}

	return orderItem, nil
}

// GetOrderItems implements Service.
func (s *service) GetOrderItems(c *gin.Context) ([]*models.OrderItem, error) {
	orderItems, err := s.repository.GetOrderItems(c)

	if err != nil {
		return nil, err
	}

	return orderItems, nil
}

// NewOrderItem implements Service.
func (s *service) NewOrderItem(c *gin.Context, orderItem *models.OrderItem) error {
	err := s.repository.NewOrderItem(c, orderItem)

	if err != nil {
		return err
	}

	return nil
}

// UpdateOrderItem implements Service.
func (s *service) UpdateOrderItem(c *gin.Context, oderItemId uint, updatedOrderItem models.OrderItem) error {
	orderItem, err := s.repository.GetOrderItemById(c, oderItemId)
	if err != nil {
		return err
	}

	if orderItem.ID == 0 {
		return errors.New("category not found")
	}

	orderItem.OrderID = updatedOrderItem.OrderID
	orderItem.ProductID = updatedOrderItem.ProductID
	orderItem.Quantity = updatedOrderItem.Quantity
	orderItem.Price = updatedOrderItem.Price

	return s.repository.UpdateOrderItem(c, *orderItem)
}
