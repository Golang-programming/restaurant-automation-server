package service

import (
	"github.co/golang-programming/restaurant/api/customer/dto"
	"github.co/golang-programming/restaurant/api/customer/repository"
	"github.co/golang-programming/restaurant/api/customer/utils"
	"github.co/golang-programming/restaurant/api/entity"
)

func CreateCustomer(input *dto.CreateCustomerInput) (*entity.Customer, error) {
	customer := &entity.Customer{
		TotalGuests: input.TotalGuests,
		Name:        input.Name,
	}
	if err := repository.CreateCustomer(customer); err != nil {
		return nil, err
	}

	utils.AddCustomerToActiveList(customer.ID, customer)
	return customer, nil
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
	return utils.RemoveCustomerFromActiveList(id)
}
