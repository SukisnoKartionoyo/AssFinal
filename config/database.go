package config

import (
	"assfinal/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB is global variable for Database connection
var DB *gorm.DB

//ConnectDB is for coonnect to DB
func ConnectDB() {
	dsn := "go-test:test123@tcp(192.168.91.134:3306)/test-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&models.Project{}, &models.Label{}, &models.User{}, &models.Task{}, &models.Status{})
	fmt.Println("Connected to Database")

	DB = db
}
