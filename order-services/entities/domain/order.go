package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Paid    bool      `json:"is_paid" form:"is_paid"`
	Total   int       `json:"grand_total" form:"grand_total"`
	OrderAt time.Time `json:"order_at" form:"order_at"`
}

type OrderTime struct {
	gorm.Model
	OrderID  uint  `json:"order_id" form:"order_id"`
	BookID   uint  `json:"book_id" form:"book_id"`
	Quantity int   `json:"quantity" form:"quantity"`
	Price    int   `json:"price" form:"price"`
	Subtotal int   `json:"subtotal" form:"subtotal"`
	Orders   Order `json:"order" gorm:"foreignKey:OrderID;reference:ID"`
	Books    Book  `json:"book" gorm:"foreignKey:BookID;reference:ID"`
}

type Book struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}
