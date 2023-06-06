package book

import (
	"net/http"
	"strconv"

	"github.com/rulanugrh/alpha/apigateway/entity/web"

	"github.com/labstack/echo/v4"
	modelBook "github.com/rulanugrh/alpha/apigateway/entity/book"
	"github.com/rulanugrh/alpha/apigateway/services/book"
)

type bookcontroller struct {
	books book.BookInterface
}

func NewBookController(book book.BookInterface) BookController {
	return &bookcontroller{books: book}
}

func (bk *bookcontroller) GetBook(ctx echo.Context) error {
	resp, err := bk.books.GetBook()
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Cannot Get Book",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Success Get Data",
		Data:   resp.Body,
	}

	return ctx.JSON(200, response)
}

func (bk *bookcontroller) GetCategory(ctx echo.Context) error {
	resp, err := bk.books.GetCategory()
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Cannot Get Book",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Success Get Data",
		Data:   resp.Body,
	}

	return ctx.JSON(200, response)
}

func (bk *bookcontroller) GetBookId(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)
	resp, err := bk.books.GetBookById(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Cannot Get Book",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Success Get Data",
		Data:   resp.Body,
	}

	return ctx.JSON(200, response)
}

func (bk *bookcontroller) GetCategoryId(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)
	resp, err := bk.books.GetCategoryId(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Cannot Get Category",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Success Get Data",
		Data:   resp.Body,
	}

	return ctx.JSON(200, response)
}

func (bk *bookcontroller) PostBook(ctx echo.Context) error {
	var book modelBook.BookModel
	ctx.Bind(&book)
	resp, err := bk.books.PostBook(book)
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Cannot Post Book",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Success Get Data",
		Data:   resp.Body,
	}

	return ctx.JSON(200, response)
}

func (bk *bookcontroller) PostCategory(ctx echo.Context) error {
	var book modelBook.Category
	ctx.Bind(&book)
	resp, err := bk.books.PostCategory(book)
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Cannot Post Category",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Success Get Data",
		Data:   resp.Body,
	}

	return ctx.JSON(200, response)
}

func (bk *bookcontroller) DeleteBookById(ctx echo.Context) error {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)
	resp, err := bk.books.DeleteBookId(uint(id))
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Cannot Delete Book",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Success Get Data",
		Data:   resp.Body,
	}

	return ctx.JSON(200, response)
}

func (bk *bookcontroller) UpdateBook(ctx echo.Context) error {
	var book modelBook.BookModel
	ctx.Bind(&book)

	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	resp, err := bk.books.UpdateBook(uint(id), book)
	if err != nil {
		response := web.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Cannot Update Book",
			Error:  err,
		}
		return ctx.JSON(http.StatusBadRequest, response)
	}

	response := web.SuccessResponse{
		Code:   200,
		Status: "Success Get Data",
		Data:   resp.Body,
	}

	return ctx.JSON(200, response)
}
