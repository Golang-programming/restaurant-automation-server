package service

import (
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
func AddFoodToMenu(menuID uint, foodID uint) error {
	menu, err := repository.GetMenuByID(menuID)
	if err != nil {
		return err
	}

	food, err := foodService.GetFoodByID(foodID)
	if err != nil {
		return err
	}

	if err := repository.AddFoodToMenu(menu, food); err != nil {
		return err
	}

	return nil
}

func RemoveFoodFromMenu(menuID uint, foodID uint) error {
	menu, err := repository.GetMenuByID(menuID)
	if err != nil {
		return err
	}

	food, err := foodService.GetFoodByID(foodID)
	if err != nil {
		return err
	}

	if err := repository.RemoveFoodFromMenu(menu, food); err != nil {
		return err
	}

	return nil
}

func AddDealToMenu(menuID uint, dealID uint) error {
	menu, err := repository.GetMenuByID(menuID)
	if err != nil {
		return err
	}

	deal, err := dealService.GetDealByID(dealID)
	if err != nil {
		return err
	}

	if err := repository.AddDealToMenu(menu, deal); err != nil {
		return err
	}

	return nil
}

func RemoveDealFromMenu(menuID uint, dealID uint) error {
	menu, err := repository.GetMenuByID(menuID)
	if err != nil {
		return err
	}

	deal, err := dealService.GetDealByID(dealID)
	if err != nil {
		return err
	}

	if err := repository.RemoveDealFromMenu(menu, deal); err != nil {
		return err
	}

	return nil
}
