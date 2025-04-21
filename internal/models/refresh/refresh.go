package refresh

import (
	"github.com/google/uuid"
	"time"
)

type Refresh struct {
	Id           int       `json:"id" gorm:"column:id"`
	UserId       uuid.UUID `json:"userId" gorm:"column:userId"`
	RefreshToken string    `json:"refreshToken" gorm:"column:refreshToken"`
	Ip           string    `json:"ip" gorm:"column:ip"`
	ExpiresIn    int64     `json:"expiresIn" gorm:"column:expiresIn"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:createdAt"`
}

type LoginRequest struct {
	UserId string `json:"userId"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}
