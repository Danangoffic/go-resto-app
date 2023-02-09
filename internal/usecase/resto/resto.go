package resto

import (
	"errors"
	"fmt"
	"resto-app/internal/model"
	"resto-app/internal/repository/menu"
	"resto-app/internal/repository/order"
	"resto-app/internal/repository/user"

	"github.com/google/uuid"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (r *restoUsecase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

func (s *restoUsecase) GetMenu(orderCode string) (model.MenuItem, error) {
	return s.menuRepo.GetMenu(orderCode)
}

func (s *restoUsecase) Order(request model.OrderMenuRequest) (model.Order, error) {
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	fmt.Printf("request order product : %v\n", request.OrderProducts)

	for i, orderPorudct := range request.OrderProducts {
		menuData, err := s.menuRepo.GetMenu(orderPorudct.OrderCode)
		if err != nil {
			fmt.Printf("failed get menu with %v", err)
			return model.Order{}, err
		}
		fmt.Printf("menu data %v\n", menuData)

		productOrderData[i] = model.ProductOrder{
			ID:         int(uuid.New().ID()),
			OrderCode:  menuData.OrderCode,
			Quantity:   orderPorudct.Quantity,
			TotalPrice: int64(menuData.Price) * int64(orderPorudct.Quantity),
			Status:     model.ProductOrderStatusPreparing,
		}
		fmt.Printf("product order data ke %d %v\n", i, productOrderData[i])
	}
	orderData := model.Order{
		UserID:        request.UserID,
		ID:            uuid.New().String(),
		Status:        model.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	fmt.Printf("order data %v", orderData)

	createdOrderData, err := s.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}
	return createdOrderData, nil
}

func (s *restoUsecase) GetOrderData(request model.GetOrderDataRequest) (model.Order, error) {
	orderData, err := s.orderRepo.GeteOrderData(request.OrderID)
	if err != nil {
		return model.Order{}, err
	}
	if orderData.UserID != request.UserID {
		return model.Order{}, errors.New("Unauthorized")
	}
	return orderData, nil
}

func (r *restoUsecase) RegisterUser(request model.RegisterRequest) (model.User, error) {
	isRegistered, err := r.userRepo.CheckRegistered(request.Username)
	if err != nil {
		return model.User{}, err
	}
	if isRegistered {
		return model.User{}, errors.New("User already registed!")
	}

	userHash, err := r.userRepo.GenerateUserHash(request.Password)
	if err != nil {
		return model.User{}, nil
	}

	userRegistered, err := r.userRepo.RegisterUser(model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	})
	if err != nil {
		return model.User{}, err
	}
	return userRegistered, nil
}

func (r *restoUsecase) Login(request model.LoginRequest) (model.UserSession, error) {
	userData, err := r.userRepo.GetUserData(request.Username)
	if err != nil {
		return model.UserSession{}, err
	}
	fmt.Printf("userdata : %v\n", userData)
	if userData.ID == "" {
		return model.UserSession{}, errors.New("User Not Found")
	}

	verified, err := r.userRepo.VerifyLogin(request.Username, request.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}
	fmt.Printf("is verified : %v\n", verified)

	if !verified {
		return model.UserSession{}, errors.New("can't verify user login")
	}

	userSession, err := r.userRepo.CreateUserSession(userData.ID)
	if err != nil {
		return model.UserSession{}, nil
	}
	fmt.Printf("user session : %v\n", userSession)
	return userSession, nil
}

func (r *restoUsecase) CheckSession(data model.UserSession) (userID string, err error) {
	userID, err = r.userRepo.CheckSession(data)
	if err != nil {
		return "", err
	}

	return userID, nil
}
