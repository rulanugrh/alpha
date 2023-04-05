package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" name:"description"`
	Books       []Book `json:"books" gorm:"foreignKey:CategoryID;references:ID"`
}
