package book

import (
	"github.com/rulanugrh/graphql/graph/model"
)

type BookInterface interface {
	GetCategory() ([]*model.Category, error)
	PostCategory(category model.Category) (*model.Category, error)
	GetCategoryId(id *string) (*model.Category, error)
	GetBook() ([]*model.Book, error)
	PostBook(books model.Book) (*model.Book, error)
	DeleteBookId(id *string) (*model.Book, error)
	UpdateBook(id *string, books model.Book) (*model.Book, error)
	GetBookById(id *string) (*model.Book, error)
}
