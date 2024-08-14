package repository

import (
	"github.co/golang-programming/restaurant/api/database"
	"github.co/golang-programming/restaurant/api/entity"
)

func CreateMenu(menu *entity.Menu) error {
	return database.ActiveDB.Create(menu).Error
}

func GetMenuDetailsByID(id uint) (*entity.Menu, error) {
	var menu entity.Menu
	if err := database.ActiveDB.Preload("Foods").Preload("Deals").First(&menu, id).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}

func GetMenuByID(id uint) (*entity.Menu, error) {
	var menu entity.Menu
	if err := database.ActiveDB.First(&menu, id).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}

func UpdateMenu(menu *entity.Menu) error {
	return database.ActiveDB.Save(menu).Error
}

func DeleteMenu(menu *entity.Menu) error {
	return database.ActiveDB.Delete(menu).Error
}

func ListMenus() ([]entity.Menu, error) {
	var menus []entity.Menu
	if err := database.ActiveDB.Preload("Foods").Preload("Deals").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
