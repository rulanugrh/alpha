package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rulanugrh/graphql/graph/model"
)

type orderservice struct{}

func NewOrderSercives() OrderInterfaces {
	return &orderservice{}
}
func (or *orderservice) Create(order model.Order) (http.Response, error) {
	json_data, _ := json.Marshal(&order)
	resp, err := http.NewRequest("POST", "http://localhost:5000/", bytes.NewBuffer(json_data))
	if err != nil {
		return http.Response{
			Status:     "Cannot create order",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp.Response, nil
}

func (or *orderservice) FindId(id *int) (http.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/%d", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return http.Response{
			Status:     "Cannot get data",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp, nil
}

func (or *orderservice) FindAll(id *int) (http.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/getAll/%d", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return http.Response{
			Status:     "cannot get data",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp, nil
}

func (or *orderservice) AddCart(order model.OrderItem) (http.Response, error) {
	json_data, _ := json.Marshal(&order)

	resp, err := http.NewRequest("POST", "http://localhost:5000/cart", bytes.NewBuffer(json_data))
	if err != nil {
		return http.Response{
			Status:     "cannot add cart",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp.Response, nil
}

func (or *orderservice) Checkout(id *int, order model.Order) (http.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/checkout/%d", id)
	json_data, _ := json.Marshal(&order)

	resp, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(json_data))

	if err != nil {
		return http.Response{
			Status:     "cannot checkout",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp.Response, nil
}

func (or *orderservice) ListCart(userId *int) (http.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d", userId)
	resp, err := http.Get(requestUrl)

	if err != nil {
		return http.Response{
			Status:     "cannot get list cart",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp, nil
}

func (or *orderservice) DeleteCart(id *int) (http.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d", id)
	resp, err := http.NewRequest("DELETE", requestUrl, nil)

	if err != nil {
		return http.Response{
			Status:     "cannot delete cart",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp.Response, nil
}

func (or *orderservice) ListNotPaid(id *int) (http.Response, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d/notpaid", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return http.Response{
			Status:     "cannot get data",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	return *resp, nil
}
