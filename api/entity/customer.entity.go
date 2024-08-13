package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	TotalGuests int     `gorm:"type:int;not null"`
	Name        string  `gorm:"type:varchar(255);null"`
	Orders      []Order `gorm:"foreignKey:CustomerID"`
}
