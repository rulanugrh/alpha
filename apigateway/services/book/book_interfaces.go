package book

import (
	"net/http"

	"github.com/rulanugrh/alpha/apigateway/entity/book"
)

type BookInterface interface {
	GetCategory() (http.Response, error)
	PostCategory(category book.Category) (http.Response, error)
	GetCategoryId(id uint) (http.Response, error)
	GetBook() (http.Response, error)
	PostBook(books book.BookModel) (http.Response, error)
	DeleteBookId(id uint) (http.Response, error)
	UpdateBook(id uint, books book.BookModel) (http.Response, error)
	GetBookById(id uint) (http.Response, error)
}
