package database

import (
	"log"

	"github.com/saifoelloh/minimalism_api/api/author"
	"github.com/saifoelloh/minimalism_api/api/book"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Host     = "172.18.0.2"
	Port     = 5432
	User     = "root"
	Password = "admin123"
	DBname   = "bookstore"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:admin123@tcp(172.18.0.2)/uwu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func Migration() {
	DB.AutoMigrate(&author.Author{}, &book.Book{})
}
