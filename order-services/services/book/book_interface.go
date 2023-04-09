package book

import (
	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
)

type BookServices interface {
	UploadBook(book domain.Book) (web.BookResponse, error)
	FindById(id uint) (web.BookResponse, error)
	FindAll() ([]domain.Book, error)
}
