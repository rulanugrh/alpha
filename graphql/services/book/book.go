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

func (bk *bookservices) GetBook() (*model.Response, error) {

	res, err := http.Get("http://localhost:3000/book")

	var sb strings.Builder
	io.Copy(&sb, res.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &data, err
}

func (bk *bookservices) GetCategory() (*model.Response, error) {
	res, err := http.Get("http://localhost:3000/category")

	var sb strings.Builder
	io.Copy(&sb, res.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &data, err

}

func (bk *bookservices) PostBook(books model.Book) (*model.Response, error) {
	json_data, _ := json.Marshal(books)
	resp, err := http.NewRequest("POST", "http://localhost:3000/book", bytes.NewBuffer(json_data))

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &data, nil
}

func (bk *bookservices) PostCategory(category model.Category) (*model.Response, error) {
	json_data, _ := json.Marshal(category)
	resp, err := http.NewRequest("POST", "http://localhost:3000/category", bytes.NewBuffer(json_data))

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &data, nil
}

func (bk *bookservices) GetBookById(id *string) (*model.Response, error) {

	requesUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)
	resp, err := http.NewRequest("GET", requesUrl, nil)
	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &data, nil

}

func (bk *bookservices) GetCategoryId(id *string) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:3000/category/%d", id)
	resp, err := http.NewRequest("GET", requestUrl, nil)

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}
	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &data, nil

}

func (bk *bookservices) UpdateBook(id *string, book model.Book) (*model.Response, error) {
	json_data, _ := json.Marshal(book)
	requestUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)
	resp, err := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(json_data))
	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &data, nil

}
func (bk *bookservices) DeleteBookId(id *string) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:3000/book/%d", id)

	resp, err := http.NewRequest("DELETE", requestUrl, nil)
	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		helper.PanicIfError(err)
		return nil, err
	}

	return &data, nil

}
