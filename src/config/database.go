package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	var err error
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
