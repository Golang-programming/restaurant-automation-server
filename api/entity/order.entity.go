package entity

import "gorm.io/gorm"

type OrderStatus string
type OrderItemStatus string
type OrderItemProductType string

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

const (
	OrderItemFood OrderItemProductType = "food"
	OrderItemDeal OrderItemProductType = "deal"
)

type Order struct {
	gorm.Model
	Status     OrderStatus `gorm:"type:order_status; not null; default:"`
	TableID    uint        `gorm:"not null"`
	CustomerID uint        `gorm:"not null"`
	Table      *Table      `gorm:"foreignKey:TableID"`
	Customer   *Customer   `gorm:"foreignKey:CustomerID"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	ProductID   uint                 `gorm:"not null"`
	ProductType OrderItemProductType `gorm:"type:order_item_product_type;not null;default:'food'"`
	OrderID     uint                 `gorm:"not null"`
	Order       *Order               `gorm:"foreignKey:OrderID"`
	Quantity    int                  `gorm:"type:int;not null"`
	Status      OrderItemStatus      `gorm:"type:order_item_status;not null;default:'ordering'"`
	Notes       []Note               `gorm:"foreignKey:OrderItemID"`
}
