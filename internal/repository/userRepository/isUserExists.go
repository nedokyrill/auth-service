package userRepository

import (
	"github.com/google/uuid"
	"github.com/nedokyrill/auth-service/internal/models/user"
)

func (rep *UserRepositoryImpl) IsUserExists(id uuid.UUID) bool {
	var existingUser user.User
	if err := rep.db.Where("id = ?", id).First(&existingUser).Error; err != nil {
		return false
	} else {
		return true
	}
}
