package book

import (
	"net/http"

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

func (bk *bookcontroller)