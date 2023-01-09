package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.Logger.Fatal(e.Start(":14045"))
}

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Kwetiaw",
			OrderCode: "kwetiaw",
			Price:     12000,
		},
		{
			Name:      "Nasi Goreng",
			OrderCode: "NG-1",
			Price:     16000,
		},
	}
	return c.JSON(200, foodMenu)
}
