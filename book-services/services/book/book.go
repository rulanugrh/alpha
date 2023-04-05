package books

import (
	"errors"

	"github.com/rulanugrh/alpha/book/entities/domain"
	"github.com/rulanugrh/alpha/book/entities/web"
	books "github.com/rulanugrh/alpha/book/repository/book"
)

type bookservice struct {
	book books.BookRepository
}

func NewBookRepository(bookrepo books.BookRepository) BookService {
	return &bookservice{
		book: bookrepo,
	}
}

func (bk *bookservice) UploadBook(book domain.Book) (web.BookResponse, error) {
	response, err := bk.book.Create(book)
	if err != nil {
		return web.BookResponse{}, errors.New("cant upload book")
	}

	resp := web.BookResponse{
		Name:   response.Name,
		Author: response.Author,
		Price:  response.Price,
		Stock:  response.Stock,
	}

	return resp, nil
}

func (bk *bookservice) UpdateBook(id uint, book domain.Book) (web.BookResponse, error) {
	response, err := bk.book.Update(id, book)
	if err != nil {
		return web.BookResponse{}, errors.New("cant update book")
	}

	return web.BookResponse{
		Name:   response.Name,
		Author: response.Author,
		Price:  response.Price,
		Stock:  response.Stock,
	}, nil
}

func (bk *bookservice) FindId(id uint) (web.BookResponse, error) {
	response, err := bk.book.FindId(id)
	if err != nil {
		return web.BookResponse{}, errors.New("cant find book")
	}

	return web.BookResponse{
		Name:   response.Name,
		Author: response.Author,
		Price:  response.Price,
		Stock:  response.Stock,
	}, nil
}

func (bk *bookservice) DeleteBook(id uint) error {
	err := bk.book.Delete(id)
	if err != nil {
		return errors.New("cant delete book")
	}

	return nil
}

func (bk *bookservice) FindAll() ([]domain.Book, error) {
	book, err := bk.book.FindAll()
	if err != nil {
		return book, errors.New("cant find all book")
	}
	return book, nil
}

func (bk *bookservice) Seller(seller domain.Seller) (web.SellerResponse, error) {
	reponse, err := bk.book.Seller(seller)
	if err != nil {
		return web.SellerResponse{}, errors.New("cant create seller")
	}

	return web.SellerResponse{
		Name: reponse.Name,
	}, nil
}
