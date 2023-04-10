package book

import "github.com/labstack/echo/v4"

type BookController interface {
	Create(ctx echo.Context) error
	FindId(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}
