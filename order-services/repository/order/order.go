package order

import (
	"errors"

	"github.com/rulanugrh/alpha/order/config"
	"github.com/rulanugrh/alpha/order/entities/domain"
)

type orderrepo struct{}

func NewOrderRepository() OrderRepository {
	return &orderrepo{}
}

func (or *orderrepo) Create(order domain.Order) (domain.Order, error) {
	err := config.DB.Create(&order).Error
	if err != nil {
		return domain.Order{}, errors.New("cant create order")
	}

	return order, nil
}

func (or *orderrepo) FindID(id uint) (domain.Order, error) {
	var order domain.Order
	err := config.DB.First(&order, id).Error
	if err != nil {
		return domain.Order{}, errors.New("cant find order by this id")
	}

	return order, nil
}

func (or *orderrepo) FindAll() ([]domain.Order, error) {
	var order []domain.Order
	err := config.DB.Find(&order).Error
	if err != nil {
		return []domain.Order{}, errors.New("cant find all order")
	}

	return order, nil
}
