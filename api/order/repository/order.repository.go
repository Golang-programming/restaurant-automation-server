package repository

import (
	"github.co/golang-programming/restaurant/api/database"
	"github.co/golang-programming/restaurant/api/entity"
)

func CreateOrder(order *entity.Order) error {
	return database.ActiveDB.Create(order).Error
}

func GetOrderDetailsByID(id uint) (*entity.Order, error) {
	var order entity.Order
	if err := database.ActiveDB.Preload("OrderItems").Preload("OrderItems.Food").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func GetOrderByID(id uint) (*entity.Order, error) {
	var order entity.Order
	if err := database.ActiveDB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func UpdateOrder(order *entity.Order) error {
	return database.ActiveDB.Save(order).Error
}

func DeleteOrder(order *entity.Order) error {
	return database.ActiveDB.Delete(order).Error
}

func ListOrders() ([]entity.Order, error) {
	var orders []entity.Order
	if err := database.ActiveDB.Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func AddOrderItem(orderItem *entity.OrderItem) error {
	return database.ActiveDB.Create(orderItem).Error
}

func GetOrderItemByID(itemID uint) (*entity.OrderItem, error) {
	var orderItem entity.OrderItem
	if err := database.ActiveDB.Where("id = ?", itemID).First(&orderItem).Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func UpdateOrderItem(orderItem *entity.OrderItem) error {
	return database.ActiveDB.Save(orderItem).Error
}

func RemoveOrderItem(orderItem *entity.OrderItem) error {
	return database.ActiveDB.Delete(orderItem).Error
}
