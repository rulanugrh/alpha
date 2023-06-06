package order

import (
	"net/http"

	"github.com/rulanugrh/alpha/apigateway/entity/order"
)

type OrderInterfaces interface {
	Create(order order.OrderModel) (http.Response, error)
	FindId(id uint) (http.Response, error)
	FindAll(id uint) (http.Response, error)
	AddCart(order order.OrderItemModel) (http.Response, error)
	Checkout(id uint, order order.OrderModel) (http.Response, error)
	ListCart(userId uint) (http.Response, error)
	DeleteCart(id uint) (http.Response, error)
	ListNotPaid(id uint) (http.Response, error)
}
