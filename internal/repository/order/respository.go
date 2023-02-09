package order

import "resto-app/internal/model"

type Repository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GeteOrderData(orderID string) (model.Order, error)
}
