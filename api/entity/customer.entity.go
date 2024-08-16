package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	TotalGuests int        `gorm:"type:int;not null"`
	Name        string     `gorm:"type:varchar(255);null"`
	Orders      []Order    `gorm:"foreignKey:CustomerID"`
	StartDate   time.Time  `gorm:"type:timestamp;not null"`
	EndDate     *time.Time `gorm:"type:timestamp;null"`
	TableID     uint       `gorm:"not null"`
	Table       *Table     `gorm:"foreignKey:TableID"`
}
