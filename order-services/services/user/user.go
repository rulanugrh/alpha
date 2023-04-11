package user

import (
	"errors"

	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
	"github.com/rulanugrh/alpha/order/repository/user"
)

type userservice struct {
	users user.UserRepository
}

func NewUserServices(user user.UserRepository) UserServices {
	return &userservice{
		users: user,
	}
}

func (usr userservice) Create(user domain.User) (web.UserResponse, error) {
	users, err := usr.users.Create(user)
	if err != nil {
		return web.UserResponse{}, errors.New("cant create user")
	}

	return web.UserResponse{
		Name: users.Name,
	}, nil
}
