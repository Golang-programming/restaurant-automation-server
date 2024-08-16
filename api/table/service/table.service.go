package service

import (
	"github.co/golang-programming/restaurant/api/entity"
	"github.co/golang-programming/restaurant/api/table/dto"
	"github.co/golang-programming/restaurant/api/table/repository"
	"github.co/golang-programming/restaurant/api/table/utils"
)

func CreateTable(input *dto.CreateTableInput) (*entity.Table, error) {
	table := &entity.Table{
		Number: input.Number,
		Status: input.Status,
	}
	if err := repository.CreateTable(table); err != nil {
		return nil, err
	}

	utils.MarkTableAvailable(table.ID)
	return table, nil
}

func GetTableByID(id uint) (*entity.Table, error) {
	return repository.GetTableByID(id)
}

func MarkTableObserved(id uint) error {
	return utils.MarkTableObserved(id)
}

func UpdateTable(id uint, input *dto.UpdateTableInput) (*entity.Table, error) {
	table, err := repository.GetTableByID(id)
	if err != nil {
		return nil, err
	}

	table.Number = input.Number
	table.Status = input.Status
	if err := repository.UpdateTable(table); err != nil {
		return nil, err
	}

	return table, nil
}

func DeleteTable(id uint) error {
	table, err := repository.GetTableByID(id)
	if err != nil {
		return err
	}
	return repository.DeleteTable(table)
}

func ListTables() ([]*entity.Table, error) {
	return repository.ListTables()
}
