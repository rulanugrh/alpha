package order

import (
	"github.com/rulanugrh/graphql/graph/model"
)

type OrderInterfaces interface {
	Create(order model.Order) (*model.Order, error)
	FindId(id *string) (*model.Order, error)
	FindAll(id *string) (*model.Order, error)
	AddCart(order model.OrderItem) (*model.OrderItem, error)
	Checkout(id *string, order model.Order) (*model.Order, error)
	ListCart(userId *string) (*model.OrderItem, error)
	DeleteCart(id *string) (*model.Order, error)
	ListNotPaid(id *string) (*model.Order, error)
}
