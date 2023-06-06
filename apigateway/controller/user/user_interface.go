package user

import "github.com/labstack/echo/v4"

type UserController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	Detail(ctx echo.Context) error
}
