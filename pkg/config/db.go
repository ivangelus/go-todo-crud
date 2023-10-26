package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func connect() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Unable to load env file")
		os.Exit(1)
	}

	d, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/task-db?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS")))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	connect()
	return db
}
