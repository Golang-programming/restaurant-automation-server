package utils

import (
	"fmt"

	"github.co/golang-programming/restaurant/api/entity"
	"github.co/golang-programming/restaurant/api/redis"
)

func SetCustomerStatus(customerID uint, status entity.CustomerStatus) error {
	key := fmt.Sprintf("customer:%d:status", customerID)
	return redis.SaveToRedis(key, status)
}

func GetActiveCustomers() ([]*entity.Customer, error) {
}
