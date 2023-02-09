package menu

import (
	"resto-app/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetMenuList(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem
	err := r.db.Where(model.MenuItem{Type: menuType}).Find(&menuData).Error
	if err != nil {
		return nil, err
	}
	return menuData, nil
}

func (r *repository) GetMenu(orderCode string) (model.MenuItem, error) {
	var menuData model.MenuItem
	if err := r.db.Where(model.MenuItem{OrderCode: orderCode}).First(&menuData).Error; err != nil {
		return menuData, err
	}

	return menuData, nil
}
