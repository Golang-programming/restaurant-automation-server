package repository

import (
	"github.co/golang-programming/restaurant/api/entity"
	"gorm.io/gorm"
)

var db *gorm.DB // This should be initialized in your main application

func CreateCustomer(customer *entity.Customer) error {
	return db.Create(customer).Error
}

func GetCustomerByID(id string) (*entity.Customer, error) {
	var customer entity.Customer
	if err := db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func UpdateCustomer(customer *entity.Customer) error {
	return db.Save(customer).Error
}

func DeleteCustomer(customer *entity.Customer) error {
	return db.Delete(customer).Error
}

func ListCustomers() ([]*entity.Customer, error) {
	var customers []*entity.Customer
	if err := db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
