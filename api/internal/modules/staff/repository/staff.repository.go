package repository

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/pkg/database"
)

func GetStaffByPhoneNumber(phoneNumber string) (*entity.Staff, error) {
	var staff entity.Staff
	if err := database.ActiveDB.Where("phone_number = ?", phoneNumber).First(&staff).Error; err != nil {
		return nil, err
	}

	return &staff, nil
}

func CreateStaff(Staff *entity.Staff) error {
	return database.ActiveDB.Create(Staff).Error
}

func GetStaffByID(id uint) (*entity.Staff, error) {
	var staff entity.Staff
	if err := database.ActiveDB.First(&staff, id).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

func GetAllStaffs() ([]*entity.Staff, error) {
	var staffs []*entity.Staff
	if err := database.ActiveDB.Find(&staffs).Error; err != nil {
		return nil, err
	}
	return staffs, nil
}

func UpdateStaff(Staff *entity.Staff) error {
	return database.ActiveDB.Save(Staff).Error
}

func DeleteStaff(Staff *entity.Staff) error {
	return database.ActiveDB.Delete(Staff).Error
}
