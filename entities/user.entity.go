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
	Address        string    `json:"address" gorm:"type:varchar(150)"`
	Password       string    `json:"-" gorm:"-"`
	HashedPassword string    `json:"-"`
	Banned         bool      `json:"-" gorm:"default:false"`
	Verified       bool      `json:"verified" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"not null"`
}

type Session struct {
	ID           uint      `json:"id,omitempty" gorm:"type:uint;primaryKey;autoIncrement"`
	UserID       uuid.UUID `json:"user_id,omitempty" gorm:"type:uuid;not null"`
	RefreshToken string    `json:"refresh_token,omitempty" gorm:"type:text;not null"`
	User         User      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
	LoggedOut    bool      `json:"logged_out" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null"`
}
