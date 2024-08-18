package service

import (
	"github.co/golang-programming/restaurant/api/entity"
	"github.co/golang-programming/restaurant/api/user/repository"
)

func GetUserByPhoneNumber(phoneNumber string) (*entity.User, error) {
	return repository.GetUserByPhoneNumber(phoneNumber)
}
