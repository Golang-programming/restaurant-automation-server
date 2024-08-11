package entity

import (
	foodEntity "github.co/golang-programming/restaurant/api/food/entity"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	Quantity  int             `gorm:"type:int;not null"`
	UnitPrice float64         `gorm:"type:decimal(10,2);not null"`
	FoodID    uint            `gorm:"not null"`
	OrderID   uint            `gorm:"not null"`
	Food      foodEntity.Food `gorm:"foreignKey:FoodID"`
	Order     Order           `gorm:"foreignKey:OrderID"`
}
