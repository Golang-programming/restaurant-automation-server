package entity

import (
	"gorm.io/gorm"
)

type StaffRole string

const (
	Chef   StaffRole = "chef"
	Waiter StaffRole = "waiter"
)

type Staff struct {
	gorm.Model
	FirstName   string    `gorm:"type:varchar(255)"`
	LastName    string    `gorm:"type:varchar(255)"`
	Avatar      string    `gorm:"type:varchar(255);"`
	PhoneNumber string    `gorm:"type:varchar(50); not null; unique"`
	CNIC        string    `gorm:"type:varchar(50); not null; unique"`
	Role        StaffRole `gorm:"type:staff_role; not null"`
}
