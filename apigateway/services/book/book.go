package book

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rulanugrh/alpha/apigateway/entity/book"
	"github.com/rulanugrh/alpha/apigateway/helper"
)

type bookservices struct{}

func NewBookServices() BookInterface {
	return &bookservices{}
}

func (bk *bookservices) GetBook() (http.Response, error) {
	res, err := http.Get("http://localhost:3000/book")

	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *res, err
}

func (bk *bookservices) GetCategory() (http.Response, error) {
	res, err := http.Get("http://localhost:3000/category")

	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *res, err
}

func (bk *bookservices) PostBook(books book.BookModel) (http.Response, error) {
	json_data, _ := json.Marshal(books)
	resp, err := http.NewRequest("POST", "http://localhost:3000/book", bytes.NewBuffer(json_data))
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil
}

func (bk *bookservices) PostCategory(category book.Category) (http.Response, error) {
	json_data, _ := json.Marshal(category)
	resp, err := http.NewRequest("POST", "http://localhost:3000/category", bytes.NewBuffer(json_data))
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil
}

func (bk *bookservices) GetBookById(id uint) (http.Response, error) {
	books := book.BookModel{}
	books.Id = id

	resp, err := http.NewRequest("GET", "http://localhost:3000/book/:id", nil)
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil

}

func (bk *bookservices) GetCategoryId(id uint) (http.Response, error) {
	category := book.Category{}
	category.Id = id

	resp, err := http.NewRequest("GET", "http://localhost:3000/category/:id", nil)
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil

}

func (bk *bookservices) UpdateBook(id uint, book book.BookModel) (http.Response, error) {
	json_data, _ := json.Marshal(book)
	resp, err := http.NewRequest("PUT", "http://localhost:3000/book/:id", bytes.NewBuffer(json_data))
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil

}
func (bk *bookservices) DeleteBookId(id uint) (http.Response, error) {
	books := book.BookModel{}
	books.Id = id

	resp, err := http.NewRequest("DELETE", "http://localhost:3000/book/:id", nil)
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil

}
