package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dbUrl := os.Getenv("DB_URL")
	envMode := os.Getenv("ENV_MODE")

	var database *gorm.DB
	var err error

	switch envMode {
	case "development":
		database, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
		})

	case "production":
		database, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{
			SkipDefaultTransaction: true,
		})

	default:
		database, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
		})
	}

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to the database")
	}

	fmt.Println("Connected to the database✅")

	return database
}
