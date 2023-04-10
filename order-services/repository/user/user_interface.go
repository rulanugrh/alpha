package user

import "github.com/rulanugrh/alpha/order/entities/domain"

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
}
