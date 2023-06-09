package order

import (
	"github.com/rulanugrh/graphql/graph/model"
)

type OrderInterfaces interface {
	Create(order model.Order) (*model.Response, error)
	FindId(id *string) (*model.Response, error)
	FindAll(id *string) (*model.Response, error)
	AddCart(order model.OrderItem) (*model.Response, error)
	Checkout(id *string, order model.Order) (*model.Response, error)
	ListCart(userId *string) (*model.Response, error)
	DeleteCart(id *string) (*model.Response, error)
	ListNotPaid(id *string) (*model.Response, error)
}
