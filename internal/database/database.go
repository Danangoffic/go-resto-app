package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func GetDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbAddress), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	db_seed(db)
	return db
}
