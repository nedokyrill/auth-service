package authRepository

import (
	"github.com/google/uuid"
	"github.com/nedokyrill/auth-service/internal/models/refresh"
)

func (rep *AuthRepositoryImpl) GetRefreshByToken(token uuid.UUID) (*refresh.Refresh, error) {
	var currRefresh refresh.Refresh
	err := rep.db.Table("refreshes").Where("refreshToken = ?", token).First(&currRefresh).Error
	if err != nil {
		return nil, err
	}
	return &currRefresh, nil
}
