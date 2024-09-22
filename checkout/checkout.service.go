package checkout

import (
	"context"
	"errors"

	"github.com/9thDuck/ecommerce-api.git/cart"
	"github.com/9thDuck/ecommerce-api.git/db"
	"github.com/9thDuck/ecommerce-api.git/orders"
	"github.com/google/uuid"
)

func processCheckout(ctx context.Context, userID uuid.UUID) (uint, error) {
	tx := db.Instance.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cartItems, err := cart.GetCartItemsByUserID(tx, userID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if len(*cartItems) == 0 {
		tx.Rollback()
		return 0, errors.New("cart is empty")
	}

	order, err := orders.Create(tx, userID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = orders.CreateBatch(tx, order.ID, *cartItems)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = cart.DeleteAllFromCart(tx, userID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return order.ID, nil
}
