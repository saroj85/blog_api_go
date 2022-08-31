package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/saroj85/blog_api_go/pkg/models"
)

var db *gorm.DB

func Connect() {

	fmt.Println("connecting to database...")
	d, err := gorm.Open("mysql", "root:123456789@/blog_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected Successfully to database...")
	db = d

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.PostCategory{})

	// defer db.Close()

}

func GetDb() *gorm.DB {
	return db
}

func CloseDb(connection *gorm.DB) {

	fmt.Println("CONNECTION CLOSED")
	db := connection.DB()
	db.Close()
}
