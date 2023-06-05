package book

import "github.com/labstack/echo/v4"

type BookController interface {
	GetBook(ctx echo.Context) error
	GetBookId(ctx echo.Context) error
	GetCategory(ctx echo.Context) error
	GetCategoryId(ctx echo.Context) error
	PostCategory(ctx echo.Context) error
	PostBook(ctx echo.Context) error
	DeleteBookById(ctx echo.Context) error
	UpdateBook(ctx echo.Context) error
}
