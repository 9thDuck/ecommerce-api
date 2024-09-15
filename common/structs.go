package common

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CartItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
