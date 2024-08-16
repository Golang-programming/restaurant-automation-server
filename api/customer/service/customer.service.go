package service

import (
	"errors"
	"fmt"
	"time"

	"github.co/golang-programming/restaurant/api/customer/dto"
	"github.co/golang-programming/restaurant/api/customer/enum"
	"github.co/golang-programming/restaurant/api/customer/repository"
	"github.co/golang-programming/restaurant/api/customer/store"
	"github.co/golang-programming/restaurant/api/entity"
	utils "github.co/golang-programming/restaurant/api/utils/encryption"
)

func CreateCustomer(input *dto.CreateCustomerInput) (string, error) {
	customer := &entity.Customer{
		TotalGuests: input.TotalGuests,
		Name:        input.Name,
		TableID:     input.TableID,
		StartDate:   time.Now(),
	}
	if err := repository.CreateCustomer(customer); err != nil {
		return "", err
	}

	storeCustomerInRedis(input, customer.ID)
	return utils.Encryptor(fmt.Sprint(customer.ID))
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

func DeactivateCustomerInRedis(customerID uint) error {
	return errors.New("utils.RemoveActiveUser(id)")
}

func storeCustomerInRedis(input *dto.CreateCustomerInput, customerID uint) {
	customer := &store.Customer{
		ID:        customerID,
		Status:    enum.CustomerActive,
		TableID:   input.TableID,
		StartTime: time.Now(),
	}
	store.Save(customer)
}
