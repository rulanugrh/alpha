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
	err := config.DB.Preload("User").First(&order, id).Error
	if err != nil {
		return domain.Order{}, errors.New("cant find order by this id")
	}

	return order, nil
}

func (or *orderrepo) FindAll(id uint) ([]domain.Order, error) {
	var order []domain.Order
	err := config.DB.Preload("User").Where("user_id = ?", id).Find(&order).Error
	if err != nil {
		return []domain.Order{}, errors.New("cant find all order")
	}

	return order, nil
}

func (or *orderrepo) AddCart(order domain.OrderItem) (domain.OrderItem, error) {
	err := config.DB.Create(&order).Error
	if err != nil {
		return domain.OrderItem{}, errors.New("cant add cart")
	}

	return order, nil
}

func (or *orderrepo) Checkout(id uint, order domain.Order) (domain.Order, error) {
	var orderUpdates domain.Order
	err := config.DB.Model(&order).Where("id = ?", id).Updates(&orderUpdates).Error
	if err != nil {
		return domain.Order{}, errors.New("cant checkout")
	}

	return orderUpdates, nil
}

func (or *orderrepo) ListCart(id uint) ([]domain.OrderItem, error) {
	var orderItem []domain.OrderItem

	err := config.DB.Preload("Books").Preload("Orders").Preload("User").Where("user_id = ?", id).Find(&orderItem).Error
	if err != nil {
		return []domain.OrderItem{}, errors.New("cant find all cart")
	}

	return orderItem, nil
}

func (or *orderrepo) DeleteCart(id uint) error {
	var orderItem domain.OrderItem
	err := config.DB.Where("id = ?", id).Delete(&orderItem).Error
	if err != nil {
		return errors.New("cant delete cart")
	}

	return nil
}

func (or *orderrepo) ListNotPaid(id uint) ([]domain.Order, error) {
	var order []domain.Order
	err := config.DB.Preload("User").Where("user_id = ?", id).Where("paid = ?", false).Find(&order).Error
	if err != nil {
		return []domain.Order{}, errors.New("cant list all not paid")
	}

	return order, nil
}
