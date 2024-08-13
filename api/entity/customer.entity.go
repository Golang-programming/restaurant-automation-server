package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255); not null"`
	PhoneNumber string  `gorm:"type:varchar(50); not null;unique"`
	TotalGuests int     `gorm:"type:int;not null"`
	Orders      []Order `gorm:"foreignKey:CustomerID"`
	Bills       []Bill  `gorm:"foreignKey:CustomerID"`
}
