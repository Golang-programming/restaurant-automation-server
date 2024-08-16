package service

import (
	"github.co/golang-programming/restaurant/api/entity"
	"github.co/golang-programming/restaurant/api/food/dto"
	"github.co/golang-programming/restaurant/api/food/repository"
)

func CreateFood(input *dto.CreateFoodInput) (*entity.Food, error) {
	food := &entity.Food{
		Name:         input.Name,
		Price:        input.Price,
		Description:  input.Description,
		Images:       input.Images,
		Status:       input.Status,
		ParentFoodID: input.ParentFoodID,
		MenuID:       input.MenuID,
		Category:     input.Category,
	}

	if err := repository.CreateFood(food); err != nil {
		return nil, err
	}

	return food, nil
}

func GetFoodByID(id uint) (*entity.Food, error) {
	return repository.GetFoodByID(id)
}

func UpdateFood(id uint, input *dto.UpdateFoodInput) (*entity.Food, error) {
	food, err := repository.GetFoodByID(id)
	if err != nil {
		return nil, err
	}

	food.Name = input.Name
	food.Price = input.Price
	food.Description = input.Description
	food.Images = input.Images
	food.Status = input.Status
	food.ParentFoodID = input.ParentFoodID
	food.MenuID = input.MenuID
	food.Category = input.Category

	if err := repository.UpdateFood(food); err != nil {
		return nil, err
	}

	return food, nil
}

func DeleteFood(id uint) error {
	food, err := repository.GetFoodByID(id)
	if err != nil {
		return err
	}
	return repository.DeleteFood(food)
}

func ListFoods() ([]*entity.Food, error) {
	return repository.ListFoods()
}
