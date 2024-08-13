package entity

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255); not null"`
	Foods []Food `gorm:"foreignKey:MenuID"`
}
