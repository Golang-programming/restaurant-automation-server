package entity

import "gorm.io/gorm"

type OrderStatus string

var (
	Pending   OrderStatus = "pending"
	Confirmed OrderStatus = "confirmed"
	Cancelled OrderStatus = "cancelled"
	Making    OrderStatus = "making"
	Placed    OrderStatus = "placed"
	Billing   OrderStatus = "billing"
	Complete  OrderStatus = "complete"
	Failed    OrderStatus = "failed"
)

type Order struct {
	gorm.Model
	Status     OrderStatus `gorm:"type:enum('pending','confirmed','cancelled','making','placed','billing','complete','failed')"`
	TableID    uint        `gorm:"not null"`
	CustomerID uint        `gorm:"not null"`
	Table      Table       `gorm:"foreignKey:TableID"`
	Customer   Customer    `gorm:"foreignKey:CustomerID"`
}
