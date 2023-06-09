package book

import (
	"github.com/rulanugrh/graphql/graph/model"
)

type BookInterface interface {
	GetCategory() (*model.Response, error)
	PostCategory(category model.Category) (*model.Response, error)
	GetCategoryId(id *string) (*model.Response, error)
	GetBook() (*model.Response, error)
	PostBook(books model.Book) (*model.Response, error)
	DeleteBookId(id *string) (*model.Response, error)
	UpdateBook(id *string, books model.Book) (*model.Response, error)
	GetBookById(id *string) (*model.Response, error)
}
