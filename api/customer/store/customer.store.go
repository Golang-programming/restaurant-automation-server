package store

import (
	"encoding/json"
	"fmt"

	"github.co/golang-programming/restaurant/api/redis"
)

type Customer struct {
	ID             uint   `json:"id"`
	Status         string `json:"status"`
	TotalGuests    int    `json:"total_guests"`
	TableID        uint   `json:"table_id"`
	CurrentOrderID uint   `json:"current_order_id"`
}

func (c *Customer) Serialize() (string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *Customer) Deserialize(data string) error {
	return json.Unmarshal([]byte(data), c)
}

func (c *Customer) Save() error {
	key := fmt.Sprintf("customer:%d", c.ID)
	data, err := c.Serialize()
	if err != nil {
		return err
	}
	return redis.Set(key, data)
}

func GetCustomer(id uint) (*Customer, error) {
	key := fmt.Sprintf("customer:%d", id)
	data, err := redis.Get(key)
	if err != nil {
		return nil, err
	}
	c := &Customer{}
	if err := c.Deserialize(data); err != nil {
		return nil, err
	}
	return c, nil
}
