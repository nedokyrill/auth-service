package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CompareTokens(hashToken string, token uuid.UUID) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashToken), []byte(token.String()))
	return err == nil
}
