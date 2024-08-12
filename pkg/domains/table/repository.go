package table

import (
	"cafe-backend/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	NewTable(c *gin.Context, table *models.Table) error
	GetTables(c *gin.Context) ([]*models.Table, error)
	GetTableByID(c *gin.Context, tableId uint) (models.Table, error)
	UpdateTable(c *gin.Context, table models.Table) error
	DeleteTable(c *gin.Context, tableId uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// DeleteTable implements Repository.
func (r *repository) DeleteTable(c *gin.Context, tableId uint) error {
	result := r.db.Delete(&models.Table{}, tableId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetTableByID implements Repository.
func (r *repository) GetTableByID(c *gin.Context, tableId uint) (models.Table, error) {
	var table models.Table

	result := r.db.First(&table, tableId)
	if result.Error != nil {
		return models.Table{}, result.Error
	}

	return table, nil
}

// GetTables implements Repository.
func (r *repository) GetTables(c *gin.Context) ([]*models.Table, error) {
	var tables []*models.Table

	result := r.db.Find(&tables)

	if result.Error != nil {
		return nil, result.Error
	}

	return tables, nil
}

// NewTable implements Repository.
func (r *repository) NewTable(c *gin.Context, table *models.Table) error {
	return r.db.Create(&table).Error
}

// UpdateTable implements Repository.
func (r *repository) UpdateTable(c *gin.Context, table models.Table) error {
	return r.db.Save(table).Error
}
