package database

import (
	"fmt"

	"resto-app/internal/model"
	"resto-app/internal/model/constant"

	"gorm.io/gorm"
)

func db_seed(db *gorm.DB) {
	migrate_db(db)

	foodMenu := []model.MenuItem{
		{
			Name:      "Kwetiaw",
			OrderCode: "kwetiaw-1",
			Price:     12000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasi-goreng-1",
			Price:     16000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "teh-121",
			Price:     5500,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Teh Panas",
			OrderCode: "teh-122",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu, drinkMenu)

	err := db.First(&model.MenuItem{}).Error
	if err == gorm.ErrRecordNotFound {
		db.Create(foodMenu)
		db.Create(drinkMenu)
		fmt.Printf("food menu inserted %v\n", foodMenu)
		fmt.Printf("drink menu inserted %v\n", drinkMenu)
	}

}

func migrate_db(db *gorm.DB) {
	if db.Migrator().HasTable(&model.ProductOrder{}) {
		db.Migrator().DropTable(&model.ProductOrder{})
	}
	if db.Migrator().HasTable(&model.Order{}) {
		db.Migrator().DropTable(&model.Order{})
	}
	db.AutoMigrate(&model.ProductOrder{}, &model.Order{})
}
