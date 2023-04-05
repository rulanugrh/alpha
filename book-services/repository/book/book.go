package books

import (
	"errors"

	"github.com/rulanugrh/alpha/book/config"
	"github.com/rulanugrh/alpha/book/entities/domain"
)

type bookrepository struct{}

func NewBookRepository() BookRepository {
	return &bookrepository{}
}

func (bk *bookrepository) Create(book domain.Book) (domain.Book, error) {
	err := config.DB.Create(&book).Error
	if err != nil {
		return book, errors.New("cant create book")
	}

	return book, nil
}

func (bk *bookrepository) Update(id uint, book domain.Book) (domain.Book, error) {
	var bookupdate domain.Book
	err := config.DB.Model(&book).Where("id = ?", id).Updates(&bookupdate).Error
	if err != nil {
		return bookupdate, errors.New("cant update book")
	}

	return bookupdate, nil
}

func (bk *bookrepository) FindAll() ([]domain.Book, error) {
	var book []domain.Book
	err := config.DB.Find(&book)
	if err != nil {
		return book, errors.New("cant find all book")
	}

	return book, nil
}

func (bk *bookrepository) FindId(id uint) (domain.Book, error) {
	var book domain.Book
	err := config.DB.First(&book, id)
	if err != nil {
		return book, errors.New("cant find book")
	}

	return book, nil
}

func (bk *bookrepository) Delete(id uint) error {
	var book domain.Book
	err := config.DB.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		panic(err)
	}
	return nil
}

func (bk *bookrepository) Seller(seller domain.Seller) (domain.Seller, error) {
	err := config.DB.Create(&seller).Error
	if err != nil {
		return seller, errors.New("cant create seller")
	}

	return seller, nil
}
