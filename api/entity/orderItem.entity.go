package entity

import "gorm.io/gorm"

type OrderItemStatus string

var (
	Ordering OrderItemStatus = "ordering"
	Making   OrderItemStatus = "making"
	Placed   OrderItemStatus = "placed"
)

type OrderItem struct {
	gorm.Model
	FoodID   uint            `gorm:"not null"`
	Food     Food            `gorm:"foreignKey:FoodID"`
	OrderID  uint            `gorm:"not null"`
	Order    Order           `gorm:"foreignKey:OrderID"`
	Quantity int             `gorm:"type:int;not null"`
	Status   OrderItemStatus `gorm:"type:enum('ordering','making','placed');not null;default:'ordering'"`
}
