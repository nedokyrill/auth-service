package authService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/nedokyrill/auth-service/internal/models/refresh"
	"github.com/nedokyrill/auth-service/internal/repository"
	"github.com/nedokyrill/auth-service/pkg/utils"
	"log"
	"net/http"
	"os"
	"strconv"
)

type AuthService struct {
	repoAuth repository.AuthRepository
	repoUser repository.UserRepository
}

func NewAuthService(repoAuth repository.AuthRepository, repoUser repository.UserRepository) *AuthService {
	return &AuthService{
		repoAuth: repoAuth,
		repoUser: repoUser,
	}
}

func (s *AuthService) Login(c *gin.Context) {
	var payload refresh.LoginRequest
	ip := c.ClientIP()

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Login: error with parsing JSON")
		return
	}

	userId, err := uuid.Parse(payload.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Login: error with parsing user id")
		return
	}

	if !s.repoUser.IsUserExists(userId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		log.Println("Login: user not found")
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Println("Login: error loading .env file")
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	expiration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))

	token, err := utils.CreateJWT(secret, expiration, userId, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Login: error creating JWT")
		return
	}

	newRefreshToken := uuid.New()

	hash, err := utils.HashToken(newRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Login: error hashing token")
		return
	}

	newRefreshExpiration, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRATION"))

	newRefresh := refresh.Refresh{
		UserId:       userId,
		RefreshToken: hash,
		Ip:           ip,
		ExpiresIn:    int64(newRefreshExpiration),
	}

	err = s.repoAuth.CreateRefresh(newRefresh)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Login: error creating refresh session")
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token, "refresh_token": newRefreshToken})
	log.Println("Login: successfully logged in")

}
