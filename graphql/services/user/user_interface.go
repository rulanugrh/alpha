package user

import (
	"github.com/rulanugrh/graphql/graph/model"
)

type UserService interface {
	Register(user model.User) (*model.Response, error)
	Login(user model.User) (*model.Response, error)
	Detail(id *string) (*model.Response, error)
}
