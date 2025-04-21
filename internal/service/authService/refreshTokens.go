package authService

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nedokyrill/auth-service/internal/models/refresh"
	"github.com/nedokyrill/auth-service/pkg/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (s *AuthService) RefreshTokens(c *gin.Context) {
	var payload refresh.RefreshRequest
	ip := c.ClientIP()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Refresh: error with parsing JSON")
		return
	}

	refreshToken, err := uuid.Parse(payload.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Refresh: error with parsing refresh token")
		return
	}

	oldRefresh, err := s.repoAuth.GetRefreshByToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Refresh: error with getting old refresh")
		return
	}

	err = s.repoAuth.DeleteRefresh(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Refresh: error with deleting old refresh")
		return
	}

	if oldRefresh.Ip != ip {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ip address has been changed"})
		log.Println("Refresh: ip address has been changed")
		//
		user, _ := s.repoUser.GetUserById(oldRefresh.UserId)
		log.Println("SEND EMAIL ABOUT CHANGED IP ADDRESS ON: ", user.Email)
		//
		return
	}

	if oldRefresh.CreatedAt.Unix()+oldRefresh.ExpiresIn < time.Now().Unix() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token is expired, please login again"})
		log.Println("Refresh: refresh token is expired")
		return
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	expiration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))

	token, err := utils.CreateJWT(secret, expiration, oldRefresh.UserId, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Refresh: error creating JWT")
		return
	}

	newRefreshToken := uuid.New()

	hash, err := utils.HashToken(newRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Refresh: error hashing token")
		return
	}

	newRefreshExpiration, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRATION"))

	newRefresh := refresh.Refresh{
		UserId:       oldRefresh.UserId,
		RefreshToken: hash,
		Ip:           ip,
		ExpiresIn:    int64(newRefreshExpiration),
	}

	err = s.repoAuth.CreateRefresh(newRefresh)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Refresh: error creating refresh session")
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token, "refresh_token": newRefreshToken})
	log.Println("Refresh: successfully refresh token")
}
