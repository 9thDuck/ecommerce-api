package entities

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	ProductID uint      `json:"product_id" gorm:"not null;uniqueIndex:idx_product_user"`
	Product   Product   `json:"-" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null;uniqueIndex:idx_product_user"`
	User      User      `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
