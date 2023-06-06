package user

import (
	"net/http"

	"github.com/rulanugrh/alpha/apigateway/entity/user"
)

type UserService interface {
	Register(user user.UserModel) (http.Response, error)
	Login(user user.UserModel) (http.Response, error)
	Detail(id uint) (http.Response, error)
}
