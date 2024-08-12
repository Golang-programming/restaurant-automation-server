package entity

import (
	"gorm.io/gorm"
)

type TableStatus string

const (
	OBSERVED  TableStatus = "observed"
	AVAILABLE TableStatus = "available"
)

type Table struct {
	gorm.Model
	Number int         `gorm:"type:int;not null;unique"`
	Status TableStatus `gorm:"type:enum('observed', 'available');not null"` // Define the enum type here
}
