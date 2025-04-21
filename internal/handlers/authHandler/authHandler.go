package authHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/nedokyrill/auth-service/internal/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", h.service.Login)
		authRouter.POST("/refresh-tokens", h.service.RefreshTokens)
	}

}
