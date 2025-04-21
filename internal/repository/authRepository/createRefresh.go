package authRepository

import (
	"github.com/nedokyrill/auth-service/internal/models/refresh"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

func (rep *AuthRepositoryImpl) CreateRefresh(payload refresh.Refresh) error {
	return rep.db.Create(&payload).Error
}
