package entity

import (
	"time"

	"gorm.io/gorm"
)

type Deal struct {
	gorm.Model
	Name     string     `gorm:"type:varchar(255); not null"`
	ValidDue time.Time  `gorm:"not null"`
	Status   string     `gorm:"type:enum('active', 'inactive'); not null"`
	Items    []DealItem `gorm:"foreignKey:DealID;constraint:OnDelete:CASCADE;"`
}

type DealItem struct {
	gorm.Model
	DealID   uint `gorm:"not null"`
	Deal     Deal `gorm:"foreignKey:DealID"`
	FoodID   uint `gorm:"not null"`
	Food     Food `gorm:"foreignKey:FoodID"`
	Quantity int  `gorm:"type:int;not null"`
}
