package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rulanugrh/alpha/book/entities/domain"
	"github.com/rulanugrh/alpha/book/entities/web"
	books "github.com/rulanugrh/alpha/book/services/book"
)

type bookcontroller struct {
	bookcontroll books.BookService
}

func NewBookController(book books.BookService) BookController {
	return &bookcontroller{
		bookcontroll: book,
	}
}

func (bk *bookcontroller) UploadBook(ctx *gin.Context) {
	book := domain.Book{}
	ctx.Bind(&book)

	response, err := bk.bookcontroll.UploadBook(book)
	if err != nil {
		failure := web.BookFailure{
			Code:   500,
			Status: "Cant Upload Book",
		}

		ctx.JSON(http.StatusBadRequest, failure)
	}

	success := web.BookSuccess{
		Code:   201,
		Status: "success create book",
		Data:   response,
	}

	ctx.JSON(http.StatusOK, success)
}

func (bk *bookcontroller) UpdateBook(ctx *gin.Context) {
	book := domain.Book{}
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	ctx.Bind(&book)
	response, err := bk.bookcontroll.UpdateBook(uint(id), book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.BookFailure{
			Code:   500,
			Status: "cant update book",
		})
	}

	ctx.JSON(http.StatusOK, web.BookSuccess{
		Code:   200,
		Status: "success update book",
		Data:   response,
	})
}

func (bk *bookcontroller) DeleteBook(ctx *gin.Context) {
	getId := ctx.Param("id")
	ID, _ := strconv.Atoi(getId)

	err := bk.bookcontroll.DeleteBook(uint(ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.BookFailure{
			Code:   500,
			Status: "cant delete book",
		})
	}

	ctx.JSON(http.StatusOK, web.BookSuccess{
		Code:   200,
		Status: "success delete book",
		Data:   nil,
	})
}

func (bk *bookcontroller) FindID(ctx *gin.Context) {
	getID := ctx.Param("id")
	id, _ := strconv.Atoi(getID)

	response, err := bk.bookcontroll.FindId(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.BookFailure{
			Code:   500,
			Status: "cant find book",
		})
	}

	ctx.JSON(http.StatusOK, web.BookSuccess{
		Code:   200,
		Status: "success find book",
		Data:   response,
	})
}

func (bk *bookcontroller) FindAll(ctx *gin.Context) {
	response, err := bk.bookcontroll.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.BookFailure{
			Code:   500,
			Status: "cant find all book",
		})
	}

	ctx.JSON(http.StatusOK, web.BookSuccess{
		Code:   200,
		Status: "sucess find all",
		Data:   response,
	})
}

func (bk *bookcontroller) Seller(ctx *gin.Context) {
	seller := domain.Seller{}
	ctx.Bind(&seller)

	response, err := bk.bookcontroll.Seller(seller)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.BookFailure{
			Code:   500,
			Status: "cant create seller",
		})
	}

	ctx.JSON(http.StatusOK, web.BookSuccess{
		Code:   200,
		Status: "success create user",
		Data:   response,
	})
}
