package service

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/internal/modules/staff/dto"
	"github.co/golang-programming/restaurant/api/internal/modules/staff/repository"
)

func GetStaffByPhoneNumber(phoneNumber string) (*entity.Staff, error) {
	return repository.GetStaffByPhoneNumber(phoneNumber)
}

func CreateStaff(input *dto.CreateStaffInput) (*entity.Staff, error) {
	Staff := &entity.Staff{
		PhoneNumber: input.PhoneNumber,
		Role:        input.Role,
		CNIC:        input.CNIC,
	}
	if err := repository.CreateStaff(Staff); err != nil {
		return nil, err
	}
	return Staff, nil
}

func UpdateStaff(id uint, input *dto.UpdateStaffInput) (*entity.Staff, error) {
	Staff, err := repository.GetStaffByID(id)
	if err != nil {
		return nil, err
	}

	Staff.FirstName = input.FirstName
	Staff.LastName = input.LastName
	Staff.Avatar = input.Avatar

	if err := repository.UpdateStaff(Staff); err != nil {
		return nil, err
	}

	return Staff, nil
}

func GetAllStaffs() ([]*entity.Staff, error) {
	return repository.GetAllStaffs()
}

func GetStaffByID(id uint) (*entity.Staff, error) {
	return repository.GetStaffByID(id)
}

func DeleteStaff(id uint) error {
	Staff, err := repository.GetStaffByID(id)
	if err != nil {
		return err
	}
	return repository.DeleteStaff(Staff)
}
