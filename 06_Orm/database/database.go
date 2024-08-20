package database

import (
	"fmt"
	"log"
	"orm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_Username = "root"
	DB_Password = ""
	DB_Name     = "test"
	DB_Host     = "localhost"
	DB_Port     = "3306"
)

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_Username,
		DB_Password,
		DB_Host,
		DB_Port,
		DB_Name,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Occured When Connection to Database", err)
	}
	log.Println("Successfully Connected!")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.Package{})
	if err != nil {
		log.Fatal("migration failed!", err)
	}
	log.Println("migration successful")
}
