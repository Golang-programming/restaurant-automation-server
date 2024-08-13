package service

import (
	"errors"

	dealService "github.co/golang-programming/restaurant/api/deal/service"
	"github.co/golang-programming/restaurant/api/entity"
	foodService "github.co/golang-programming/restaurant/api/food/service"
	"github.co/golang-programming/restaurant/api/menu/dto"
	"github.co/golang-programming/restaurant/api/menu/repository"
)

func CreateMenu(input *dto.CreateMenuInput) (*entity.Menu, error) {
	menu := &entity.Menu{
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
	}
	if err := repository.CreateMenu(menu); err != nil {
		return nil, err
	}
	return menu, nil
}

func GetMenuDetailsByID(id uint) (*entity.Menu, error) {
	return repository.GetMenuDetailsByID(id)
}

func GetMenuByID(id uint) (*entity.Menu, error) {
	return repository.GetMenuByID(id)
}

func UpdateMenu(id uint, input *dto.UpdateMenuInput) (*entity.Menu, error) {
	menu, err := repository.GetMenuByID(id)
	if err != nil {
		return nil, err
	}
	if input.Name != "" {
		menu.Name = input.Name
	}
	if input.Description != "" {
		menu.Description = input.Description
	}
	if input.Status != "" {
		menu.Status = input.Status
	}
	if err := repository.UpdateMenu(menu); err != nil {
		return nil, err
	}
	return menu, nil
}

func DeleteMenu(id uint) error {
	menu, err := repository.GetMenuByID(id)
	if err != nil {
		return err
	}
	return repository.DeleteMenu(menu)
}

func ListMenus() ([]entity.Menu, error) {
	return repository.ListMenus()
}

/******* add and remove items from menu *******/
func AddFoodToMenu(menuID uint, input *dto.AddFoodToMenuInput) (*entity.Menu, error) {
	menu, err := repository.GetMenuByID(menuID)
	if err != nil {
		return nil, err
	}

	food, err := foodService.GetFoodByID(input.FoodID)
	if err != nil {
		return nil, errors.New("food not found")
	}

	menu.Foods = append(menu.Foods, *food)
	if err := repository.UpdateMenu(menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func RemoveFoodFromMenu(menuID uint, input *dto.RemoveFoodFromMenuInput) (*entity.Menu, error) {
	menu, err := repository.GetMenuByID(menuID)
	if err != nil {
		return nil, err
	}

	var updatedFoods []entity.Food
	for _, food := range menu.Foods {
		if food.ID != input.FoodID {
			updatedFoods = append(updatedFoods, food)
		}
	}

	menu.Foods = updatedFoods
	if err := repository.UpdateMenu(menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func AddDealToMenu(menuID uint, input *dto.AddDealToMenuInput) (*entity.Menu, error) {
	menu, err := repository.GetMenuByID(menuID)
	if err != nil {
		return nil, err
	}

	deal, err := dealService.GetDealByID(input.DealID)
	if err != nil {
		return nil, errors.New("deal not found")
	}

	menu.Deals = append(menu.Deals, *deal)
	if err := repository.UpdateMenu(menu); err != nil {
		return nil, err
	}

	return menu, nil
}

func RemoveDealFromMenu(menuID uint, input *dto.RemoveDealFromMenuInput) (*entity.Menu, error) {
	menu, err := repository.GetMenuByID(menuID)
	if err != nil {
		return nil, err
	}

	var updatedDeals []entity.Deal
	for _, deal := range menu.Deals {
		if deal.ID != input.DealID {
			updatedDeals = append(updatedDeals, deal)
		}
	}

	menu.Deals = updatedDeals
	if err := repository.UpdateMenu(menu); err != nil {
		return nil, err
	}

	return menu, nil
}
