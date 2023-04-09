package book

import (
	"errors"

	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
	"github.com/rulanugrh/alpha/order/repository/book"
)

type bookservices struct {
	books book.BookRepository
}

func NewBookServices(books book.BookRepository) BookServices {
	return &bookservices{
		books: books,
	}
}

func (bk *bookservices) UploadBook(book domain.Book) (web.BookResponse, error) {
	books, err := bk.books.Create(book)
	if err != nil {
		return web.BookResponse{}, errors.New("cant craete book")
	}

	return web.BookResponse{
		Name:  books.Name,
		Price: books.Price,
	}, nil
}

func (bk *bookservices) FindById(id uint) (web.BookResponse, error) {
	books, err := bk.books.FindID(id)
	if err != nil {
		return web.BookResponse{}, errors.New("cant find book by this id")
	}

	return web.BookResponse{
		Name:  books.Name,
		Price: books.Price,
	}, nil
}

func (bk *bookservices) FindAll() ([]domain.Book, error) {
	books, err := bk.books.FindAll()
	if err != nil {
		return []domain.Book{}, errors.New("cant find all book")
	}

	return books, nil
}
