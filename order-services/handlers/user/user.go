package user

import (
	"github.com/labstack/echo/v4"
	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
	"github.com/rulanugrh/alpha/order/services/user"
)

type usercontroller struct {
	users user.UserServices
}

func NewUserController(user user.UserServices) UserController {
	return &usercontroller{
		users: user,
	}
}

func (usr *usercontroller) Create(ctx echo.Context) error {
	var user domain.User
	ctx.Bind(&user)

	users, err := usr.users.Create(user)
	if err != nil {
		return ctx.JSON(400, web.UserFailure{
			Code:   400,
			Status: "cant creaet user",
			Error:  err,
		})
	}

	return ctx.JSON(201, web.UserSuccess{
		Code:   201,
		Status: "success create user",
		Data:   users,
	})
}
