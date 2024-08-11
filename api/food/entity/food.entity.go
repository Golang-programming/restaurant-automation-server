package entity

import (
	menuEntity "github.co/golang-programming/restaurant/api/menu/entity"
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name      string          `gorm:"type:varchar(255);not null"`
	Price     float64         `gorm:"type:decimal(10,2);not null"`
	FoodImage string          `gorm:"type:varchar(255);"`
	MenuID    uint            `gorm:"not null"`
	Menu      menuEntity.Menu `gorm:"foreignKey:MenuID"`
}
