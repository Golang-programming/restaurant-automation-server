package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255); not null"`
	LastName  string `gorm:"type:varchar(255); not null"`
	Avatar    string `gorm:"type:varchar(255);"`
	Phone     string `gorm:"type:varchar(50); not null"`
	CNIC      string `gorm:"type:varchar(50); not null"`
}
