package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username       string    `json:"username" gorm:"unique;not null;type:varchar(20)"`
	Role           int       `json:"role" gorm:"type:int;not null"`
	Email          string    `json:"email" gorm:"unique;not null;type:varchar(50)"`
	Password       string    `json:"-" gorm:"-"`
	HashedPassword string    `json:"-"`
	BlackListed    bool      `json:"-" gorm:"default:false"`
	Verified       bool      `json:"verified" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"not null"`
}
