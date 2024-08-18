package entity

import (
	"time"

	"gorm.io/gorm"
)

type CustomerStatus string

const (
	CustomerActive   CustomerStatus = "active"
	CustomerInactive CustomerStatus = "inactive"
)

type Customer struct {
	gorm.Model
	TotalGuests int            `gorm:"type:int;not null"`
	Name        string         `gorm:"type:varchar(255);null"`
	Orders      []Order        `gorm:"foreignKey:CustomerID"`
	Status      CustomerStatus `gorm:"type:customer_status; default:active"`
	StartTime   time.Time      `gorm:"type:timestamp;not null"`
	EndTime     *time.Time     `gorm:"type:timestamp;null"`
	TableID     uint           `gorm:"not null"`
	Table       *Table         `gorm:"foreignKey:TableID"`
}
