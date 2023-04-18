package books

import (
	"github.com/rulanugrh/alpha/book/entities/domain"
	"github.com/rulanugrh/alpha/book/entities/web"
)

type BookService interface {
	UploadBook(book domain.Book) (web.BookResponse, error)
	UpdateBook(id uint, book domain.Book) (web.BookResponse, error)
	FindId(id uint) (web.BookResponse, error)
	DeleteBook(id uint) (web.BookDeleteRespons, error)
	FindAll() ([]domain.Book, error)
	Seller(seller domain.Seller) (web.SellerResponse, error)
}
