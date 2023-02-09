package order

import (
	"resto-app/internal/model"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) CreateOrder(order model.Order) (model.Order, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (r *orderRepo) GeteOrderData(orderID string) (model.Order, error) {
	var data model.Order

	if err := r.db.Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
