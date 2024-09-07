package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username       string    `json:"username" gorm:"unique"`
	Email          string    `json:"email" gorm:"unique"`
	Password       string    `json:"-" gorm:"-"`
	HashedPassword string    `json:"-"`
	Verified       bool      `json:"verified"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt      time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}
