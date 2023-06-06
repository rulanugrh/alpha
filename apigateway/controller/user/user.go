package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	userMod "github.com/rulanugrh/alpha/apigateway/entity/user"
	"github.com/rulanugrh/alpha/apigateway/entity/web"
	"github.com/rulanugrh/alpha/apigateway/services/user"
)

type usercontoller struct {
	users user.UserService
}

func NewUserController(user user.UserService) UserController {
	return &usercontoller{users: user}
}

func (usr *usercontoller) Register(ctx echo.Context) error {
	var user userMod.UserModel
	ctx.Bind(&user)

	resp, err := usr.users.Register(user)
	if err != nil {
		response := web.ErrorResponse{
			Status: "cannot register",
			Code:   http.StatusBadRequest,
			Error:  err,
		}

		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Status: "success register",
		Code:   200,
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (usr *usercontoller) Login(ctx echo.Context) error {
	var user userMod.UserModel
	ctx.Bind(&user)

	resp, err := usr.users.Login(user)
	if err != nil {
		response := web.ErrorResponse{
			Status: "cannot login",
			Code:   http.StatusBadRequest,
			Error:  err,
		}

		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Status: "success login",
		Code:   200,
		Data:   resp,
	}

	return ctx.JSON(200, response)
}

func (usr *usercontoller) Detail(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	resp, err := usr.users.Detail(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Status: "cannot seee",
			Code:   http.StatusBadRequest,
			Error:  err,
		}

		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Status: "success",
		Code:   200,
		Data:   resp,
	}

	return ctx.JSON(200, response)
}
