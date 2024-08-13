package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	TotalGuests int     `gorm:"type:int;not null"`
	Email       *string `gorm:"type:varchar(100)"`
	Phone       *string `gorm:"type:varchar(50)"`
	Orders      []Order `gorm:"foreignKey:CustomerID"` // One-to-many relationship with Order
}
