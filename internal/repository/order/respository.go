package order

import (
	"context"
	"resto-app/internal/model"
)

type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GeteOrderData(ctx context.Context, ID string) (model.Order, error)
}
