package service

import "github.com/gin-gonic/gin"

type AuthService interface {
	Login(c *gin.Context)
	RefreshTokens(c *gin.Context)
}
