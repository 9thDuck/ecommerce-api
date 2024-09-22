package orders

import (
	"github.com/9thDuck/ecommerce-api.git/db"
	"gorm.io/gorm"
)

func CancelOrder(tx *gorm.DB, order *Order) error {
	if tx == nil {
		tx = db.Instance
	}

	if err := GetOrder(tx, order); err != nil {
		return err
	}

	order.Status = OrderStatusCancelled

	return Update(tx, order)
}

func GetOrder(tx *gorm.DB, order *Order) error {
	if tx == nil {
		tx = db.Instance
	}

	return Get(tx, order)
}
