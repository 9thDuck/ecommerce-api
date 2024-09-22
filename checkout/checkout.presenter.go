package checkout

import (
	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/9thDuck/ecommerce-api.git/orders"
)

func successCheckoutResponse(orderID uint) common.Response {
	return common.Response{
		Message: "Checkout successful",
		Data:    orderID,
	}
}

func errorCheckoutResponse() common.Response {
	return common.Response{
		Message: "Checkout failed",
	}
}

func errorCheckoutForbiddenResponse() common.Response {
	return common.Response{
		Message: "Forbidden",
	}
}

func errorEmptyCartResponse() common.Response {
	return common.Response{
		Message: "Cart is empty",
	}
}

func errorCheckoutInternalServerErrorResponse() common.Response {
	return common.Response{
		Message: "Something went wrong",
	}
}

func errorCheckoutBadRequestResponse() common.Response {
	return common.Response{
		Message: "Bad request",
	}
}

func successCancelCheckoutResponse(order *orders.Order) common.Response {
	return common.Response{
		Message: "Cancel checkout successful",
		Data:    order,
	}
}

func errorCancelCheckoutResponse() common.Response {
	return common.Response{
		Message: "Cancel checkout failed",
	}
}
