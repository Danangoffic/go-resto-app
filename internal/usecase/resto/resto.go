package resto

import (
	"context"
	"errors"
	"fmt"
	"resto-app/internal/model"
	"resto-app/internal/repository/menu"
	"resto-app/internal/repository/order"
	"resto-app/internal/repository/user"
	"resto-app/internal/tracing"

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

func (r *restoUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	return r.menuRepo.GetMenuList(ctx, menuType)
}

func (s *restoUsecase) GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenu")
	defer span.End()
	return s.menuRepo.GetMenu(ctx, orderCode)
}

func (s *restoUsecase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "Order")
	defer span.End()

	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	fmt.Printf("request order product : %v\n", request.OrderProducts)

	for i, orderPorudct := range request.OrderProducts {
		menuData, err := s.menuRepo.GetMenu(ctx, orderPorudct.OrderCode)
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

	createdOrderData, err := s.orderRepo.CreateOrder(ctx, orderData)
	if err != nil {
		return model.Order{}, err
	}
	return createdOrderData, nil
}

func (s *restoUsecase) GetOrderData(ctx context.Context, request model.GetOrderDataRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderData")
	defer span.End()

	orderData, err := s.orderRepo.GeteOrderData(ctx, request.OrderID)
	if err != nil {
		return model.Order{}, err
	}
	if orderData.UserID != request.UserID {
		return model.Order{}, errors.New("Unauthorized")
	}
	return orderData, nil
}

func (r *restoUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	isRegistered, err := r.userRepo.CheckRegistered(ctx, request.Username)
	if err != nil {
		return model.User{}, err
	}
	if isRegistered {
		return model.User{}, errors.New("User already registed!")
	}

	userHash, err := r.userRepo.GenerateUserHash(ctx, request.Password)
	if err != nil {
		return model.User{}, nil
	}

	userRegistered, err := r.userRepo.RegisterUser(ctx, model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	})
	if err != nil {
		return model.User{}, err
	}
	return userRegistered, nil
}

func (r *restoUsecase) Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "Login")
	defer span.End()

	userData, err := r.userRepo.GetUserData(ctx, request.Username)
	if err != nil {
		return model.UserSession{}, err
	}
	fmt.Printf("userdata : %v\n", userData)
	if userData.ID == "" {
		return model.UserSession{}, errors.New("User Not Found")
	}

	verified, err := r.userRepo.VerifyLogin(ctx, request.Username, request.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}
	fmt.Printf("is verified : %v\n", verified)

	if !verified {
		return model.UserSession{}, errors.New("can't verify user login")
	}

	userSession, err := r.userRepo.CreateUserSession(ctx, userData.ID)
	if err != nil {
		return model.UserSession{}, nil
	}
	fmt.Printf("user session : %v\n", userSession)
	return userSession, nil
}

func (r *restoUsecase) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	userID, err = r.userRepo.CheckSession(ctx, data)
	if err != nil {
		return "", err
	}

	return userID, nil
}
