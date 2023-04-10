package order

import "github.com/rulanugrh/alpha/order/entities/domain"

type OrderRepository interface {
	Create(order domain.Order) (domain.Order, error)
	FindID(id uint) (domain.Order, error)
	FindAll(id uint) ([]domain.Order, error)
	AddCart(orderItem domain.OrderItem) (domain.OrderItem, error)
	Checkout(id uint, order domain.Order) (domain.Order, error)
	ListCart(userId uint) ([]domain.OrderItem, error)
	DeleteCart(id uint) error
}
