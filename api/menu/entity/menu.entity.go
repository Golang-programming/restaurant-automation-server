package entity

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(255);not null"`
	Category  string    `gorm:"type:varchar(255);not null"`
	StartDate time.Time `gorm:"type:date; not null"`
	EndDate   time.Time `gorm:"type:date"`
}
