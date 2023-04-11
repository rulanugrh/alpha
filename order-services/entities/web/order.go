package web

type OrderSuccess struct {
	Code   int         `json:"code" form:"code"`
	Status string      `json:"status" form:"status"`
	Data   interface{} `json:"data" form:"data"`
}

type OrderFailure struct {
	Code   int    `json:"code" form:"code"`
	Status string `json:"status" form:"status"`
}

type OrderResponse struct {
	OrderID  uint `json:"order_id" form:"order_id"`
	Subtotal int  `json:"subtotal" form:"subtotal"`
	Paid     bool `json:"paid" form:"paid"`
}

type OrderItemResponse struct {
	Quantity int  `json:"quantity" form:"quantity"`
	BookID   uint `json:"book_id" form:"book_id"`
	Subtotal int  `json:"subtotal" form:"subtotal"`
}
