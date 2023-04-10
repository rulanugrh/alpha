package user

import "github.com/labstack/echo/v4"

type UserController interface {
	Create(ctx echo.Context) error
}
