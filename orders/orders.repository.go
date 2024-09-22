package orders

import (
	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order entities.Order
type OrderItem entities.OrderItem

const (
	OrderStatusPending        = 1
	OrderStatusPaymentFailed  = 2
	OrderStatusPaymentSuccess = 3
	OrderStatusProcessing     = 4
	OrderStatusShipped        = 5
	OrderStatusCompleted      = 6
	OrderStatusCancelled      = 7
	OrderStatusRefunded       = 8
	OrderStatusFailed         = 9
	OrderStatusDelivered      = 10
)

func (o *Order) GetStatus() string {
	switch o.Status {
	case OrderStatusPending:
		return "PENDING"
	case OrderStatusCancelled:
		return "CANCELLED"
	case OrderStatusCompleted:
		return "COMPLETED"
	case OrderStatusRefunded:
		return "REFUNDED"
	case OrderStatusFailed:
		return "FAILED"
	case OrderStatusDelivered:
		return "DELIVERED"
	default:
		return "UNKNOWN"
	}
}

func Create(tx *gorm.DB, userID uuid.UUID) (*Order, error) {
	if tx == nil {
		tx = db.Instance
	}
	order := &Order{
		UserID: userID,
		Status: OrderStatusPending,
	}
	err := tx.Create(order).Error
	return order, err
}

func CreateBatch(tx *gorm.DB, orderID uint, items []common.CartItem) error {
	if tx == nil {
		tx = db.Instance
	}
	orderItems := make([]OrderItem, len(items))
	for i, item := range items {
		orderItems[i] = OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}
	return tx.Create(&orderItems).Error
}

func Get(tx *gorm.DB, order *Order) error {
	if tx == nil {
		tx = db.Instance
	}

	if err := tx.First(order, order).Error; err != nil {
		return err
	}

	return nil
}

func Update(tx *gorm.DB, order *Order) error {
	if tx == nil {
		tx = db.Instance
	}

	if err := tx.First(order, order.ID).Error; err != nil {
		return err
	}

	return tx.Model(order).Update("status", order.Status).Error
}
