package entity

import "gorm.io/gorm"

type OrderStatus string
type OrderItemStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderConfirmed OrderStatus = "confirmed"
	OrderCancelled OrderStatus = "cancelled"
	OrderMaking    OrderStatus = "making"
	OrderPlaced    OrderStatus = "placed"
	OrderBilling   OrderStatus = "billing"
	OrderComplete  OrderStatus = "complete"
	OrderFailed    OrderStatus = "failed"
)

const (
	OrderItemOrdering OrderItemStatus = "ordering"
	OrderItemMaking   OrderItemStatus = "making"
	OrderItemPlaced   OrderItemStatus = "placed"
)

type Order struct {
	gorm.Model
	Status     OrderStatus `gorm:"type:enum('pending','confirmed','cancelled','making','placed','billing','complete','failed')"`
	TableID    uint        `gorm:"not null"`
	CustomerID uint        `gorm:"not null"`
	UserID     uint        `gorm:"not null"` // The user who created the order
	User       *User       `gorm:"foreignKey:UserID"`
	Table      *Table      `gorm:"foreignKey:TableID"`
	Customer   *Customer   `gorm:"foreignKey:CustomerID"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	FoodID   uint   `gorm:"not null"`
	Food     *Food  `gorm:"foreignKey:FoodID"`
	OrderID  uint   `gorm:"not null"`
	Order    *Order `gorm:"foreignKey:OrderID"`
	Quantity int    `gorm:"type:int;not null"`
	Status   string `gorm:"type:enum('ordering','making','placed');not null;default:'ordering'"`
	Notes    []Note `gorm:"foreignKey:OrderItemID"` // One-to-many with Notes
}
