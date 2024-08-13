package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	TotalGuests int   `gorm:"type:int;not null"`
	TableID     uint  `gorm:"not null"`
	Table       Table `gorm:"foreignKey:TableID"`
}
