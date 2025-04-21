package authRepository

import (
	"github.com/google/uuid"
	"github.com/nedokyrill/auth-service/internal/models/refresh"
)

func (rep *AuthRepositoryImpl) DeleteRefresh(token uuid.UUID) error {
	return rep.db.Where("refreshToken = ?", token).Delete(&refresh.Refresh{}).Error
}
