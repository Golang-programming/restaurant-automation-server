package service

import (
	"github.co/golang-programming/restaurant/api/entity"
	"github.co/golang-programming/restaurant/api/order/dto"
	"github.co/golang-programming/restaurant/api/order/repository"
)

func CreateOrder(input *dto.CreateOrderInput) (*entity.Order, error) {
	order := &entity.Order{
		Status:     entity.OrderPending,
		TableID:    input.TableID,
		CustomerID: input.CustomerID,
	}

	if err := repository.CreateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func GetOrderByID(id uint) (*entity.Order, error) {
	return repository.GetOrderByID(id)
}

func UpdateOrder(id uint, input *dto.UpdateOrderInput) (*entity.Order, error) {
	order, err := repository.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	if input.Status != "" {
		order.Status = input.Status
	}

	if err := repository.UpdateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func DeleteOrder(id uint) error {
	order, err := repository.GetOrderByID(id)
	if err != nil {
		return err
	}

	return repository.DeleteOrder(order)
}

func ListOrders() ([]entity.Order, error) {
	return repository.ListOrders()
}

// Order Items
func AddOrderItem(orderID uint, input *dto.AddOrderItemInput) (*entity.OrderItem, error) {
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	orderItem := &entity.OrderItem{
		FoodID:   input.FoodID,
		OrderID:  orderID,
		Quantity: input.Quantity,
		Status:   entity.OrderItemOrdering,
	}

	order.OrderItems = append(order.OrderItems, *orderItem)

	if err := repository.UpdateOrder(order); err != nil {
		return nil, err
	}

	return orderItem, nil
}

func UpdateOrderItem(orderID, itemID uint, input *dto.UpdateOrderItemInput) (*entity.OrderItem, error) {
	orderItem, err := repository.GetOrderItemByID(orderID, itemID)
	if err != nil {
		return nil, err
	}

	if input.Status != "" {
		orderItem.Status = input.Status
	}

	if input.Quantity > 0 {
		orderItem.Quantity = input.Quantity
	}

	if err := repository.UpdateOrderItem(orderItem); err != nil {
		return nil, err
	}

	return orderItem, nil
}

func RemoveOrderItem(orderID, itemID uint) error {
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	orderItem, err := repository.GetOrderItemByID(orderID, itemID)
	if err != nil {
		return err
	}

	if err := repository.DeleteOrderItem(orderItem); err != nil {
		return err
	}

	return repository.UpdateOrder(order)
}
