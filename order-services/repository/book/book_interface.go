package book

import "github.com/rulanugrh/alpha/order/entities/domain"

type BookRepository interface {
	Create(book domain.Book) (domain.Book, error)
	FindAll() ([]domain.Book, error)
	FindID(id uint) (domain.Book, error)
}
