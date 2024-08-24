package service

import (
	"os"
	"strconv"

	"github.co/golang-programming/restaurant/api/bill/dto"
	"github.co/golang-programming/restaurant/api/bill/repository"
	customerService "github.co/golang-programming/restaurant/api/customer/service"
	"github.co/golang-programming/restaurant/api/entity"
	orderService "github.co/golang-programming/restaurant/api/order/service"
)

func CreateBill(input *dto.CreateBillInput) (*entity.Bill, error) {
	customer, err := customerService.GetCustomerByID(input.CustomerID)
	if err != nil {
		return nil, err
	}

	order, err := orderService.GetOrderByID(input.OrderID)
	if err != nil {
		return nil, err
	}

	var subtotal float64
	for _, item := range order.OrderItems {
		subtotal += item.Price
	}

	discount := CalculateDiscount(subtotal)
	total := subtotal - discount

	bill := &entity.Bill{
		Subtotal:   subtotal,
		Discount:   discount,
		Total:      total,
		OrderID:    input.OrderID,
		TableID:    customer.TableID,
		CustomerID: customer.ID,
	}

	if err := repository.CreateBill(bill); err != nil {
		return nil, err
	}

	return bill, nil
}

func CalculateDiscount(subtotal float64) float64 {
	discountPriceReach, _ := strconv.ParseFloat(os.Getenv("DISCOUNT_REACH"), 64)
	discountPercentage, _ := strconv.ParseFloat(os.Getenv("DISCOUNT_PERCENTAGE"), 64)

	if subtotal >= discountPriceReach {
		return subtotal * discountPercentage / 100
	}

	return 0
}

func GetBillByID(id uint) (*entity.Bill, error) {
	return repository.GetBillByID(id)
}

func GetBillByTableID(id uint) (*entity.Bill, error) {
	return repository.GetBillByTableID(id)
}

func RefreshBill(id uint) error {
	bill, err := repository.GetBillByID(id)
	if err != nil {
		return err
	}

	var subtotal float64

	for _, item := range bill.Order.OrderItems {
		subtotal += item.Price
	}

	discount := CalculateDiscount(subtotal)
	total := subtotal - discount

	bill.Subtotal = subtotal
	bill.Discount = discount
	bill.Total = total

	if err := repository.UpdateBill(bill); err != nil {
		return err
	}

	return nil
}

func UpdateDiscount(id uint, input *dto.UpdateDiscountInput) error {
	bill, err := repository.GetBillByID(id)
	if err != nil {
		return err
	}

	total := bill.Subtotal - input.DiscountAmount

	bill.Discount = input.DiscountAmount
	bill.Total = total

	if err := repository.UpdateBill(bill); err != nil {
		return err
	}

	return nil
}

func MarkBillPaid(id uint, input *dto.MarkBillPaidInput) error {
	bill, err := repository.GetBillByID(id)
	if err != nil {
		return err
	}

	bill.Paid = input.PaidAmount
	bill.Status = entity.BillPaid

	if err := repository.UpdateBill(bill); err != nil {
		return err
	}

	return nil
}

func DeleteBill(id uint) error {
	bill, err := repository.GetBillByID(id)
	if err != nil {
		return err
	}
	return repository.DeleteBill(bill)
}

func ListBills() ([]*entity.Bill, error) {
	return repository.ListBills()
}
