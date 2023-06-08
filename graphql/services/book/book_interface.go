package book

import (
	"net/http"

	"github.com/rulanugrh/graphql/graph/model"
)

type BookInterface interface {
	GetCategory() (http.Response, error)
	PostCategory(category model.Category) (http.Response, error)
	GetCategoryId(id *int) (http.Response, error)
	GetBook() (http.Response, error)
	PostBook(books model.Book) (http.Response, error)
	DeleteBookId(id *int) (http.Response, error)
	UpdateBook(id *int, books model.Book) (http.Response, error)
	GetBookById(id *int) (http.Response, error)
}
