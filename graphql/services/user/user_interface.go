package user

import (
	"github.com/rulanugrh/graphql/graph/model"
)

type UserService interface {
	Register(user model.User) (*model.User, error)
	Login(user model.User) (*model.User, error)
	Detail(id *string) (*model.User, error)
}
