package order

type OrderModel struct {
	Paid  bool `json:"is_paid" form:"is_paid"`
	Total int  `json:"grand_total" form:"grand_total"`
	// OrderAt string `json:"order_at" form:"order_at"`
	UserID uint `json:"user_id" form:"user_id"`
}

type OrderItemModel struct {
	Quantity int  `json:"quantity" form:"quantity"`
	OrderID  uint `json:"order_id" form:"order_id"`
	BookID   uint `json:"book_id" form:"book_id"`
	Price    int  `json:"price" form:"price"`
	Subtotal int  `json:"subtotal" form:"subtotal"`
	UserID   uint `json:"user_id" form:"user_id"`
}
