package service

import (
	"github.co/golang-programming/restaurant/api/entity"
	"github.co/golang-programming/restaurant/api/food/repository"
)

func GetFoodByID(id uint) (*entity.Food, error) {
	return repository.GetFoodByID(id)
}
