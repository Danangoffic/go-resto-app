package resto

import "resto-app/internal/model"

type Usecase interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)

	Order(request model.OrderMenuRequest) (model.Order, error)
	GetOrderData(request model.GetOrderDataRequest) (model.Order, error)

	RegisterUser(request model.RegisterRequest) (model.User, error)
	Login(request model.LoginRequest) (model.UserSession, error)
	CheckSession(data model.UserSession) (userID string, err error)
}
