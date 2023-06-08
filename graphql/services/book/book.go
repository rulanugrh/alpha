package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rulanugrh/graphql/graph/model"
	"github.com/rulanugrh/graphql/helper"
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

func (bk *bookservices) PostBook(books model.Book) (http.Response, error) {
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

func (bk *bookservices) PostCategory(category model.Category) (http.Response, error) {
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

func (bk *bookservices) GetBookById(id *int) (http.Response, error) {
	books := model.Book{}
	books.ID = id

	requesUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)
	resp, err := http.NewRequest("GET", requesUrl, nil)
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil

}

func (bk *bookservices) GetCategoryId(id *int) (http.Response, error) {
	category := model.Category{}
	category.ID = id

	requestUrl := fmt.Sprintf("http://localhost:3000/category/%d", id)
	resp, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil

}

func (bk *bookservices) UpdateBook(id *int, book model.Book) (http.Response, error) {
	json_data, _ := json.Marshal(book)
	requestUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)
	resp, err := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(json_data))
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil

}
func (bk *bookservices) DeleteBookId(id *int) (http.Response, error) {
	books := model.Book{}
	books.ID = id

	requestUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)

	resp, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		helper.PanicIfError(err)
		return http.Response{
			Status:     "cant request to server api",
			StatusCode: 500,
		}, err
	}

	return *resp.Response, nil

}
