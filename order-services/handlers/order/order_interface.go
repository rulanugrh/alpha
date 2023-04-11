package order

import "github.com/labstack/echo/v4"

type OrderController interface {
	CreateOrder(ctx echo.Context) error
	FindId(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	AddCart(ctx echo.Context) error
	Checkout(ctx echo.Context) error
	ListCart(ctx echo.Context) error
	DeleteCart(ctx echo.Context) error
	ListNotPaid(ctx echo.Context) error
}
