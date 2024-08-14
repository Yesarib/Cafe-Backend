package category

import (
	"cafe-backend/pkg/models"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	NewCategory(c *gin.Context, category *models.Category) error
	GetCategories(c *gin.Context) ([]*models.Category, error)
	GetCategoryById(c *gin.Context, categoryId uint) (*models.Category, error)
	UpdateCategory(c *gin.Context, categoryId uint, updatedCategory models.Category) error
	DeleteCategory(c *gin.Context, categoryId uint) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// DeleteCategory implements Service.
func (s *service) DeleteCategory(c *gin.Context, categoryId uint) error {
	err := s.repository.DeleteCategory(c, categoryId)
	if err != nil {
		return err
	}

	return nil
}

// GetCategories implements Service.
func (s *service) GetCategories(c *gin.Context) ([]*models.Category, error) {
	categories, err := s.repository.GetCategories(c)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryById implements Service.
func (s *service) GetCategoryById(c *gin.Context, categoryId uint) (*models.Category, error) {
	category, err := s.repository.GetCategoryById(c, categoryId)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// NewCategory implements Service.
func (s *service) NewCategory(c *gin.Context, category *models.Category) error {
	err := s.repository.NewCategory(c, category)

	if err != nil {
		return err
	}

	return nil
}

// UpdateCategory implements Service.
func (s *service) UpdateCategory(c *gin.Context, categoryId uint, updatedCategory models.Category) error {
	category, err := s.repository.GetCategoryById(c, categoryId)

	if err != nil {
		return err
	}

	if category.ID == 0 {
		return errors.New("category not found")
	}

	category.Name = updatedCategory.Name
	category.Description = updatedCategory.Description

	return s.repository.UpdateCategory(c, *category)
}
