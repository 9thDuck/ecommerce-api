package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null;type:varchar(100)" validate:"required,min=3,max=100"`
	Description string    `json:"description" gorm:"not null;type:varchar(1000)" validate:"required,min=10,max=1000"`
	Price       float64   `json:"price" gorm:"not null;type:decimal(10,2)" validate:"required,min=0.01"`
	Stock       int       `json:"stock" gorm:"not null" validate:"required,min=0"`
	Image       string    `json:"image" gorm:"type:varchar(255)" validate:"omitempty,url"`
	Hidden      bool      `json:"hidden" gorm:"not null;default:false"`
	CategoryID  uint      `json:"category_id" gorm:"not null" validate:"required"`
	Category    Category  `json:"-" gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedBy   uuid.UUID `json:"created_by" gorm:"type:uuid;not null"`
	User        User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null"`
}
