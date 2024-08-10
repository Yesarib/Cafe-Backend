package product

import (
	"cafe-backend/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	NewProduct(c *gin.Context, product *models.Product) error
	GetProducts(c *gin.Context) ([]*models.Product, error)
	GetProductByID(c *gin.Context, productId uint) (models.Product, error)
	UpdateProduct(c *gin.Context, product models.Product) error
	DeleteProduct(c *gin.Context, productId uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// DeleteProduct implements Repository.
func (r *repository) DeleteProduct(c *gin.Context, productId uint) error {
	result := r.db.Delete(&models.Product{}, productId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetProductByID implements Repository.
func (r *repository) GetProductByID(c *gin.Context, productId uint) (models.Product, error) {
	var product models.Product

	result := r.db.First(&product, productId)
	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return product, nil
}

// GetProducts implements Repository.
func (r *repository) GetProducts(c *gin.Context) ([]*models.Product, error) {
	var products []*models.Product

	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

// NewProduct implements Repository.
func (r *repository) NewProduct(c *gin.Context, product *models.Product) error {
	return r.db.Create(&product).Error
}

// UpdateProduct implements Repository.
func (r *repository) UpdateProduct(c *gin.Context, product models.Product) error {
	return r.db.Save(product).Error
}
