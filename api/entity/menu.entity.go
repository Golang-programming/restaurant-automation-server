package entity

type Menu struct {
	Name string `gorm:"type:varchar(255); not null"`
}
