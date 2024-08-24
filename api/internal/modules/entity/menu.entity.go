package entity

import "gorm.io/gorm"

type MenuStatus string

const (
	MenuActive   MenuStatus = "active"
	MenuInactive MenuStatus = "inactive"
)

type Menu struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255); not null; unique"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"type:menu_status; not null; default: 'active'"`
	Foods       []Food `gorm:"foreignKey:MenuID"`
	Deals       []Deal `gorm:"foreignKey:MenuID"`
}
