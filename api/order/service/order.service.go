package service

import (
	billService "github.co/golang-programming/restaurant/api/bill/service"
	dealService "github.co/golang-programming/restaurant/api/deal/service"
	"github.co/golang-programming/restaurant/api/entity"
	foodService "github.co/golang-programming/restaurant/api/food/service"
	"github.co/golang-programming/restaurant/api/order/dto"
	"github.co/golang-programming/restaurant/api/order/repository"
)

func CreateOrder(input *dto.CreateOrderInput) (*entity.Order, error) {
	order := &entity.Order{
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
	order, err := GetOrderByID(orderID)
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
	order, err := GetOrderByID(orderID)
	if err != nil {
		return err
	}

	return repository.DeleteOrder(order)
}

func ListOrders() ([]entity.Order, error) {
	return repository.ListOrders()
}

func AddOrderItem(orderID uint, input *dto.AddOrderItemInput) (*entity.OrderItem, error) {
	order, err := GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	price, err := getProductByIDAndType(input.ProductID, input.ProductType)
	if err != nil {
		return nil, err
	}

	orderItem := &entity.OrderItem{
		ProductID:   input.ProductID,
		ProductType: input.ProductType,
		OrderID:     orderID,
		Price:       float64(input.Quantity) * price,
		Quantity:    input.Quantity,
		Status:      entity.OrderItemOrdering,
	}

	if err := repository.AddOrderItem(orderItem); err != nil {
		return nil, err
	}

	if !order.Bill {
		billService.RefreshBill(order.bill)
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

func GetCustomerOrders(customerID uint) (*[]entity.Order, error) {
	orders, err := repository.GetCustomerOrders(customerID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(orderID uint) (order *entity.Order, err error) {
	order, err = repository.GetOrderByID(orderID)

	return
}

func getOrderItemByID(itemID uint) (orderItem *entity.OrderItem, err error) {
	orderItem, err = repository.GetOrderItemByID(itemID)

	return
}

func getProductByIDAndType(productID uint, productType entity.OrderItemProductType) (float64, error) {
	if productType == entity.OrderItemFood {
		food, err := foodService.GetFoodByID(productID)
		if err != nil {
			return 0, err
		}

		return food.Price, nil
	}

	_, err := dealService.GetDealByID(productID)
	if err != nil {
		return 0, err
	}

	return 0, nil
}
