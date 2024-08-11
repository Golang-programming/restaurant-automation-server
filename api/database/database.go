package database

import (
	"log"

	foodEntity "github.co/golang-programming/restaurant/api/food/entity"
	invoiceEntity "github.co/golang-programming/restaurant/api/invoice/entity"
	menuEntity "github.co/golang-programming/restaurant/api/menu/entity"
	noteEntity "github.co/golang-programming/restaurant/api/note/entity"
	orderEntity "github.co/golang-programming/restaurant/api/order/entity"
	tableEntity "github.co/golang-programming/restaurant/api/table/entity"
	userEntity "github.co/golang-programming/restaurant/api/user/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {

	dsn := "root:password@tcp(127.0.0.1:3307)/restaurant_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	DB = db
	/*
		DB.AutoMigrate(&foodEntity.Food{})
		DB.AutoMigrate(&userEntity.User{})
		DB.AutoMigrate(&invoiceEntity.Invoice{})
		DB.AutoMigrate(&orderEntity.Order{})
		DB.AutoMigrate(&orderEntity.OrderItem{})
		DB.AutoMigrate(&tableEntity.Table{})
		DB.AutoMigrate(&noteEntity.Note{})
		DB.AutoMigrate(&menuEntity.Menu{}) */

	DB.AutoMigrate(
		&foodEntity.Food{},
		&userEntity.User{},
		&invoiceEntity.Invoice{},
		&orderEntity.Order{},
		&orderEntity.OrderItem{},
		&tableEntity.Table{},
		&noteEntity.Note{},
		&menuEntity.Menu{},
	)

}
