package utils

import (
	"fmt"

	"github.co/golang-programming/restaurant/api/entity"
	"github.co/golang-programming/restaurant/api/redis"
)

func AddCustomerToActiveList(customerID uint, customer *entity.Customer) error {
	key := fmt.Sprintf("activeCustomer:%d", customerID)
	return redis.SaveToRedis(key, customer)
}

func GetActiveCustomers() ([]*entity.Customer, error) {
	var customers []*entity.Customer
	err := redis.GetFromRedis("activeCustomers", &customers)
	return customers, err
}

func RemoveCustomerFromActiveList(customerID uint) error {
	key := fmt.Sprintf("activeCustomer:%d", customerID)
	return redis.DeleteFromRedis(key)
}
