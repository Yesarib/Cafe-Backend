package routes

import (
	"cafe-backend/pkg/api/handlers"
	"cafe-backend/pkg/domains/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, service auth.Service) {
	handler := handlers.NewHandler(service)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", handler.Register)
		authRoutes.POST("/login", handler.Login)
	}
}
