package service

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/internal/modules/deal/repository"
)

func GetDealByID(id uint) (*entity.Deal, error) {
	return repository.GetDealByID(id)
}
