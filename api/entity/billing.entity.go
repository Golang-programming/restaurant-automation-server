package entity

import (
	"gorm.io/gorm"
)

type BillStatus string

const (
	Pending  BillStatus = "pending"
	Paid     BillStatus = "paid"
	Refunded BillStatus = "refunded"
)

type Bill struct {
	gorm.Model
	Subtotal   float64    `gorm:"type:decimal(10,2); not null"`
	Discount   float64    `gorm:"type:decimal(10,2); not null; default:0"`
	Total      float64    `gorm:"type:decimal(10,2); not null"`
	Paid       float64    `gorm:"type:decimal(10,2); not null; default:0"`
	Status     BillStatus `gorm:"type:enum('pending','paid', 'refunded'); not null; default:'pending'"`
	OrderID    uint       `gorm:"not null"`
	Order      Order      `gorm:"foreignKey:OrderID"`
	CustomerID uint       `gorm:"not null"`
	Customer   Customer   `gorm:"foreignKey:CustomerID"`
	InvoiceID  uint       `gorm:"not null"`
	Invoice    Invoice    `gorm:"foreignKey:InvoiceID"`
}
