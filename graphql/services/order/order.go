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
func (or *orderservice) Create(order model.Order) (*model.Order, error) {
	json_data, _ := json.Marshal(&order)
	resp, err := http.NewRequest("POST", "http://localhost:5000/", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return &order, nil
}

func (or *orderservice) FindId(id *string) (*model.Order, error) {
	var order model.Order
	requestUrl := fmt.Sprintf("http://localhost:5000/%d", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)

	return &order, nil
}

func (or *orderservice) FindAll(id *string) (*model.Order, error) {
	var order model.Order
	requestUrl := fmt.Sprintf("http://localhost:5000/getAll/%d", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)
	return &order, nil
}

func (or *orderservice) AddCart(order model.OrderItem) (*model.OrderItem, error) {
	json_data, _ := json.Marshal(&order)

	resp, err := http.NewRequest("POST", "http://localhost:5000/cart", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)
	return &order, nil
}

func (or *orderservice) Checkout(id *string, order model.Order) (*model.Order, error) {
	requestUrl := fmt.Sprintf("http://localhost:5000/checkout/%d", id)
	json_data, _ := json.Marshal(&order)

	resp, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(json_data))

	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)
	return &order, nil
}

func (or *orderservice) ListCart(userId *string) (*model.OrderItem, error) {
	var order model.OrderItem
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d", userId)
	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)
	return &order, nil
}

func (or *orderservice) DeleteCart(id *string) (*model.Order, error) {
	var order model.Order
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d", id)
	resp, err := http.NewRequest("DELETE", requestUrl, nil)

	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)
	return &order, nil
}

func (or *orderservice) ListNotPaid(id *string) (*model.Order, error) {
	var order model.Order
	requestUrl := fmt.Sprintf("http://localhost:5000/cart/%d/notpaid", id)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Body)
	return &order, nil
}
