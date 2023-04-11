package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}
