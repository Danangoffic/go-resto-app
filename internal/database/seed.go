package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
