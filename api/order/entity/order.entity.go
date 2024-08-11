package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderDate time.Time `gorm:"type:datetime;not null"`
	TableID   string    `gorm:"type:varchar(255);"`
}
