package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name       string   `json:"name" form:"name" validate:"required"`
	Price      int      `json:"price" form:"price" validate:"required"`
	Stock      int      `json:"stock_quantity" form:"stock_quantity" validate:"requiered"`
	Author     string   `json:"author_name" form:"author_name" validate:"required"`
	Examplar   int      `json:"examplar" form:"examplar" validate:"required"`
	SellerID   Seller   `json:"seller" gorm:"foreignKey:SellerID;references:ID"`
	CategoryID Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}

type Seller struct {
	gorm.Model
	Name  string `json:"name" form:"name" validate:"required"`
	Books []Book `json:"books" gorm:"foreignKey:SellerID;references:ID"`
}

type Category struct {
	gorm.Model
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" name:"description"`
	Books       []Book `json:"books" gorm:"foreignKey:CategoryID;references:ID"`
}
