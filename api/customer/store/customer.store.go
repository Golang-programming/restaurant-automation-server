package store

import (
	"encoding/json"
	"fmt"
	"time"

	"github.co/golang-programming/restaurant/api/customer/enum"
	"github.co/golang-programming/restaurant/api/redis"
	"github.co/golang-programming/restaurant/api/utils"
)

type Customer struct {
	ID        uint                `json:"id"`
	Status    enum.CustomerStatus `json:"status"`
	StartTime time.Time           `json:"start_time"`
	EndTime   time.Time           `json:"end_time"`
	TableID   uint                `json:"table_id"`
}

func Serialize(customer *Customer) (string, error) {
	bytes, err := json.Marshal(customer)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Deserialize(customer *Customer, data string) error {
	return json.Unmarshal([]byte(data), customer)
}

func Save(customer *Customer) error {
	key := fmt.Sprintf("customer:%d", customer.ID)
	data, err := Serialize(customer)
	if err != nil {
		return err
	}

	expiration := utils.CalculateSecondsUntilMidnight()

	return redis.Set(key, data, time.Duration(expiration)*time.Second)
}

func GetCustomer(id uint) (*Customer, error) {
	key := fmt.Sprintf("customer:%d", id)
	data, err := redis.Get(key)
	if err != nil {
		return nil, err
	}

	customer := &Customer{}
	if err := Deserialize(customer, data); err != nil {
		return nil, err
	}
	return customer, nil
}

func GetAllCustomers() ([]*Customer, error) {
	pattern := "customer:*"
	keys, err := redis.Keys(pattern)
	if err != nil {
		return nil, err
	}

	var customers []*Customer
	for _, key := range keys {
		data, err := redis.Get(key)
		if err != nil {
			return nil, err
		}
		customer := &Customer{}
		if err := Deserialize(customer, data); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func DeactivateCustomer(id uint) error {
	customer, err := GetCustomer(id)
	if err != nil {
		return err
	}

	customer.Status = enum.CustomerInactive
	return Save(customer)
}

func DeleteCustomer(id uint) error {
	key := fmt.Sprintf("customer:%d", id)
	return redis.Del(key)
}
