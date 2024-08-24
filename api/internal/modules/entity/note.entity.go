package entity

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Description string     `gorm:"type:varchar(255);not null"`
	OrderItemID uint       `gorm:"not null"`
	OrderItem   *OrderItem `gorm:"foreignKey:OrderItemID"`
}
