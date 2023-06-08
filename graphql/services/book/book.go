package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rulanugrh/graphql/graph/model"
	"github.com/rulanugrh/graphql/helper"
)

type bookservices struct{}

func NewBookServices() BookInterface {
	return &bookservices{}
}

func (bk *bookservices) GetBook() ([]*model.Book, error) {

	res, err := http.Get("http://localhost:3000/book")
	var book model.Book

	var sb strings.Builder
	io.Copy(&sb, res.Body)
	resData := sb.String()

	book.Name = &resData

	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	var books []*model.Book
	books = append(books, &book)

	return books, err
}

func (bk *bookservices) GetCategory() ([]*model.Category, error) {
	res, err := http.Get("http://localhost:3000/category")

	var category model.Category

	var sb strings.Builder
	io.Copy(&sb, res.Body)
	resData := sb.String()

	category.Name = &resData

	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	var categories []*model.Category
	categories = append(categories, &category)

	return categories, err

}

func (bk *bookservices) PostBook(books model.Book) (*model.Book, error) {
	json_data, _ := json.Marshal(books)
	_, err := http.NewRequest("POST", "http://localhost:3000/book", bytes.NewBuffer(json_data))
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &books, nil
}

func (bk *bookservices) PostCategory(category model.Category) (*model.Category, error) {
	json_data, _ := json.Marshal(category)
	_, err := http.NewRequest("POST", "http://localhost:3000/category", bytes.NewBuffer(json_data))
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &category, nil
}

func (bk *bookservices) GetBookById(id *string) (*model.Book, error) {
	var books model.Book

	requesUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)
	resp, err := http.NewRequest("GET", requesUrl, nil)
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	books.Name = &resData

	return &books, nil

}

func (bk *bookservices) GetCategoryId(id *string) (*model.Category, error) {
	var category model.Category

	requestUrl := fmt.Sprintf("http://localhost:3000/category/%d", id)
	resp, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	category.Name = &resData
	return &category, nil

}

func (bk *bookservices) UpdateBook(id *string, book model.Book) (*model.Book, error) {
	json_data, _ := json.Marshal(book)
	requestUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)
	_, err := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(json_data))
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &book, nil

}
func (bk *bookservices) DeleteBookId(id *string) (*model.Book, error) {
	var books model.Book

	requestUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)

	resp, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	books.Name = &resData
	return &books, nil

}
