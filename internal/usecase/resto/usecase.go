package resto

import (
	"context"
	"resto-app/internal/model"
)

type Usecase interface {
	GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error)
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)

	Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error)
	GetOrderData(ctx context.Context, request model.GetOrderDataRequest) (model.Order, error)

	RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error)
	Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error)
	CheckSession(ctx context.Context, data model.UserSession) (userID string, err error)
}
