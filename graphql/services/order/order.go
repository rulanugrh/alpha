package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rulanugrh/graphql/graph/model"
)

type orderservice struct{}

func NewOrderSercives() OrderInterfaces {
	return &orderservice{}
}
func (or *orderservice) Create(order model.Order) (*model.Response, error) {
	json_data, _ := json.Marshal(&order)
	resp, err := http.NewRequest("POST", "http://localhost:5000/", bytes.NewBuffer(json_data))
	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (or *orderservice) FindId(id *string) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/%d", id)
	resp, err := http.Get(requestUrl)
	var sb strings.Builder
	io.Copy(&sb, resp.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (or *orderservice) FindAll(id *string) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/getAll/%d", id)
	resp, err := http.Get(requestUrl)

	var sb strings.Builder
	io.Copy(&sb, resp.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (or *orderservice) AddCart(order model.OrderItem) (*model.Response, error) {
	json_data, _ := json.Marshal(&order)

	resp, err := http.NewRequest("POST", "http://localhost:5000/cart", bytes.NewBuffer(json_data))

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (or *orderservice) Checkout(id *string, order model.Order) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/checkout/%d", id)
	json_data, _ := json.Marshal(&order)

	resp, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(json_data))

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (or *orderservice) ListCart(userId *string) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d", userId)
	resp, err := http.Get(requestUrl)

	var sb strings.Builder
	io.Copy(&sb, resp.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (or *orderservice) DeleteCart(id *string) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d", id)
	resp, err := http.NewRequest("DELETE", requestUrl, nil)

	var sb strings.Builder
	io.Copy(&sb, resp.Response.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (or *orderservice) ListNotPaid(id *string) (*model.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d/notpaid", id)
	resp, err := http.Get(requestUrl)

	var sb strings.Builder
	io.Copy(&sb, resp.Body)
	resData := sb.String()

	data := model.Response{
		Data: &resData,
	}

	if err != nil {
		return nil, err
	}

	return &data, nil

}
