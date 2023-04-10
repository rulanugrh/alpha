package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
}
