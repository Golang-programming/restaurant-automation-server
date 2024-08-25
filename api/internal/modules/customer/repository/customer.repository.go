package repository

import (
	"github.co/golang-programming/restaurant/api/internal/entity"
	"github.co/golang-programming/restaurant/api/pkg/database"
)

func CreateCustomer(customer *entity.Customer) error {
	return database.ActiveDB.Create(customer).Error
}

func GetCustomerByID(id uint) (*entity.Customer, error) {
	var customer entity.Customer
	if err := database.ActiveDB.First(&customer, id).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func UpdateCustomer(customer *entity.Customer) error {
	return database.ActiveDB.Save(customer).Error
}

func DeleteCustomer(customer *entity.Customer) error {
	return database.ActiveDB.Delete(customer).Error
}

func ListCustomers() ([]*entity.Customer, error) {
	var customers []*entity.Customer
	if err := database.ActiveDB.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
