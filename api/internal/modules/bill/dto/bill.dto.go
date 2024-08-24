package dto

type CreateBillInput struct {
	CustomerID uint `json:"customer_id" binding:"required"`
	OrderID    uint `json:"order_id" binding:"required"`
}

type MarkBillPaidInput struct {
	PaidAmount float64 `json:"paid_amount" binding:"required"`
}

type UpdateDiscountInput struct {
	DiscountAmount     float64 `json:"discount_amount" binding:"omitempty"`
}
