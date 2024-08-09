package employee

import (
	"cafe-backend/pkg/dto"
	"cafe-backend/pkg/models"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	NewEmployee(c *gin.Context, employee *models.Employee) error
	GetEmployee(c *gin.Context) ([]*models.Employee, error)
	GetEmployeeByID(c *gin.Context, employeeID string) (models.Employee, error)
	UpdateEmployee(c *gin.Context, employeeId uint, employee dto.UpdateEmployeeRequestDTO) error
	DeleteEmployee(c *gin.Context, employeeId uint) error
}

type service struct {
	repository Repository
}

// DeleteEmployee implements Service.
func (s *service) DeleteEmployee(c *gin.Context, employeeId uint) error {
	err := s.repository.DeleteEmployee(c, employeeId)

	if err != nil {
		return err
	}

	return nil
}

// GetEmployee implements Service.
func (s *service) GetEmployee(c *gin.Context) ([]*models.Employee, error) {
	employees, err := s.repository.GetEmployee(c)

	if err != nil {
		return []*models.Employee{}, err
	}

	return employees, nil
}

// GetEmployeeByID implements Service.
func (s *service) GetEmployeeByID(c *gin.Context, employeeID string) (models.Employee, error) {
	employee, err := s.repository.GetEmployeeByID(c, employeeID)
	if err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

// NewEmployee implements Service.
func (s *service) NewEmployee(c *gin.Context, employee *models.Employee) error {
	// Mevcut kullanıcıyı kontrol et
	existingEmployee, err := s.repository.GetEmployeeByUserName(c, employee.Username)
	if err == nil && &existingEmployee != nil {
		return errors.New("employee already exists")
	}

	// Şifreyi hashle
	err = employee.PassHash()
	if err != nil {
		return errors.New("failed when trying to hash password")
	}

	// Yeni kullanıcıyı oluştur
	err = s.repository.NewEmployee(c, employee)
	if err != nil {
		return err
	}

	return nil
}

// UpdateEmployee implements Service.
func (s *service) UpdateEmployee(c *gin.Context, employeeId uint, employee dto.UpdateEmployeeRequestDTO) error {
	err := s.repository.UpdateEmployee(c, employeeId, employee)
	if err != nil {
		// Eğer bir hata varsa, hatayı döndür
		return err
	}

	// Başarılı bir güncelleme durumunda nil hatasını döndür
	return nil
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
