package entity

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	Chef   UserRole = "chef"
	Waiter UserRole = "waiter"
)

type User struct {
	gorm.Model
	FirstName string   `gorm:"type:varchar(255); not null"`
	LastName  string   `gorm:"type:varchar(255); not null"`
	Avatar    string   `gorm:"type:varchar(255);"`
	Phone     string   `gorm:"type:varchar(50); not null"`
	CNIC      string   `gorm:"type:varchar(50); not null"`
	Role      UserRole `gorm:"type:enum('chef', 'waiter'); not null"`
	Orders    []Order  `gorm:"foreignKey:UserID"`
}
