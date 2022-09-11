package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/saroj85/blog_api_go/pkg/models"
	"github.com/saroj85/blog_api_go/pkg/utils"
)

var db *gorm.DB

func GetDbUrl() string {

	mode := utils.GetDotEnvVariable("MODE")

	fmt.Println("APP MODE:", mode)
	if mode == "production" {
		return utils.GetDotEnvVariable("DB_CONNECTION_URL_PROD")
	} else {
		return utils.GetDotEnvVariable("DB_CONNECTION_URL_DEV")
	}
}

func Connect() {

	fmt.Println("connecting to database...")

	db_url := GetDbUrl()
	d, err := gorm.Open("mysql", db_url)
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
