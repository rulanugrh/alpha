package user

import (
	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
)

type UserServices interface {
	Create(user domain.User) (web.UserResponse, error)
}
