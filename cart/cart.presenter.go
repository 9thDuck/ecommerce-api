package cart

import (
	"github.com/9thDuck/ecommerce-api.git/common"
)

type cartResponse struct {
	Items []common.CartItem `json:"items"`
}

func successGetCartResponse(cartItems *[]common.CartItem) *common.Response {
	return &common.Response{
		Message: "Cart retrieved successfully",
		Data: cartResponse{
			Items: *cartItems,
		},
	}
}

func successAddToCartResponse() *common.Response {
	return &common.Response{
		Message: "Item added to cart successfully",
	}
}

func successRemoveFromCartResponse() *common.Response {
	return &common.Response{
		Message: "Item removed from cart successfully",
	}
}

func errorAddToCartResponse() *common.Response {
	return &common.Response{
		Message: "Failed to add item to cart",
	}
}

func errorRemoveFromCartResponse() *common.Response {
	return &common.Response{
		Message: "Failed to remove item from cart",
	}
}

func errorDeleteProductFromCartResponse() *common.Response {
	return &common.Response{
		Message: "Failed to delete product from cart",
	}
}

func errorClearCartResponse() *common.Response {
	return &common.Response{
		Message: "Failed to clear cart",
	}
}

func successClearCartResponse() *common.Response {
	return &common.Response{
		Message: "Cart cleared successfully",
	}
}

func errorGetCartResponse() *common.Response {
	return &common.Response{
		Message: "Failed to get cart",
	}
}
