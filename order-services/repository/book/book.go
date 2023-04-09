package book

import (
	"errors"

	"github.com/rulanugrh/alpha/order/config"
	"github.com/rulanugrh/alpha/order/entities/domain"
)

type bookrepo struct{}

func NewBookRepository() BookRepository {
	return &bookrepo{}
}

func (bk *bookrepo) Create(book domain.Book) (domain.Book, error) {
	err := config.DB.Create(&book).Error
	if err != nil {
		return domain.Book{}, errors.New("cant create book")
	}

	return book, nil
}

func (bk *bookrepo) FindAll() ([]domain.Book, error) {
	var book []domain.Book
	err := config.DB.Find(&book).Error
	if err != nil {
		return []domain.Book{}, errors.New("cant find book")
	}

	return book, nil
}

func (bk *bookrepo) FindID(id uint) (domain.Book, error) {
	var book domain.Book
	err := config.DB.First(&book, id).Error
	if err != nil {
		return domain.Book{}, errors.New("cant find book by this id")
	}

	return book, nil
}
