package order

import (
	"errors"

	"github.com/rulanugrh/alpha/order/entities/domain"
	"github.com/rulanugrh/alpha/order/entities/web"
	"github.com/rulanugrh/alpha/order/repository/order"
)

type orderservice struct {
	orders order.OrderRepository
}

func NewOrderServices(order order.OrderRepository) OrderService {
	return &orderservice{
		orders: order,
	}
}

func (or *orderservice) CreateOrder(order domain.Order) (web.OrderResponse, error) {
	ors, err := or.orders.Create(order)
	if err != nil {
		return web.OrderResponse{}, errors.New("cant create orders")
	}

	return web.OrderResponse{
		OrderID:  int(ors.ID),
		Subtotal: order.Total,
		Paid:     ors.Paid,
	}, nil
}

func (or *orderservice) FindId(id uint) (web.OrderResponse, error) {
	order, err := or.orders.FindID(id)
	if err != nil {
		return web.OrderResponse{}, errors.New("cant find order")
	}

	return web.OrderResponse{
		OrderID:  int(id),
		Subtotal: order.Total,
		Paid:     order.Paid,
	}, nil
}

func (or *orderservice) FindAll(id uint) ([]domain.Order, error) {
	orders, err := or.orders.FindAll(id)
	if err != nil {
		return []domain.Order{}, errors.New("can find orders")
	}

	return orders, nil
}

func (or *orderservice) AddCart(orderItem domain.OrderItem) (web.OrderItemResponse, error) {
	orders, err := or.orders.AddCart(orderItem)
	if err != nil {
		return web.OrderItemResponse{}, errors.New("cant create orderItem")
	}

	return web.OrderItemResponse{
		Quantity: orders.Quantity,
		BookID:   orders.BookID,
		Subtotal: orders.Subtotal,
	}, nil
}

func (or *orderservice) Checkout(id uint, order domain.Order) (web.OrderResponse, error) {
	ors, err := or.orders.Checkout(id, order)
	if err != nil {
		return web.OrderResponse{}, errors.New("cant checkout")
	}

	return web.OrderResponse{
		OrderID:  int(ors.ID),
		Paid:     ors.Paid,
		Subtotal: ors.Total,
	}, nil
}

func (or *orderservice) ListCart(id uint) ([]domain.OrderItem, error) {
	orders, err := or.orders.ListCart(id)
	if err != nil {
		return []domain.OrderItem{}, errors.New("cant find all cart")
	}

	return orders, nil
}

func (or *orderservice) DeleteCart(id uint) error {
	err := or.orders.DeleteCart(id)
	if err != nil {
		return errors.New("cant delete cart")
	}

	return nil
}

func (or *orderservice) ListNotPaid(id uint) ([]domain.Order, error) {
	orders, err := or.orders.ListNotPaid(id)
	if err != nil {
		return []domain.Order{}, errors.New("cant find not paid")
	}

	return orders, nil
}
