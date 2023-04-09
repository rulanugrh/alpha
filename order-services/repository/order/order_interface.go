package order

import "github.com/rulanugrh/alpha/order/entities/domain"

type OrderRepository interface {
	Create(order domain.Order) (domain.Order, error)
	FindID(id uint) (domain.Order, error)
	FindAll() ([]domain.Order, error)
	//	PayOrder(id uint, orderItem domain.OrderItem) (domain.OrderItem, error)
}
