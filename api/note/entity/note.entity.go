package entity

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Text  string `gorm:"type:text;not null"`
	Title string `gorm:"type:varchar(255);not null"`
}
