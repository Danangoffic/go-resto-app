package menu

import (
	"context"
	"resto-app/internal/model"
	"resto-app/internal/tracing"

	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (r *menuRepo) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	var menuData []model.MenuItem
	err := r.db.WithContext(ctx).Where(model.MenuItem{Type: menuType}).Find(&menuData).Error
	if err != nil {
		return nil, err
	}
	return menuData, nil
}

func (r *menuRepo) GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenu")
	defer span.End()

	var menuData model.MenuItem
	if err := r.db.WithContext(ctx).Where(model.MenuItem{OrderCode: orderCode}).First(&menuData).Error; err != nil {
		return menuData, err
	}

	return menuData, nil
}
