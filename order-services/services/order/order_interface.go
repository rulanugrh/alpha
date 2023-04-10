package order

import (
	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
)

type OrderService interface {
	CreateOrder(order domain.Order) (web.OrderResponse, error)
	FindId(id uint) (web.OrderResponse, error)
	FindAll(id uint) ([]domain.Order, error)
	AddCart(orderItem domain.OrderItem) (web.OrderItemResponse, error)
	Checkout(id uint, order domain.Order) (web.OrderResponse, error)
	ListCart(id uint) ([]domain.OrderItem, error)
	DeleteCart(id uint) error
}
