package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB(
	MYSQL_HOST string,
	MYSQL_USER string,
	MYSQL_PORT string,
	MYSQL_DB_NAME string,
	MYSQL_PASSWORD string,
) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MYSQL_USER,
		MYSQL_PASSWORD,
		MYSQL_HOST,
		MYSQL_PORT,
		MYSQL_DB_NAME,
	)

	fmt.Println("Connecting to MYSQL database")

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprint("ERROR: ", err))
	}

	fmt.Println("Connected")

	db = database
}

func GetDb() *gorm.DB {
	if db == nil {
		StartDB(
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DB_NAME"),
			os.Getenv("MYSQL_PASSWORD"),
		)
	}

	return db
}
