package entity

import (
	"time"

	"gorm.io/gorm"
)

type Deal struct {
	gorm.Model
	ValidDue time.Time  `gorm:"type:timestamp"`
	Title    string     `json:"type:varchar(255)"`
	Status   string     `json:"status"`
	MenuID   uint       `gorm:"not null"`
	Menu     *Menu      `gorm:"foreignKey:MenuID"`
	Items    []DealItem `json:"items" gorm:"foreignKey:DealID;constraint:OnDelete:CASCADE;"`
}

type DealItem struct {
	gorm.Model
	DealID   uint  `gorm:"not null"`
	Deal     *Deal `gorm:"foreignKey:DealID"`
	FoodID   uint  `gorm:"not null"`
	Food     *Food `gorm:"foreignKey:FoodID"`
	Quantity int   `gorm:"type:int;not null"`
}
