package entities

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"not null;type:varchar(50);unique" validate:"required,min=3,max=50"`
	Description string `json:"description" gorm:"not null;type:varchar(500)" validate:"required,min=10,max=500"`
	Hidden      bool      `json:"hidden" gorm:"not null;default:false"`
	CreatedBy   uuid.UUID `json:"created_by" gorm:"not null;foreignKey:CreatedBy"`
	User        User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null"`
}