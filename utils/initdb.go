package utils

import (
	"fmt"
	"os"
	"web-api-sql/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	DSN := os.Getenv("DB_URI")
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	DB = db

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	DB.AutoMigrate(&models.Todo{})
}
