package menu

import (
	"context"
	"resto-app/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=MenuRepository=MockMenuRepository -destination=../../mocks/menu_repository_mock.go -source=repository.go

type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error)
}
