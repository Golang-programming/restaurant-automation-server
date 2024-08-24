package dto

import "github.co/golang-programming/restaurant/api/entity"

type CreateStaffInput struct {
	PhoneNumber string           `json:"phone_number" validate:"required"`
	Role        entity.StaffRole `json:"role" validate:"required,oneof=chef waiter"`
	CNIC        string           `json:"cnic" validate:"required"`
}

type UpdateStaffInput struct {
	FirstName string `json:"first_name" validate:"omitempty"`
	LastName  string `json:"last_name" validate:"omitempty"`
	Avatar    string `json:"avatar" validate:"omitempty"`
}
