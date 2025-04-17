package refreshSession

import (
	"github.com/google/uuid"
	"time"
)

type RefreshSession struct {
	Id           uuid.UUID `json:"id"`
	UserId       uuid.UUID `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	Ip           string    `json:"ip"`
	ExpiresIn    int64     `json:"expires"`
	CreatedAt    time.Time `json:"created"`
}
