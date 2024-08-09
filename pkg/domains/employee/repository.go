package employee

import (
	"cafe-backend/pkg/dto"
	"cafe-backend/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	NewEmployee(c *gin.Context, employee *models.Employee) error
	GetEmployee(c *gin.Context) ([]*models.Employee, error)
	GetEmployeeByID(c *gin.Context, employeeID string) (models.Employee, error)
	UpdateEmployee(c *gin.Context, employeeId uint, employee dto.UpdateEmployeeRequestDTO) error
	DeleteEmployee(c *gin.Context, employeeId uint) error
	GetEmployeeByUserName(c *gin.Context, userName string) (models.Employee, error)
}

type repository struct {
	db *gorm.DB
}

// GetEmployeeByUserName implements Repository.
func (r *repository) GetEmployeeByUserName(c *gin.Context, userName string) (models.Employee, error) {
	var employee models.Employee

	err := r.db.Table("employees").Where("username = ?", userName).First(&employee).Error

	return employee, err
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// DeleteEmployee implements Repository.
func (r *repository) DeleteEmployee(c *gin.Context, employeeId uint) error {
	result := r.db.Delete(&models.Employee{}, employeeId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetEmployee implements Repository.
func (r *repository) GetEmployee(c *gin.Context) ([]*models.Employee, error) {
	var employees []*models.Employee

	result := r.db.Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}

	return employees, nil
}

// GetEmployeeByID implements Repository.
func (r *repository) GetEmployeeByID(c *gin.Context, employeeID string) (models.Employee, error) {
	var employee models.Employee

	result := r.db.First(&employee, employeeID)

	if result.Error != nil {
		return models.Employee{}, result.Error
	}

	return employee, nil
}

// NewEmployee implements Repository.
func (r *repository) NewEmployee(c *gin.Context, employee *models.Employee) error {
	return r.db.Create(&employee).Error
}

// UpdateEmployee implements Repository.
func (r *repository) UpdateEmployee(c *gin.Context, employeeId uint, employee dto.UpdateEmployeeRequestDTO) error {
	var existingEmployee models.Employee
	if err := r.db.First(&existingEmployee, employeeId).Error; err != nil {
		return err
	}

	existingEmployee.Username = employee.Username
	existingEmployee.Email = employee.Email
	existingEmployee.FirstName = employee.FirstName
	existingEmployee.LastName = employee.LastName
	existingEmployee.Position = employee.Position
	existingEmployee.HourlyPrice = employee.HourlyPrice

	if err := r.db.Save(&existingEmployee).Error; err != nil {
		return err
	}
	return nil
}
