package order

import (
	"net/http"

	"github.com/rulanugrh/graphql/graph/model"
)

type OrderInterfaces interface {
	Create(order model.Order) (http.Response, error)
	FindId(id *int) (http.Response, error)
	FindAll(id *int) (http.Response, error)
	AddCart(order model.OrderItem) (http.Response, error)
	Checkout(id *int, order model.Order) (http.Response, error)
	ListCart(userId *int) (http.Response, error)
	DeleteCart(id *int) (http.Response, error)
	ListNotPaid(id *int) (http.Response, error)
}
