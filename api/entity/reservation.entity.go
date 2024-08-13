package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	CustomerID uint      `gorm:"not null"`
	Customer   *Customer `gorm:"foreignKey:CustomerID"`
	TableID    uint      `gorm:"not null"`
	Table      *Table    `gorm:"foreignKey:TableID"`
	Date       time.Time `gorm:"not null"`
}
