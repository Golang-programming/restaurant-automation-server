package entity

import (
	"gorm.io/gorm"
)

type BillStatus string

const (
	BillPending  BillStatus = "pending"
	BillPaid     BillStatus = "paid"
	BillRefunded BillStatus = "refunded"
)

type Bill struct {
	gorm.Model
	Subtotal   float64    `gorm:"type:decimal(10,2); not null"`
	Discount   float64    `gorm:"type:decimal(10,2); not null; default:0"`
	Total      float64    `gorm:"type:decimal(10,2); not null"`
	Paid       float64    `gorm:"type:decimal(10,2); not null; default:0"`
	Status     BillStatus `gorm:"type:bill_status; not null; default:'pending'"`
	OrderID    uint       `gorm:"not null"`
	Order      *Order     `gorm:"foreignKey:OrderID"`
	TableID    uint       `gorm:"not null"`
	Table      *Table     `gorm:"foreignKey:TableID"`
	CustomerID uint       `gorm:"not null"`
	Customer   *Customer  `gorm:"foreignKey:CustomerID"`
	Invoices   Invoice    `gorm:"foreignKey:BillID"`
}
