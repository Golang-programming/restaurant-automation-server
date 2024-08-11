package entity

import (
	"time"

	orderEntity "github.co/golang-programming/restaurant/api/order/entity"
	"gorm.io/gorm"
)

type PaymentMethod string
type PaymentStatus string

const (
	// PaymentMethod values
	CreditCard   PaymentMethod = "CreditCard"
	DebitCard    PaymentMethod = "DebitCard"
	PayPal       PaymentMethod = "PayPal"
	BankTransfer PaymentMethod = "BankTransfer"

	// PaymentStatus values
	Pending   PaymentStatus = "Pending"
	Completed PaymentStatus = "Completed"
	Failed    PaymentStatus = "Failed"
)

type Invoice struct {
	gorm.Model
	PaymentMethod  PaymentMethod     `gorm:"type:varchar(255);not null"`
	PaymentStatus  PaymentStatus     `gorm:"type:varchar(255);not null"`
	PaymentDueDate time.Time         `gorm:"type:date;not null"`
	OrderID        uint              `gorm:"not null"`
	Order          orderEntity.Order `gorm:"foreignKey:OrderID;references:ID"`
}
