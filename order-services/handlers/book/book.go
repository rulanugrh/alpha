package book

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
	"github.com/rulanugrh/alpha/order/services/book"
)

type bookcontroller struct {
	books book.BookServices
}

func NewBookController(book book.BookServices) BookController {
	return &bookcontroller{
		books: book,
	}
}

func (bk *bookcontroller) Create(ctx echo.Context) error {
	var book domain.Book
	ctx.Bind(&book)

	books, err := bk.books.UploadBook(book)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.BookFailuer{
			Code:   400,
			Status: "cant create book",
		})
	}

	return ctx.JSON(201, web.BookSuccess{
		Code:   201,
		Status: "success upload book",
		Data:   books,
	})
}

func (bk *bookcontroller) FindId(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	books, err := bk.books.FindById(uint(id))
	if err != nil {
		return ctx.JSON(400, web.BookFailuer{
			Code:   400,
			Status: "cant find book by this id",
		})
	}

	return ctx.JSON(200, web.BookSuccess{
		Code:   200,
		Status: "book find",
		Data:   books,
	})
}

func (bk *bookcontroller) FindAll(ctx echo.Context) error {
	books, err := bk.books.FindAll()
	if err != nil {
		return ctx.JSON(400, web.BookFailuer{
			Code:   400,
			Status: "data null",
		})
	}

	return ctx.JSON(200, web.BookSuccess{
		Code:   200,
		Status: "data found",
		Data:   books,
	})
}
