package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

/*
* Database Initilization
 */
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("file not found")
	}
}

func DBConnection() *gorm.DB {

	db, err := gorm.Open(
		os.Getenv("DB_CONNECTION"),
		os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_DATABASE"))

	if err != nil {
		panic(err.Error())
	}
	return db
}
