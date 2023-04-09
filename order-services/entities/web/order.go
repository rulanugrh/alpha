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
	OrderID  int `json:"order_id" form:"order_id"`
	Price    int `json:"price" form:"price"`
	Subtotal int `json:"subtotal" form:"subtotal"`
}
