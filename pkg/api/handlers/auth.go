package handlers

import (
	"cafe-backend/pkg/domains/auth"
	"cafe-backend/pkg/dto"
	"cafe-backend/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service auth.Service
}

func NewHandler(service auth.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registeredUser, err := h.Service.Register(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, registeredUser)
}

func (h *Handler) Login(c *gin.Context) {
	var loginReq dto.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginDto, err := h.Service.Login(c, &loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginDto)
}
