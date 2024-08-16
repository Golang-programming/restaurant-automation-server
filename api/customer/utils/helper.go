/* package utils

import (
	"fmt"
	"strconv"





	"github.co/golang-programming/restaurant/api/redis"
)

const activeUsersKey = "users:active"

func AddActiveUser(userID uint) error {
	return redis.SAdd(activeUsersKey, userID)
}

func RemoveActiveUser(userID uint) error {
	return redis.SRem(activeUsersKey, userID)
}

func GetActiveUsers() ([]uint, error) {
	members, err := redis.SMembers(activeUsersKey)
	if err != nil {
		return nil, err
	}

	var userIDs []uint
	for _, member := range members {
		id, err := strconv.ParseUint(member, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse userID: %v", err)
		}
		userIDs = append(userIDs, uint(id))
	}

	return userIDs, nil
}
*/

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.co/golang-programming/restaurant/api/redis"
)

// Initialize a context for Redis operations.
var ctx = context.Background()

// Customer model for Redis
type Customer struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	TotalGuests    int    `json:"total_guests"`
	TableID        uint   `json:"table_id"`
	CurrentOrderID uint   `json:"current_order_id"`
}

// Order model for Redis
type Order struct {
	ID         uint   `json:"id"`
	Status     string `json:"status"`
	CustomerID uint   `json:"customer_id"`
	TableID    uint   `json:"table_id"`
	Items      string `json:"items"` // Store as a JSON string
}

// Table model for Redis
type Table struct {
	ID               uint   `json:"id"`
	Status           string `json:"status"`
	CurrentOrderID   uint   `json:"current_order_id"`
	CurrentCustomers []uint `json:"current_customers"`
}

// Bill model for Redis
type Bill struct {
	ID         uint    `json:"id"`
	Subtotal   float64 `json:"subtotal"`
	Discount   float64 `json:"discount"`
	Total      float64 `json:"total"`
	Paid       float64 `json:"paid"`
	Status     string  `json:"status"`
	OrderID    uint    `json:"order_id"`
	CustomerID uint    `json:"customer_id"`
}

// Food model for Redis
type Food struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Category    string  `json:"category"`
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

// Similar methods for other models...

func (o *Order) Serialize() (string, error) {
	bytes, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (o *Order) Deserialize(data string) error {
	return json.Unmarshal([]byte(data), o)
}

func (o *Order) Save() error {
	key := fmt.Sprintf("order:%d", o.ID)
	data, err := o.Serialize()
	if err != nil {
		return err
	}
	return redis.Set(key, data)
}

func GetOrder(id uint) (*Order, error) {
	key := fmt.Sprintf("order:%d", id)
	data, err := redis.Get(key)
	if err != nil {
		return nil, err
	}
	o := &Order{}
	if err := o.Deserialize(data); err != nil {
		return nil, err
	}
	return o, nil
}

// You can create similar methods for the `Table`, `Bill`, and `Food` structs.
