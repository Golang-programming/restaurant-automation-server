package repository

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/pkg/database"
)

func CreateFood(food *entity.Food) error {
	return database.ActiveDB.Create(food).Error
}

func GetFoodByID(id uint) (*entity.Food, error) {
	var food entity.Food
	if err := database.ActiveDB.First(&food, id).Error; err != nil {
		return nil, err
	}
	return &food, nil
}

func UpdateFood(food *entity.Food) error {
	return database.ActiveDB.Save(food).Error
}

func DeleteFood(food *entity.Food) error {
	return database.ActiveDB.Delete(food).Error
}

func ListFoods() ([]*entity.Food, error) {
	var foods []*entity.Food
	if err := database.ActiveDB.Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}
