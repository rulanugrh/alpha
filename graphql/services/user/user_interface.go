package user

import (
	"net/http"

	"github.com/rulanugrh/graphql/graph/model"
)

type UserService interface {
	Register(user model.User) (http.Response, error)
	Login(user model.User) (http.Response, error)
	Detail(id *int) (http.Response, error)
}
