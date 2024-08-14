package category

import (
	"cafe-backend/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	NewCategory(c *gin.Context, category *models.Category) error
	GetCategories(c *gin.Context) ([]*models.Category, error)
	GetCategoryById(c *gin.Context, categoryId uint) (*models.Category, error)
	UpdateCategory(c *gin.Context, category models.Category) error
	DeleteCategory(c *gin.Context, categoryId uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// DeleteCategory implements Repository.
func (r *repository) DeleteCategory(c *gin.Context, categoryId uint) error {
	result := r.db.Delete(&models.Category{}, categoryId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetCategories implements Repository.
func (r *repository) GetCategories(c *gin.Context) ([]*models.Category, error) {
	var categories []*models.Category

	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

// GetCategoryById implements Repository.
func (r *repository) GetCategoryById(c *gin.Context, categoryId uint) (*models.Category, error) {
	var category models.Category

	result := r.db.First(&category, categoryId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

// NewCategory implements Repository.
func (r *repository) NewCategory(c *gin.Context, category *models.Category) error {
	return r.db.Create(&category).Error
}

// UpdateCategory implements Repository.
func (r *repository) UpdateCategory(c *gin.Context, category models.Category) error {
	return r.db.Save(category).Error
}
