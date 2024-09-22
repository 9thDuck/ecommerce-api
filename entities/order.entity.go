package entities

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID `json:"user_id" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Status    int       `json:"status" gorm:"type:int;not null"`
	User      User      `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Product   Product   `json:"-" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

type OrderItem struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	OrderID   uint      `json:"order_id" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Order     Order     `json:"-" gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Product   Product   `json:"-" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
