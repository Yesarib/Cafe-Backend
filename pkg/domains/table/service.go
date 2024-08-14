package table

import (
	"cafe-backend/pkg/models"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	NewTable(c *gin.Context, table *models.Table) error
	GetTables(c *gin.Context) ([]*models.Table, error)
	GetTableByID(c *gin.Context, tableId uint) (models.Table, error)
	UpdateTable(c *gin.Context, tableId uint, updatedTable models.Table) error
	DeleteTable(c *gin.Context, tableId uint) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// DeleteTable implements Service.
func (s *service) DeleteTable(c *gin.Context, tableId uint) error {
	err := s.repository.DeleteTable(c, tableId)
	if err != nil {
		return err
	}

	return nil
}

// GetTableByID implements Service.
func (s *service) GetTableByID(c *gin.Context, tableId uint) (models.Table, error) {
	table, err := s.repository.GetTableByID(c, tableId)

	if err != nil {
		return models.Table{}, err
	}

	return table, nil
}

// GetTables implements Service.
func (s *service) GetTables(c *gin.Context) ([]*models.Table, error) {
	tables, err := s.repository.GetTables(c)
	if err != nil {
		return []*models.Table{}, err
	}

	return tables, nil
}

// NewTable implements Service.
func (s *service) NewTable(c *gin.Context, table *models.Table) error {
	err := s.repository.NewTable(c, table)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTable implements Service.
func (s *service) UpdateTable(c *gin.Context, tableId uint, updatedTable models.Table) error {
	table, err := s.repository.GetTableByID(c, tableId)

	if err != nil {
		return err
	}

	if table.ID == 0 {
		return errors.New("product not found")
	}

	table.TableNo = updatedTable.TableNo
	table.Status = updatedTable.Status

	return s.repository.UpdateTable(c, table)
}
