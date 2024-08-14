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

func GetOrderDetailsByID(orderID uint) (*entity.Order, error) {
	return repository.GetOrderDetailsByID(orderID)
}

func UpdateOrderStatus(orderID uint, input *dto.UpdateOrderStatusInput) (*entity.Order, error) {
	order, err := getOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = input.Status
	if err := repository.UpdateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}

func DeleteOrder(orderID uint) error {
	order, err := getOrderByID(orderID)
	if err != nil {
		return err
	}

	return repository.DeleteOrder(order)
}

func ListOrders() ([]entity.Order, error) {
	return repository.ListOrders()
}

func AddOrderItem(orderID uint, input *dto.AddOrderItemInput) (*entity.OrderItem, error) {
	_, err := getOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	orderItem := &entity.OrderItem{
		FoodID:   input.FoodID,
		OrderID:  orderID,
		Quantity: input.Quantity,
		Status:   entity.OrderItemOrdering,
	}

	if err := repository.AddOrderItem(orderItem); err != nil {
		return nil, err
	}

	return orderItem, nil
}

func UpdateOrderItemStatus(orderID, itemID uint, input *dto.UpdateOrderItemStatusInput) (*entity.OrderItem, error) {
	orderItem, err := getOrderItemByID(orderID)
	if err != nil {
		return nil, err
	}

	orderItem.Status = input.Status
	if err := repository.UpdateOrderItem(orderItem); err != nil {
		return nil, err
	}

	return orderItem, nil
}

func RemoveOrderItem(orderID, itemID uint) error {
	orderItem, err := getOrderItemByID(itemID)
	if err != nil {
		return err
	}
	return repository.RemoveOrderItem(orderItem)
}

func getOrderByID(orderID uint) (order *entity.Order, err error) {
	order, err = repository.GetOrderByID(orderID)

	return
}

func getOrderItemByID(itemID uint) (orderItem *entity.OrderItem, err error) {
	orderItem, err = repository.GetOrderItemByID(itemID)

	return
}
