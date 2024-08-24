package service

import (
	"github.co/golang-programming/restaurant/api/deal/repository"
	"github.co/golang-programming/restaurant/api/entity"
)

func GetDealByID(id uint) (*entity.Deal, error) {
	return repository.GetDealByID(id)
}
