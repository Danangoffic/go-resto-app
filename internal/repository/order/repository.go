package order

import (
	"context"
	"resto-app/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=OrderRepository=MockOrderRepository -destination=../../mocks/order_repository_mock.go -source=repository.go
type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GeteOrderData(ctx context.Context, ID string) (model.Order, error)
}
