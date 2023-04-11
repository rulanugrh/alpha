package domain

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Paid  bool `json:"is_paid" form:"is_paid"`
	Total int  `json:"grand_total" form:"grand_total"`
	// OrderAt string `json:"order_at" form:"order_at"`
	UserID uint `json:"user_id" form:"user_id"`
	User   User `json:"user" form:"user"`
}

type OrderItem struct {
	gorm.Model
	Quantity int   `json:"quantity" form:"quantity"`
	OrderID  uint  `json:"order_id" form:"order_id"`
	BookID   uint  `json:"book_id" form:"book_id"`
	Price    int   `json:"price" form:"price"`
	Subtotal int   `json:"subtotal" form:"subtotal"`
	UserID   uint  `json:"user_id" form:"user_id"`
	User     User  `json:"user" gorm:"foreignKey:UserID"`
	Books    Book  `json:"books" gorm:"foreignKey:BookID"`
	Orders   Order `json:"orders" gorm:"foreignKey:OrderID"`
}

type Book struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}
