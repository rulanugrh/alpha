package routers

import (
	"github.com/labstack/echo/v4"
	bookControll "github.com/rulanugrh/alpha/apigateway/controller/book"
	"github.com/rulanugrh/alpha/apigateway/helper"
)

func App(books bookControll.BookController) {

	serv := echo.New()
	apiBook := serv.Group("/book")
	{
		apiBook.GET("/", books.GetBook)
		apiBook.GET("/:id", books.GetBookId)
		apiBook.POST("/", books.PostBook)
		apiBook.PUT("/:id", books.UpdateBook)
		apiBook.DELETE("/:id", books.DeleteBookById)
	}

	apiCate := serv.Group("/category")
	{
		apiCate.GET("/", books.GetCategory)
		apiCate.POST("/", books.GetCategory)
		apiCate.GET("/:id", books.GetCategoryId)
	}

	err := serv.Start(":8080")
	if err != nil {
		helper.PanicIfError(err)
	}
}
