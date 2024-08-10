package product

import (
	"cafe-backend/pkg/models"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	NewProduct(c *gin.Context, product *models.Product) error
	GetProducts(c *gin.Context) ([]*models.Product, error)
	GetProductByID(c *gin.Context, productId uint) (models.Product, error)
	UpdateProduct(c *gin.Context, productId uint, product models.Product) error
	DeleteProduct(c *gin.Context, productId uint) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) NewProduct(c *gin.Context, product *models.Product) error {
	err := s.repository.NewProduct(c, product)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetProducts(c *gin.Context) ([]*models.Product, error) {
	products, err := s.repository.GetProducts(c)

	if err != nil {
		return []*models.Product{}, err
	}

	return products, nil
}

func (s *service) GetProductByID(c *gin.Context, productId uint) (models.Product, error) {
	product, err := s.repository.GetProductByID(c, productId)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *service) UpdateProduct(c *gin.Context, productId uint, updatedProduct models.Product) error {
	product, err := s.repository.GetProductByID(c, productId)
	if err != nil {
		return err
	}

	if product.ID == 0 {
		return errors.New("product not found")
	}

	product.Name = updatedProduct.Name
	product.Description = updatedProduct.Description
	product.Price = updatedProduct.Price
	product.CategoryID = updatedProduct.CategoryID

	return s.repository.UpdateProduct(c, product)
}

func (s *service) DeleteProduct(c *gin.Context, productId uint) error {
	err := s.repository.DeleteProduct(c, productId)
	if err != nil {
		return err
	}

	return nil
}
