package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:@tcp(127.0.0.1:3306)/resto?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	// db_seed()
	e := echo.New()
	// localhost:14045/menu/food
	e.GET("/menu", getMenu)

	e.Logger.Fatal(e.Start(":14045"))
}

type MenuType string

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

type MenuItem struct {
	Name      string `json:"name"`
	OrderCode string `json:"order_code"`
	Price     int    `json:"price"`
	Type      string `json:"type"`
}

func db_seed() {
	foodMenu := []MenuItem{
		{
			Name:      "Kwetiaw",
			OrderCode: "kwetiaw-1",
			Price:     12000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasi-goreng-1",
			Price:     16000,
			Type:      MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "teh-121",
			Price:     5500,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Teh Panas",
			OrderCode: "teh-122",
			Price:     5000,
			Type:      MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu, drinkMenu)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&MenuItem{})

	err = db.First(&MenuItem{}).Error
	if err == gorm.ErrRecordNotFound {
		db.Create(foodMenu)
		db.Create(drinkMenu)
	}

}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var menuData []MenuItem

	db.Where(MenuItem{Type: menuType}).Find(&menuData).Order("name ASC")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":   menuData,
		"status": "success",
	})
}
