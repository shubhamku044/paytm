package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() *gorm.DB {
	dbUrl := os.Getenv("DB_URL")

	database, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to the database")
	}

	fmt.Println("Connected to the database✅")

	return database
}
