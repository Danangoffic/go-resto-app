package main

import (
	"crypto/rand"
	"crypto/rsa"
	"resto-app/internal/database"
	"resto-app/internal/delivery/rest"
	"time"

	// "resto-app/internal/menu"
	// "resto-app/internal/order"
	mRepo "resto-app/internal/repository/menu"
	oRepo "resto-app/internal/repository/order"
	uRepo "resto-app/internal/repository/user"
	rUsecase "resto-app/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dsn = "root:@tcp(127.0.0.1:3306)/resto?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	// db_seed()
	e := echo.New()
	// localhost:14045/menu/food
	// e.GET("/menu", getMenu)

	db := database.GetDB(dsn)

	//load menu repository
	menuRepo := mRepo.GetRepository(db)
	// menuService := menu.NewService(menuRepo)

	//load order repository
	orderRepo := oRepo.GetRepository(db)

	secret := "AES256Key-32Character12345678910"
	// orderService := order.NewService(orderRepo, menuRepo)
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	accessTimeExp := 60 * time.Second

	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, accessTimeExp)
	if err != nil {
		panic(err)
	}

	restoCase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)
	// load handlers
	handler := rest.NewHandler(restoCase)
	// handler := rest.NewHandler(menuService, orderService)

	rest.LoadMiddelwares(e)
	rest.LoadRouters(e, handler)

	e.Logger.Fatal(e.Start(":14045"))
}
