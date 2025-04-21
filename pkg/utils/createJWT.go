package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func CreateJWT(JWTSecret []byte, JWTExpiration int, userId uuid.UUID, ip string) (string, error) {
	expiration := time.Second * time.Duration(JWTExpiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userId":    userId,
		"ip":        ip,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenStr, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
