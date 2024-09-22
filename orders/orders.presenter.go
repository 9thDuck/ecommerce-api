package orders

import "github.com/9thDuck/ecommerce-api.git/common"

func errorCancelOrderResponse() *common.Response {
	return &common.Response{
		Message: "Failed to cancel order",
	}
}

func successCancelOrderResponse() *common.Response {
	return &common.Response{
		Message: "Order cancelled successfully",
	}
}

func successGetOrderResponse(order *Order) *common.Response {
	return &common.Response{
		Message: "Order retrieved successfully",
		Data:    *order,
	}
}
