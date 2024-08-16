package entity

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Number int     `gorm:"type:int;not null;unique"`
	Orders []Order `gorm:"foreignKey:TableID"` // One-to-many relationship with Order
}
