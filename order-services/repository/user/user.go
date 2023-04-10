package user

import (
	"errors"

	"github.com/rulanugrh/alpha/order/config"
	"github.com/rulanugrh/alpha/order/entities/domain"
)

type userrepo struct{}

func NewUserRepository() UserRepository {
	return &userrepo{}
}

func (usr *userrepo) Create(user domain.User) (domain.User, error) {
	err := config.DB.Create(&user).Error
	if err != nil {
		return domain.User{}, errors.New("cant create user")
	}

	return user, nil
}
