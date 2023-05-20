package config

import (
	"fmt"
	"server/models"

	"log"
	// "server/seeders"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DATA BASE ERROR")
	}

	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Borrow{})

	// err = seeders.SeedBooks(db)

	// if err != nil {
	// 	panic("Seed Error") // hanya sekali seeder
	// }

	DB = db
	log.Println("Succes Connect database")
}
