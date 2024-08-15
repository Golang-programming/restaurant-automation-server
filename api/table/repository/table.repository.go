package repository

import (
	"github.co/golang-programming/restaurant/api/database"
	"github.co/golang-programming/restaurant/api/entity"
)

func CreateTable(table *entity.Table) error {
	return database.ActiveDB.Create(table).Error
}

func GetTableByID(id uint) (*entity.Table, error) {
	var table entity.Table
	if err := database.ActiveDB.First(&table, id).Error; err != nil {
		return nil, err
	}
	return &table, nil
}

func UpdateTable(table *entity.Table) error {
	return database.ActiveDB.Save(table).Error
}

func DeleteTable(table *entity.Table) error {
	return database.ActiveDB.Delete(table).Error
}

func ListTables() ([]*entity.Table, error) {
	var tables []*entity.Table
	if err := database.ActiveDB.Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}
