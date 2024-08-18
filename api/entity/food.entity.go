package entity

import "gorm.io/gorm"

type FoodStatus string

const (
	FoodActive   FoodStatus = "active"
	FoodInactive FoodStatus = "inactive"
)

type Food struct {
	gorm.Model
	Name         string     `gorm:"type:varchar(255); not null"`
	Price        float64    `gorm:"type:decimal(10,2); not null"`
	Description  string     `gorm:"type:text; not null"`
	Images       []string   `gorm:"type:json"`
	Status       FoodStatus `gorm:"type:food_status; not null"`
	ParentFoodID *uint      `gorm:"index"`
	ParentFood   *Food      `gorm:"foreignKey:ParentFoodID"`
	MenuID       uint       `gorm:"not null"`
	Menu         *Menu      `gorm:"foreignKey:MenuID"`
	Category     string     `gorm:"type:varchar(100); not null"`
}
