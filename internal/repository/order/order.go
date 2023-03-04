package order

import (
	"context"
	"resto-app/internal/model"
	"resto-app/internal/tracing"

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

func (r *orderRepo) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateOrder")
	defer span.End()

	if err := r.db.WithContext(ctx).Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (r *orderRepo) GeteOrderData(ctx context.Context, orderID string) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateOrder")
	defer span.End()

	var data model.Order

	if err := r.db.WithContext(ctx).Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
