package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/saroj85/blog_api_go/pkg/models"
)

var db *gorm.DB

func GetDbUrl() string {

	// dbURL := "u170347778_bdbharti:Meraki2022@tcp(sql765.main-hosting.eu:3306)/u170347778_hourly?charset=utf8mb4&parseTime=True&loc=Local"
	db_url_dev := "root:123456789@/blog_db?charset=utf8&parseTime=True&loc=Local"
	db_url_prod := "sql6516409:ECm75aXYp8@tcp(sql6.freesqldatabase.com)/sql6516409?charset=utf8mb4&parseTime=True&loc=Local"
	mode := "production" //utils.GetDotEnvVariable("MODE")

	if mode == "production" {
		return db_url_prod
	}
	return db_url_dev

}

func Connect() {

	// root:123456789@/blog_db?charset=utf8&parseTime=True&loc=Local
	// gorm:gorm@tcp(3.7.144.210)/gorm?charset=utf8&parseTime=True&loc=Loca
	// nf_remote_two:iiu**syd^4ewdsccnxjsy^@tcp(3.7.144.210)/saroj_test?charset=utf8mb4&parseTime=True&loc=Local

	// nf_remote_two:iiu**syd^4ewdsccnxjsy^@tcp(3.7.144.210)/saroj_test?charset=utf8mb4&parseTime=True&loc=Local
	// Host: sql6.freesqldatabase.com
	// Database name: sql6516409
	// Database user: sql6516409
	// Database password: ECm75aXYp8
	// Port number: 3306

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
