package repository

import (
	"github.com/google/uuid"
	"github.com/nedokyrill/auth-service/internal/models/refresh"
	"github.com/nedokyrill/auth-service/internal/models/user"
)

type AuthRepository interface {
	GetRefreshByToken(token uuid.UUID) (*refresh.Refresh, error)
	DeleteRefresh(token uuid.UUID) error
	CreateRefresh(refreshSession refresh.Refresh) error
}

type UserRepository interface {
	GetUserById(id uuid.UUID) (*user.User, error)
	IsUserExists(id uuid.UUID) bool
}
