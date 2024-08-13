package entity

import "gorm.io/gorm"

type FoodStatus string

const (
	Active   FoodStatus = "active"
	Inactive FoodStatus = "inactive"
)

type Food struct {
	gorm.Model
	Name        string      `gorm:"type:varchar(255); not null"`
	Price       float64     `gorm:"type:decimal(10,2); not null"`
	Description string      `gorm:"type:text; not null"`
	Images      []string    `gorm:"type:json; not null"`
	Status      FoodStatus  `gorm:"type:enum('active', 'inactive'); not null"`
	MenuID      uint        `gorm:"not null"`
	Menu        Menu        `gorm:"foreignKey:MenuID"`
	OrderItems  []OrderItem `gorm:"foreignKey:FoodID"`
	DealItems   []DealItem  `gorm:"foreignKey:FoodID"`
}
