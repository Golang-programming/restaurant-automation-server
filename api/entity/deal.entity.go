package entity

import (
	"time"

	"gorm.io/gorm"
)

type Deal struct {
	gorm.Model
	ValidDue time.Time  `json:"valid_due"`
	Status   string     `json:"status"`
	Items    []DealItem `json:"items" gorm:"foreignKey:DealID;constraint:OnDelete:CASCADE;"` // Relationship with DealItem
}

type DealItem struct {
	gorm.Model
	DealID   uint `json:"deal_id"`
	FoodID   uint `json:"food_id"`
	Quantity int  `json:"quantity"`
}
