package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashToken(token uuid.UUID) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token.String()), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
