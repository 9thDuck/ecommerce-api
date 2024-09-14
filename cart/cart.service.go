package cart

import (
	"github.com/google/uuid"
)

type cartItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func addToCartOrIncrement(userID uuid.UUID, productID uint, quantity int) error {
	return upsertCartItem(userID, productID, quantity)
}

func getCartByUserID(userID uuid.UUID) (*[]cartItem, error) {
	return getCartItemsByUserID(userID)
}

func remove(userID uuid.UUID, productID uint) error {
	return decrement(userID, productID)
}

func deleteCartItem(userID uuid.UUID, productID uint) error {
	return delete(userID, productID)
}

func deleteAllFromCart(userID uuid.UUID) error {
	return deleteAll(userID)
}
