package entity

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255); not null; unique"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"type:enum('active', 'inactive'); not null; default: 'active'"`
	Foods       []Food `gorm:"foreignKey:MenuID"`
	Deals       []Deal `gorm:"foreignKey:MenuID"`
}
