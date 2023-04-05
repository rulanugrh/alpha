package books

import "github.com/rulanugrh/alpha/book/entities/domain"

type BookRepository interface {
	Create(book domain.Book) (domain.Book, error)
	Update(id uint, book domain.Book) (domain.Book, error)
	FindAll() ([]domain.Book, error)
	FindId(id uint) (domain.Book, error)
	Delete(id uint) error
	Seller(seller domain.Seller) (domain.Seller, error)
}
