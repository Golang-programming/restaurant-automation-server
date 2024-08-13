package entity

import "gorm.io/gorm"

type InvoiceItem struct {
	gorm.Model
}

type Invoice struct {
	gorm.Model
	Items  []InvoiceItem `gorm:"type:json;default:[]"`
	PdfUrl string        `gorm:"type:varchar(255); not null"`
	BillID uint          `gorm:"not null"`
	Bill   Bill          `gorm:"foreignKey:BillID"`
}
