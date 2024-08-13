package repository

import (
	"github.co/golang-programming/restaurant/api/entity"
	"gorm.io/gorm"
)

var db *gorm.DB // This should be initialized in your main application

func GetFoodByID(id uint) (*entity.Food, error) {
	var food entity.Food
	if err := db.First(&food, id).Error; err != nil {
		return nil, err
	}
	return &food, nil
}
