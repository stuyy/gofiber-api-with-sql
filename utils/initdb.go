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
	DB, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	DB.AutoMigrate(&models.Todo{})
}
