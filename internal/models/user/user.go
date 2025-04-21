package user

import "github.com/google/uuid"

type User struct {
	Id    uuid.UUID `json:"id" gorm:"column:id"`
	Name  string    `json:"name" gorm:"column:name"`
	Email string    `json:"email" gorm:"column:email"`
}
