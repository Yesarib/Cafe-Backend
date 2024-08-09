package auth

import (
	"cafe-backend/pkg/dto"
	"cafe-backend/pkg/middleware"
	"cafe-backend/pkg/models"
	"cafe-backend/pkg/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Register(c *gin.Context, user *models.User) (models.User, error)
	Login(c *gin.Context, req *dto.LoginReq) (dto.LoginDTO, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Register(c *gin.Context, user *models.User) (models.User, error) {
	_, err := s.repository.GetUserByUserName(c, user.Username)

	if err == nil {
		return models.User{}, errors.New(`this user already exists`)
	}
	err = user.PassHash()
	if err != nil {
		return models.User{}, errors.New(`failed when try to hashing password`)
	}
	err = s.repository.Save(c, user)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (s *service) Login(c *gin.Context, req *dto.LoginReq) (dto.LoginDTO, error) {
	var loginDto dto.LoginDTO

	user, err := s.repository.GetUserByUserName(c, req.UserName)
	if err != nil {
		return loginDto, errors.New("failed login")
	}

	checkPassword := utils.PasswordControl(user.Password, req.Password)
	if !checkPassword {
		return loginDto, errors.New("email or password doesnt match")
	}

	token, err := middleware.SignAccessToken(string(user.ID))

	if err != nil {
		return loginDto, errors.New("there is an error when get access token")
	}

	loginDto = dto.LoginDTO{
		UserID:   user.ID,
		UserName: user.Username,
		Token:    token,
	}

	return loginDto, nil
}
