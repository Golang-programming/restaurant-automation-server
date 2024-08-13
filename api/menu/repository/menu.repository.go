package repository

import (
	"github.co/golang-programming/restaurant/api/entity"
	"gorm.io/gorm"
)

var db *gorm.DB // This should be initialized in your main application

func CreateMenu(menu *entity.Menu) error {
	return db.Create(menu).Error
}

func GetMenuByID(id uint) (*entity.Menu, error) {
	var menu entity.Menu
	if err := db.Preload("Foods").Preload("Deals").First(&menu, id).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func UpdateMenu(menu *entity.Menu) error {
	return db.Save(menu).Error
}

func DeleteMenu(menu *entity.Menu) error {
	return db.Delete(menu).Error
}

func ListMenus() ([]entity.Menu, error) {
	var menus []entity.Menu
	if err := db.Preload("Foods").Preload("Deals").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
