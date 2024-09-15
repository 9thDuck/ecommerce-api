package cart

import (
	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func addToCartOrIncrement(userID uuid.UUID, productID uint, quantity int) error {
	return upsertCartItem(userID, productID, quantity)
}

func getCartByUserID(userID uuid.UUID) (*[]common.CartItem, error) {
	return GetCartItemsByUserID(nil, userID)
}

func remove(userID uuid.UUID, productID uint) error {
	return decrement(userID, productID)
}

func deleteCartItem(userID uuid.UUID, productID uint) error {
	return delete(userID, productID)
}

func DeleteAllFromCart(tx *gorm.DB, userID uuid.UUID) error {
	return deleteAll(tx, userID)
}
