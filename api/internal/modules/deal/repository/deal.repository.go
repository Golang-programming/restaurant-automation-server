package repository

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"gorm.io/gorm"
)

var db *gorm.DB // This should be initialized in your main application

func GetDealByID(id uint) (*entity.Deal, error) {
	var deal entity.Deal
	if err := db.First(&deal, id).Error; err != nil {
		return nil, err
	}

	return &deal, nil
}
