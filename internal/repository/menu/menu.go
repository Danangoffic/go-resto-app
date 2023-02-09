package menu

import (
	"resto-app/internal/model"

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

func (r *menuRepo) GetMenuList(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem
	err := r.db.Where(model.MenuItem{Type: menuType}).Find(&menuData).Error
	if err != nil {
		return nil, err
	}
	return menuData, nil
}

func (r *menuRepo) GetMenu(orderCode string) (model.MenuItem, error) {
	var menuData model.MenuItem
	if err := r.db.Where(model.MenuItem{OrderCode: orderCode}).First(&menuData).Error; err != nil {
		return menuData, err
	}

	return menuData, nil
}
