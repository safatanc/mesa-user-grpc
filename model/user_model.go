package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid; primaryKey; default:gen_random_uuid()"`
	Username  string    `gorm:"unique" validate:"required, unique, min=3, max=20"`
	FullName  string    `validate:"required, min=3, max=50"`
	Email     string    `validate:"required, email"`
	Phone     string    `validate:"required, e164"`
	Password  string    `validate:"required, sha256"`
	CreatedAt int64     `gorm:"autoCreateTime" validate:"datetime"`
	UpdatedAt int64     `gorm:"autoUpdateTime" validate:"datetime"`
}
