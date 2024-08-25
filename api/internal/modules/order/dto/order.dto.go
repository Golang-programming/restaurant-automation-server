package dto

import "github.co/golang-programming/restaurant/api/internal/entity"

type CreateOrderInput struct {
	CustomerID uint `json:"customer_id" validate:"required"`
}

type UpdateOrderStatusInput struct {
	Status entity.OrderStatus `json:"status" validate:"required,oneof='pending' 'confirmed' 'cancelled' 'making' 'placed' 'billing' 'complete' 'failed'"`
}

type AddOrderItemInput struct {
	ProductID   uint                        `json:"product_id" validate:"required"`
	ProductType entity.OrderItemProductType `json:"status" validate:"omitempty,oneof='food', 'deal'"`
	Quantity    int                         `json:"quantity" validate:"omitempty,min=1"`
}

type UpdateOrderItemStatusInput struct {
	Status entity.OrderItemStatus `json:"status" validate:"required,oneof='ordering' 'making' 'placed'"`
}

type OrderResponse struct {
	ID         uint                `json:"id"`
	Status     entity.OrderStatus  `json:"status"`
	CustomerID uint                `json:"customer_id"`
	OrderItems []OrderItemResponse `json:"order_items"`
}

type OrderItemResponse struct {
	ID        uint                   `json:"id"`
	ProductID uint                   `json:"food_id"`
	Quantity  int                    `json:"quantity"`
	Status    entity.OrderItemStatus `json:"status"`
}

type OrderListResponse struct {
	Orders []OrderResponse `json:"orders"`
}
