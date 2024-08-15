package service

import (
	"fmt"

	"github.co/golang-programming/restaurant/api/customer/dto"
	"github.co/golang-programming/restaurant/api/customer/repository"
	"github.co/golang-programming/restaurant/api/customer/utils"
	"github.co/golang-programming/restaurant/api/entity"

	commonUtils "github.co/golang-programming/restaurant/api/utils/encryption"
)

func CreateCustomer(input *dto.CreateCustomerInput) (string, error) {
	customer := &entity.Customer{
		TotalGuests: input.TotalGuests,
		Name:        input.Name,
	}
	if err := repository.CreateCustomer(customer); err != nil {
		return "", err
	}

	utils.SetCustomerStatus(customer.ID, entity.CustomerActive)
	return commonUtils.Encryptor(fmt.Sprint(customer.ID))
}

func GetCustomerByID(id uint) (*entity.Customer, error) {
	return repository.GetCustomerByID(id)
}

func UpdateCustomer(id uint, input *dto.UpdateCustomerInput) (*entity.Customer, error) {
	customer, err := repository.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}

	customer.TotalGuests = input.TotalGuests
	customer.Name = input.Name
	if err := repository.UpdateCustomer(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func DeleteCustomer(id uint) error {
	customer, err := repository.GetCustomerByID(id)
	if err != nil {
		return err
	}
	return repository.DeleteCustomer(customer)
}

func ListCustomers() ([]*entity.Customer, error) {
	return repository.ListCustomers()
}

func RemoveCustomerFromActiveList(id uint) error {
	return utils.SetCustomerStatus(id, entity.CustomerInactive)
}
