package repository

import (
	"github.co/golang-programming/restaurant/api/database"
	"github.co/golang-programming/restaurant/api/entity"
)

func GetUserByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User
	if err := database.ActiveDB.First(&user, phoneNumber).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
