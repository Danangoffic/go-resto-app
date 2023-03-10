package order

import (
	"resto-app/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderData(orderID string) (model.Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateOrder(order model.Order) (model.Order, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (r *repository) GetOrderData(orderID string) (model.Order, error) {
	var data model.Order

	if err := r.db.Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
