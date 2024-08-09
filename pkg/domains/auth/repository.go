package auth

import (
	"cafe-backend/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository interface {
	Save(c *gin.Context, user *models.User) error
	GetUserByID(c *gin.Context, userId uint) (models.User, error)
	GetUserByUserName(c *gin.Context, userName string) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// GetUserByID implements Repository.
func (r *repository) GetUserByID(c *gin.Context, userId uint) (models.User, error) {
	var user models.User

	err := r.db.Table("users").Where("id = ?", userId).First(&user).Error

	return user, err
}

// Save implements Repository.
func (r *repository) Save(c *gin.Context, user *models.User) error {
	return r.db.Create(&user).Error
}

// GetUserByUserName implements Repository.
func (r *repository) GetUserByUserName(c *gin.Context, userName string) (models.User, error) {
	var user models.User

	err := r.db.Table("users").Where("username = ?", userName).First(&user).Error

	return user, err
}
