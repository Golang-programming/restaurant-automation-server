package entity

import "gorm.io/gorm"

type OrderStatus string

const (
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
	Table      Table       `gorm:"foreignKey:TableID"`
	CustomerID uint        `gorm:"not null"`
	Customer   Customer    `gorm:"foreignKey:CustomerID"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	Bill       Bill        `gorm:"foreignKey:OrderID"`
}
