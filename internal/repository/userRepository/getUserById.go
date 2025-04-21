package userRepository

import (
	"github.com/google/uuid"
	"github.com/nedokyrill/auth-service/internal/models/user"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (rep *UserRepositoryImpl) GetUserById(id uuid.UUID) (*user.User, error) {
	var userWithEmail user.User
	err := rep.db.Where("id = ?", id).First(&userWithEmail).Error
	if err != nil {
		return nil, err
	}
	return &userWithEmail, nil
}
