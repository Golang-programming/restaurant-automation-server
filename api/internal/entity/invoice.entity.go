package entity

import "gorm.io/gorm"

type InvoiceItem struct {
	gorm.Model
	Description string   `gorm:"type:varchar(255); not null"`
	Quantity    int      `gorm:"type:int; not null"`
	Price       float64  `gorm:"type:decimal(10,2); not null"`
	Total       float64  `gorm:"type:decimal(10,2); not null"`
	InvoiceID   uint     `gorm:"not null"`
	Invoice     *Invoice `gorm:"foreignKey:InvoiceID"`
	FoodID      uint     `gorm:"not null"`
	Food        *Food    `gorm:"foreignKey:FoodID"`
}

type Invoice struct {
	gorm.Model
	Items  []InvoiceItem `gorm:"foreignKey:InvoiceID"`
	PdfUrl string        `gorm:"type:varchar(255); not null"`
	BillID uint          `gorm:"not null"`
	Bill   *Bill         `gorm:"foreignKey:BillID"`
}
