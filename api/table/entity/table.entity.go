package entity

import (
	"gorm.io/gorm"
)

// Table represents the structure of a table in the database
type Table struct {
	gorm.Model
	NumberOfGuests int `gorm:"type:int;not null"`
	TableNumber    int `gorm:"type:int;not null"`
}
